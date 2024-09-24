package checkapi

import "github.com/roca/ugo-sfd-k8s/foundation/web"

func Routes(app *web.App) {
	app.HandleFuncNoMiddleware("GET /liveliness", liveliness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
}
