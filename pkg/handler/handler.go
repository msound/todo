package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msound/todo/pkg/db"
	"github.com/msound/todo/pkg/todo"
	"github.com/msound/todo/pkg/todoservice"
	"github.com/msound/todo/pkg/view"
	"github.com/rs/zerolog/log"
)

type App struct {
	S todo.Todoer
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
			log.Error().Err(err).Msg("Error getting list")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Info().Str("listID", listID).Msg("getting list")
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
		log.Info().Str("list_id", list.ID.Hex()).Msg("new list created")
	}

	view.Render(w, "index.html.tpl", list)
}

func (app *App) TaskDoneHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("list_id")
	if err != nil {
		log.Error().Err(err).Msg("List ID cookie missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	listID := cookie.Value
	vars := mux.Vars(r)
	taskID := vars["id"]
	task, err := app.S.TaskDone(listID, taskID)
	if err != nil {
		log.Error().Err(err).Str("listID", listID).Str("taskID", taskID).Msg("Error marking task as done")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Str("listID", listID).Str("taskID", taskID).Msg("Marking task as done")
	view.Render(w, "task.html.tpl", task)
}

func (app *App) TaskUndoHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("list_id")
	if err != nil {
		log.Error().Err(err).Msg("List ID cookie missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	listID := cookie.Value
	vars := mux.Vars(r)
	taskID := vars["id"]
	task, err := app.S.TaskUndo(listID, taskID)
	if err != nil {
		log.Error().Err(err).Str("listID", listID).Str("taskID", taskID).Msg("Error marking task as done")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Str("listID", listID).Str("taskID", taskID).Msg("Undo task")
	view.Render(w, "task.html.tpl", task)
}

func (app *App) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("list_id")
	if err != nil {
		log.Error().Err(err).Msg("List ID cookie missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	listID := cookie.Value

	task, err := app.S.AddTask(listID, r.FormValue("newtask"))
	if err != nil {
		log.Error().Err(err).Str("listID", listID).Msg("Cannot add task")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	view.Render(w, "add-task.html.tpl", task)
}
