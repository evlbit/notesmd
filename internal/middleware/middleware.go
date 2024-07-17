package middleware

import "net/http"

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

type Middleware func(http.Handler) http.Handler

func CreateStack(ms ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range ms {
			next = m(next)
		}

		return next
	}
}
