package aghio

import (
	"net/http"
)

// LimitRequestBody substitutes body of the request with LimitedReadCloser.
func LimitRequestBody(h http.Handler) (limited http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = LimitReadCloser(r.Body, 1024)

		h.ServeHTTP(w, r)
	})
}
