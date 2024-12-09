package middleware

import "net/http"

type Option func(CHandlerFunc) CHandlerFunc

func Chain(mw ...Option) func(CHandlerFunc) http.Handler {
	return func(handler CHandlerFunc) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next := handler
			for k := len(mw) - 1; k >= 0; k-- {
				currHandler := mw[k]
				nextHandler := next

				next = currHandler(nextHandler)
			}

			next(w, r) // nolint: errcheck
		})
	}
}
