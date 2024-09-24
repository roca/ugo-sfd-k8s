package mid

import (
	"context"
	"net/http"

	"github.com/roca/ugo-sfd-k8s/app/api/mid"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// Metrics updates program counters using the middleware functionality.
func Metrics() web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Metrics(ctx, hdl)
		}

		return h
	}

	return m
}
