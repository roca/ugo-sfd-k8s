package checkapi

import (
	"github.com/roca/ugo-sfd-k8s/apis/services/api/mid"
	"github.com/roca/ugo-sfd-k8s/app/api/auth"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

func Routes(app *web.App, a *auth.Auth) {
	authen := mid.Authorization(a)
	athAdminOnly := mid.Authorize(a, auth.RuleAdminOnly)

	app.HandleFuncNoMiddleware("GET /liveliness", liveliness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
	app.HandleFunc("GET /testauth", liveliness, authen, athAdminOnly)
}
