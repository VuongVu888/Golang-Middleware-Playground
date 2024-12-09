package middleware

import (
	"fmt"
	"log"
	"net/http"

	cerror "playground/middleware/error"
	"playground/middleware/middleware/common"
)

type Middleware interface {
	FirstMiddleware(handler common.HandlerWithError) common.HandlerWithError
	SecondMiddleware(handler common.HandlerWithError) common.HandlerWithError
}

type middleware struct {
	logger *log.Logger
}

func NewMiddleware(logger *log.Logger) *middleware {
	return &middleware{
		logger: logger,
	}
}

func (m *middleware) FirstChainingMiddleware(handler common.HandlerWithError) common.HandlerWithError {
	return common.HandlerWithError(func(w http.ResponseWriter, req *http.Request) error {
		wrapper := common.NewResponseWriter(w)

		if err := handler(wrapper, req); err != nil {
			if cerr, ok := cerror.As(err); ok {
				http.Error(w, cerr.Error(), int(cerr.StatusCode().HTTPStatusCode()))
				return cerr
			}
			fmt.Println("unknown error handled: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		w.WriteHeader(wrapper.StatusCode())
		if _, err := w.Write(wrapper.Body().Bytes()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte("Hello from first middleware!\n"))

		return nil
	})
}

func (m *middleware) SecondChainingMiddleware(handler common.HandlerWithError) common.HandlerWithError {
	return common.HandlerWithError(func(w http.ResponseWriter, req *http.Request) error {
		wrapper := common.NewResponseWriter(w)

		if err := handler(wrapper, req); err != nil {
			if cerr, ok := cerror.As(err); ok {
				http.Error(w, cerr.Error(), int(cerr.StatusCode().HTTPStatusCode()))

				return cerr
			}
			fmt.Println("unknown error handled: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return nil
		}

		w.WriteHeader(wrapper.StatusCode())
		if _, err := w.Write(wrapper.Body().Bytes()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte("Hello from second middleware!\n"))

		return nil
	})
}
