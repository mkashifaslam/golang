package main

import (
	"fmt"
	"net/http"
)

func main() {
	_ = http.ListenAndServe(":9000", http.HandlerFunc(hello))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	//_, _ = fmt.Fprintf(w, "hello world")
}
