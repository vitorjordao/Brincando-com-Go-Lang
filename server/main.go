package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8887", nil)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
