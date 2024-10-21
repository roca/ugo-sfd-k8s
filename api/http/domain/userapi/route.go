package userapi

import (
	"github.com/roca/ugo-sfd-k8s/api/http/api/mid"
	"github.com/roca/ugo-sfd-k8s/app/api/auth"
	"github.com/roca/ugo-sfd-k8s/app/api/authclient"
	"github.com/roca/ugo-sfd-k8s/app/domain/userapp"
	"github.com/roca/ugo-sfd-k8s/business/domain/userbus"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)
	ruleAuthorizeUser := mid.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOrSubject)
	ruleAuthorizeAdmin := mid.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOnly)

	api := newAPI(userapp.NewApp(cfg.UserBus))
	app.HandleFunc("GET /users", api.query, authen, ruleAdmin)
	app.HandleFunc("GET /users/{user_id}", api.queryByID, authen, ruleAuthorizeUser)
	app.HandleFunc("POST /users", api.create, authen, ruleAdmin)
	app.HandleFunc("PUT /users/role/{user_id}", api.updateRole, authen, ruleAuthorizeAdmin)
	app.HandleFunc("PUT /users/{user_id}", api.update, authen, ruleAuthorizeUser)
	app.HandleFunc("DELETE /users/{user_id}", api.delete, authen, ruleAuthorizeUser)
}
