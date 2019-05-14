package main

import (
	"net/http"
)
func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello world"))
	})
	http.ListenAndServe(":3000", nil)
}