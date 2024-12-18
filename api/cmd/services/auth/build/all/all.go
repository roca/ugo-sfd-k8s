// Package all binds all the routes into the specified app.
package all

import (
	"github.com/roca/ugo-sfd-k8s/api/http/api/mux"
	"github.com/roca/ugo-sfd-k8s/api/http/domain/authapi"
	"github.com/roca/ugo-sfd-k8s/api/http/domain/checkapi"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	authapi.Routes(app, authapi.Config{
		Auth: cfg.Auth,
	})
}
