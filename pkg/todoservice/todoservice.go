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
	t := todo.Task{Title: "Example TODO", Done: false}
	list := todo.List{}
	list.Tasks = make([]todo.Task, 1)
	list.Tasks[0] = t
	list.Created = time.Now()
	err := s.Stor.SaveList(&list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (s *TodoService) GetList(id string) (*todo.List, error) {
	list, err := s.Stor.GetList(id)
	return list, err
}
