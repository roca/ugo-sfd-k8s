package web

import (
	"context"
	"errors"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/google/uuid"
)

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mw       []MidHandler
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...MidHandler) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}

// SignalShutdown is used to gracfully shut down the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// HandleFunc sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) HandleFunc(pattern string, handler Handler, mw ...MidHandler) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		v := Values{
			TraceID: uuid.NewString(),
			Now:     time.Now().UTC(),
		}
		ctx := setValues(r.Context(), &v)

		if err := handler(ctx, w, r); err != nil {
			if validateError(err) {
				a.SignalShutdown()
				return
			}
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}

func validateError(err error) bool {

	// Ignore the error if it is a known safe error.
	// syscall.EPIPE and syscall.ECONNRESET are known safe errors.
	// These are errors that can happen during a write to a connection
	// that has been closed on the other end. We don't care about these.

	switch {
	case errors.Is(err, syscall.EPIPE):
		return false
	case errors.Is(err, syscall.ECONNRESET):
		return false
	}

	return true
}
