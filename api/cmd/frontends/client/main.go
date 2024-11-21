package main

import (
	"context"
	"net/http"
	"os"

	"github.com/roca/ugo-sfd-k8s/api/cmd/frontends/client/handlers"
	"github.com/roca/ugo-sfd-k8s/foundation/logger"
	"github.com/roca/ugo-sfd-k8s/foundation/web"
)

var build = "develop"

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT *******")
		},
	}

	traceIDFn := func(ctx context.Context) string {
		return web.GetTraceID(ctx)
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES", traceIDFn, events)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	mux := routes()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return (err)
	}

	return nil
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Hello)
	mux.HandleFunc("/users", handlers.Users)
	return mux
}
