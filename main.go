package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/msound/todo/pkg/db"
	"github.com/msound/todo/pkg/handler"
	"github.com/msound/todo/pkg/view"
)

func main() {
	fmt.Println("TODO")
	view.Load()
	dbClient, err := db.NewClient(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}
	app := handler.NewApp(dbClient)
	r := mux.NewRouter()
	r.HandleFunc("/", app.IndexHandler).Methods("GET")
	http.ListenAndServe(":8000", r)
}
