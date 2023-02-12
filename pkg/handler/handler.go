package handler

import (
	"fmt"
	"net/http"

	"github.com/msound/todo/pkg/db"
	"github.com/msound/todo/pkg/todo"
	"github.com/msound/todo/pkg/todoservice"
	"github.com/msound/todo/pkg/view"
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
	} else {
		// Create a new list, this is a new visitor
		list, err = app.S.NewList()
		cookie := http.Cookie{Name: "list_id", Value: list.ID}
		http.SetCookie(w, &cookie)
	}
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	view.Render(w, "index.html", list)
}
