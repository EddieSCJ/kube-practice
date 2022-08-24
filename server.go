package main

import (
	"os"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	value := os.Getenv("MESSAGE")
	if value == "" {
		value = "Hello World"
	}
	w.Write([]byte(value))
}