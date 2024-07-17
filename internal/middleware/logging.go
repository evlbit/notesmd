package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		start := time.Now()

		next.ServeHTTP(wrappedWriter, r)

		log.Println(wrappedWriter.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
