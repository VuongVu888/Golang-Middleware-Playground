package common

import "net/http"

type HandlerWithError func(w http.ResponseWriter, req *http.Request) error
