package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(
		writer http.ResponseWriter,
		request *http.Request) {
		fmt.Fprintln(writer,
			"<h1>Hello World!</h1>")
	})
}
