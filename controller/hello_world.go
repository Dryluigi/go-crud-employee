package controller

import "net/http"

func NewHelloWorldController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}
}
