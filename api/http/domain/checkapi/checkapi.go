// Package checkapi maintains the web based api for system access.
package checkapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/roca/ugo-sfd-k8s/business/api/sqldb"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/tooling/environment"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

type api struct {
	build string
	log   *logger.Logger
	db    *sqlx.DB
}

func newAPI(build string, log *logger.Logger, db *sqlx.DB) *api {
	return &api{
		build: build,
		db:    db,
		log:   log,
	}
}

func (api *api) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	status := "ok"
	statusCode := http.StatusOK

	sslMode := "require"
	if environment.GetBoolEnv("SALES_DB_DISABLE_TLS", true) {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	// psql --host=$SALES_DB_HOST_PORT --port=5432 --dbname=$SALES_DB_NAME --username=$SALES_DB_USER

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(environment.GetStrEnv("SALES_DB_USER", "postgres"), environment.GetStrEnv("SALES_DB_PASSWORD", "postgres")),
		Host:     environment.GetStrEnv("SALES_DB_HOST_PORT", "database-service.sales-system.svc.cluster.local"),
		Path:     environment.GetStrEnv("SALES_DB_NAME", "postgres"),
		RawQuery: q.Encode(),
	}

	if err := sqldb.StatusCheck(ctx, api.db); err != nil {
		status = fmt.Sprintf("db not ready: %s, '%s'", u.String(), err)
		statusCode = http.StatusInternalServerError
		api.log.Info(ctx, "readiness failure", "status", status)
	}

	data := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	return web.Respond(ctx, w, data, statusCode)
}

func (api *api) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := struct {
		Status     string `json:"status,omitempty"`
		Build      string `json:"build,omitempty"`
		Host       string `json:"host,omitempty"`
		Name       string `json:"name,omitempty"`
		PodIP      string `json:"podIP,omitempty"`
		Node       string `json:"node,omitempty"`
		Namespace  string `json:"namespace,omitempty"`
		GOMAXPROCS int    `json:"GOMAXPROCS,omitempty"`
	}{
		Status:     "up",
		Build:      api.build,
		Host:       host,
		Name:       os.Getenv("KUBERNETES_NAME"),
		PodIP:      os.Getenv("KUBERNETES_POD_IP"),
		Node:       os.Getenv("KUBERNETES_NODE_NAME"),
		Namespace:  os.Getenv("KUBERNETES_NAMESPACE"),
		GOMAXPROCS: runtime.GOMAXPROCS(0),
	}

	// This handler provides a free timer loop.

	return web.Respond(ctx, w, data, http.StatusOK)

}
