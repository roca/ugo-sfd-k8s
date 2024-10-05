package mid

import (
	"context"
	"net/http"

	"github.com/roca/ugo-sfd-k8s/app/api/auth"
	"github.com/roca/ugo-sfd-k8s/app/api/mid"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// Authorize executes the authorize middleware functionality.
func Authorize(auth *auth.Auth, rule string) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Authorize(ctx, auth, rule, hdl)
		}

		return h
	}

	return m
}
