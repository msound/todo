package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
	"github.com/msound/todo/pkg/db"
	"github.com/msound/todo/pkg/handler"
	"github.com/msound/todo/pkg/view"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Info().Msg("Starting TODO")
	view.Load()
	dbClient, err := db.NewClient(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}
	app := handler.NewApp(dbClient)
	r := mux.NewRouter()
	r.HandleFunc("/", app.IndexHandler).Methods("GET")
	r.HandleFunc("/task/{id}/done", app.TaskDoneHandler).Methods("POST")
	r.HandleFunc("/task/{id}/undo", app.TaskUndoHandler).Methods("POST")
	exit := http.ListenAndServe(":8000", r)
	log.Fatal().Err(exit)
}
