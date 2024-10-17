package checkapi

import (
	"github.com/jmoiron/sqlx"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(build string, app *web.App, log *logger.Logger, db *sqlx.DB) {
	api := newAPI(build, log, db)
	app.HandleFuncNoMiddleware("GET /liveliness", api.liveliness)
	app.HandleFuncNoMiddleware("GET /readiness", api.readiness)
}
