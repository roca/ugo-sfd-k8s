package checkapi

import "github.com/roca/ugo-sfd-k8s/foundation/web"

func Routes(app *web.App) {
	app.HandleFunc("GET /liveliness", liveliness)
	app.HandleFunc("GET /readiness", readiness)
}
