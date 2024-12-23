package middleware

import (
	"net/http"

	"playground/middleware/middleware/common"
)

type Option func(common.HandlerWithError) common.HandlerWithError

func Chain(mw ...Option) func(common.HandlerWithError) http.Handler {
	return func(handler common.HandlerWithError) http.Handler {
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
