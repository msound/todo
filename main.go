package main

import (
	"fmt"
	"net/http"

	"github.com/msound/todo/pkg/handler"
	"github.com/msound/todo/pkg/view"
)

func main() {
	fmt.Println("TODO")
	view.Load()
	http.HandleFunc("/", handler.IndexHandler)
	http.ListenAndServe(":8000", nil)
}
