package mid

import (
	"context"
	"errors"

	"github.com/roca/ugo-sfd-k8s/app/api/authclient"
	"github.com/roca/ugo-sfd-k8s/app/api/errs"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
)

// ErrInvalidID represents a condition where the id is not a uuid.
var ErrInvalidID = errors.New("ID is not in its proper form")

// Authorize executes the specified role and does not extract any domain data.
func Authorize(ctx context.Context, log *logger.Logger, client *authclient.Client, rule string, handler Handler) error {
	userID, err := GetUserID(ctx)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	auth := authclient.Authorize{
		Claims: GetClaims(ctx),
		UserID: userID,
		Rule:   rule,
	}

	if err := client.Authorize(ctx, auth); err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	return handler(ctx)
}
