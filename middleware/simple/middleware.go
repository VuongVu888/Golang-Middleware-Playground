package middleware

import (
	"fmt"
	"net/http"

	cerror "playground/middleware/error"
	"playground/middleware/middleware/common"
)

type HandlerWithError func(w http.ResponseWriter, req *http.Request) error

func SimpleMiddleware(handler HandlerWithError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		wrapper := common.NewResponseWriter(w)
		if err := handler(wrapper, req); err != nil {
			if cerr, ok := cerror.As(err); ok {
				http.Error(w, cerr.Error(), int(cerr.StatusCode().HTTPStatusCode()))
				return
			}
			fmt.Println("unknown error handled: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(wrapper.StatusCode())
		if _, err := w.Write(wrapper.Body().Bytes()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte("Hello from Simple Middleware!\n"))
	})
}
