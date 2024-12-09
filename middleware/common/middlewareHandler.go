package common

type HandlerWithError func(w http.ResponseWriter, req *http.Request) error