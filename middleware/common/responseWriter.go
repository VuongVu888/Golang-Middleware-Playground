package common

import (
	"bytes"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.statusCode = code
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, http.StatusOK, &bytes.Buffer{}}
}

func (w *ResponseWriter) StatusCode() int {
	return w.statusCode
}

func (w *ResponseWriter) Body() *bytes.Buffer {
	return w.body
}
