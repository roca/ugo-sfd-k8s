// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/roca/ugo-sfd-k8s/apis/services/api/mid"
	"github.com/roca/ugo-sfd-k8s/apis/services/sales/route/sys/checkapi"
	"github.com/roca/ugo-sfd-k8s/app/api/auth"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, auth *auth.Auth, shutdown chan os.Signal) *web.App {

	loggerMiddleware := mid.Logger(log)
	errorsMiddleware := mid.Errors(log)
	metricsMiddleware := mid.Metrics()
	panicMiddleware := mid.Panics()

	mux := web.NewApp(shutdown, loggerMiddleware, errorsMiddleware, metricsMiddleware, panicMiddleware)

	checkapi.Routes(mux, auth)

	return mux
}
