package main

import (
	"log"
	"net/http"

	"playground/middleware/handler"
	middleware2 "playground/middleware/middleware/chaining"
	middleware1 "playground/middleware/middleware/simple"
)

func SimpleMiddlewareHandler() http.Handler {
	return middleware1.SimpleMiddleware(handler.Handler)
}

func ChainingMiddlewareHandler() http.Handler {
	logger := &log.Logger{}
	m := middleware2.NewMiddleware(logger)

	middlewares := []middleware2.Option{
		m.SecondChainingMiddleware,
		m.FirstChainingMiddleware,
	}

	return middleware2.Chain(middlewares...)(handler.Handler)
}

func main() {
	simpleHandler := SimpleMiddlewareHandler()
	chainingHandler := ChainingMiddlewareHandler()

	http.Handle("/simple", simpleHandler)
	http.Handle("/chain", chainingHandler)
	http.ListenAndServe(":8080", nil)
}
