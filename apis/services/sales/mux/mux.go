// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/roca/ugo-sfd-k8s/apis/services/api/mid"
	"github.com/roca/ugo-sfd-k8s/apis/services/sales/route/sys/checkapi"
	"github.com/roca/ugo-sfd-k8s/app/api/authclient"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(build string, log *logger.Logger, db *sqlx.DB, authClient *authclient.Client, shutdown chan os.Signal) *web.App {

	loggerMiddleware := mid.Logger(log)
	errorsMiddleware := mid.Errors(log)
	metricsMiddleware := mid.Metrics()
	panicMiddleware := mid.Panics()

	app := web.NewApp(shutdown, loggerMiddleware, errorsMiddleware, metricsMiddleware, panicMiddleware)

	checkapi.Routes(build, app, log, db, authClient)

	return app
}
