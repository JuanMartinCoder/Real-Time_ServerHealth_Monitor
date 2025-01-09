package main

import "net/http"

func HandlerMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
