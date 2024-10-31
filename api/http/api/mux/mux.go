// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/roca/ugo-sfd-k8s/api/http/api/mid"
	"github.com/roca/ugo-sfd-k8s/app/api/auth"
	"github.com/roca/ugo-sfd-k8s/app/api/authclient"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build      string
	Log        *logger.Logger
	Auth       *auth.Auth
	AuthClient *authclient.Client
	DB         *sqlx.DB
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg Config)
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config, routeAdder RouteAdder) *web.App {
	logger := func(ctx context.Context, msg string, v ...any) {
		cfg.Log.Info(ctx, msg, v...)
	}

	app := web.NewApp(
		logger,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
	)

	routeAdder.Add(app, cfg)

	return app
}
