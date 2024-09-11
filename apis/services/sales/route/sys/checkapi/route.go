package checkapi

import "net/http"

func Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /liveliness", liveliness)
	mux.HandleFunc("GET /readiness", readiness)
}
