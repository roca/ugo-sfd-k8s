// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/roca/ugo-sfd-k8s/apis/services/sales/route/sys/checkapi"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
