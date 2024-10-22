// Package all binds all the routes into the specified app.
package all

import (
	"github.com/roca/ugo-sfd-k8s/api/http/api/mux"
	"github.com/roca/ugo-sfd-k8s/api/http/domain/checkapi"
	"github.com/roca/ugo-sfd-k8s/api/http/domain/testapi"
	"github.com/roca/ugo-sfd-k8s/api/http/domain/userapi"
	"github.com/roca/ugo-sfd-k8s/business/domain/userbus"
	"github.com/roca/ugo-sfd-k8s/business/domain/userbus/stores/userdb"
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
	userBus := userbus.NewBusiness(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB))

	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	testapi.Routes(app, testapi.Config{
		Log:        cfg.Log,
		AuthClient: cfg.AuthClient,
	})

	userapi.Routes(app, userapi.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		AuthClient: cfg.AuthClient,
	})
}