// Pckage: mid provides app level middleware support.
package mid

import (
	"context"
)

// Handler represents the handler function that need to be called.
type Handler func(context.Context) error
