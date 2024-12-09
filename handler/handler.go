package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) error {
	testString := []byte("This is Go's handler!\n")
	w.Write(testString)
	return nil
}
