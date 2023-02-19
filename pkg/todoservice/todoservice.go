package todoservice

import (
	"errors"
	"time"

	"github.com/msound/todo/pkg/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	Stor todo.Storage
}

func NewTodoService(s todo.Storage) *TodoService {
	todoService := TodoService{Stor: s}
	return &todoService
}

func (s *TodoService) NewList() (*todo.List, error) {
	list := todo.List{}
	list.Tasks = make([]todo.Task, 5)
	list.Tasks[0] = todo.Task{ID: s.Stor.GetNewID(), Title: "Write a todo list handler", Done: true, Created: time.Now()}
	list.Tasks[1] = todo.Task{ID: s.Stor.GetNewID(), Title: "Add support for database", Done: true, Created: time.Now()}
	list.Tasks[2] = todo.Task{ID: s.Stor.GetNewID(), Title: "Refactor the code", Done: true, Created: time.Now()}
	list.Tasks[3] = todo.Task{ID: s.Stor.GetNewID(), Title: "Throw HTMX into the mix", Done: false, Created: time.Now()}
	list.Tasks[4] = todo.Task{ID: s.Stor.GetNewID(), Title: "Et voilà", Done: false, Created: time.Now()}
	list.Created = time.Now()
	err := s.Stor.SaveList(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
	// err = s.AddTask(list.ID.Hex(), "Example TODO")
	// if err != nil {
	// 	return nil, err
	// }

	// return s.GetList(list.ID.Hex())
}

func (s *TodoService) AddTask(listID string, task string) error {
	t := todo.Task{ID: s.Stor.GetNewID(), Title: task, Done: false, Created: time.Now()}
	return s.Stor.AddTask(listID, t)
}

func (s *TodoService) GetList(id string) (*todo.List, error) {
	list, err := s.Stor.GetList(id)
	return list, err
}

func (s *TodoService) TaskDone(listID string, taskID string) (*todo.Task, error) {
	return s.markTaskAsDone(listID, taskID, true)
}

func (s *TodoService) TaskUndo(listID string, taskID string) (*todo.Task, error) {
	return s.markTaskAsDone(listID, taskID, false)
}

func (s *TodoService) markTaskAsDone(listID string, taskID string, done bool) (*todo.Task, error) {
	var output todo.Task

	err := s.Stor.TaskDone(listID, taskID, done)
	if err != nil {
		return nil, err
	}

	list, err := s.Stor.GetList(listID)
	if err != nil {
		return nil, err
	}

	taskOID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, err
	}

	found := false
	for _, task := range list.Tasks {
		if task.ID == taskOID {
			output = task
			found = true
		}
	}

	if !found {
		return nil, errors.New("task is missing in list")
	}

	return &output, nil
}
