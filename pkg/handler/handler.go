package handler

import (
	"net/http"

	"github.com/msound/todo/pkg/db"
	"github.com/msound/todo/pkg/todo"
	"github.com/msound/todo/pkg/todoservice"
	"github.com/msound/todo/pkg/view"
	"github.com/rs/zerolog/log"
)

type App struct {
	S *todoservice.TodoService
}

func NewApp(dbClient *db.Client) *App {
	var app App
	app.S = todoservice.NewTodoService(dbClient)
	return &app
}

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var list *todo.List
	cookie, err := r.Cookie("list_id")
	if err == nil {
		// Get existing list
		listID := cookie.Value
		list, err = app.S.GetList(listID)
		if err != nil {
			log.Error().Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		// Create a new list, this is a new visitor
		log.Debug().Msg("new visitor")
		list, err = app.S.NewList()
		if err != nil {
			log.Error().Err(err).Msg("Error creating new list")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cookie := http.Cookie{Name: "list_id", Value: list.ID.Hex()}
		http.SetCookie(w, &cookie)
		log.Debug().Str("list_id", list.ID.Hex()).Msg("new list created")
	}

	view.Render(w, "index.html", list)
}
