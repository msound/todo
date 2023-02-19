package todoservice

import (
	"time"

	"github.com/msound/todo/pkg/todo"
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
