package main

import (
	"fmt"
	"net/http"
)
xyz
type ServerType bool

func (m ServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	var k ServerType
	http.ListenAndServe("localhost:8000", k)
}
