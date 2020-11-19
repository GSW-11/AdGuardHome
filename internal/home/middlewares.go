package home

import (
	"net/http"
	"strings"

	"github.com/AdguardTeam/AdGuardHome/internal/aghio"
)

// middlerware is a wrapper function signature.
type middleware func(http.Handler) http.Handler

// withMiddlewares consequently wraps h with all the middlewares.
func withMiddlewares(h http.Handler, middlewares ...middleware) (wrapped http.Handler) {
	wrapped = h

	for _, mw := range middlewares {
		wrapped = mw(wrapped)
	}

	return wrapped
}

// RequestBodySizeLimit is maximum request body length in bytes.
const RequestBodySizeLimit = 64 * 1024

// LimitRequestBody substitutes body of the request with LimitedReadCloser.
func limitRequestBody(h http.Handler) (limited http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = aghio.LimitReadCloser(r.Body, RequestBodySizeLimit)

		h.ServeHTTP(w, r)
	})
}

// TODO(a.garipov): We currently have to use this, because everything registers
// its HTTP handlers in http.DefaultServeMux.  In the future, refactor our HTTP
// API initialization process and stop using the gosh darn http.DefaultServeMux
// for anything at all.  Gosh darn global variables.
func filterPProf(h http.Handler) (filtered http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/debug/pprof") {
			http.NotFound(w, r)

			return
		}

		h.ServeHTTP(w, r)
	})
}
