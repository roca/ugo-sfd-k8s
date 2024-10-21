package mid

import (
	"context"
	"net/http"

	"github.com/roca/ugo-sfd-k8s/app/api/authclient"
	"github.com/roca/ugo-sfd-k8s/app/api/mid"
	"github.com/roca/ugo-sfd-k8s/business/domain/userbus"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

// Authorize executes the authorize middleware functionality.
func Authorize(log *logger.Logger, client *authclient.Client, rule string) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.Authorize(ctx, log, client, rule, hdl)
		}

		return h
	}

	return m
}

// AuthorizeUser executes the specified role and extracts the specified
// user from the DB if a user id is specified in the call. Depending on the rule
// specified, the userid from the claims may be compared with the specified
// user id.
func AuthorizeUser(log *logger.Logger, client *authclient.Client, userBus *userbus.Business, rule string) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.AuthorizeUser(ctx, log, client, userBus, rule, web.Param(r, "user_id"), hdl)
		}

		return h
	}

	return m
}
