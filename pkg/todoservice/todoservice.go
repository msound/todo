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
	list.Tasks = make([]todo.Task, 0)
	list.Created = time.Now()
	err := s.Stor.SaveList(&list)
	if err != nil {
		return nil, err
	}
	err = s.AddTask(list.ID.Hex(), "Example TODO")
	if err != nil {
		return nil, err
	}

	return s.GetList(list.ID.Hex())
}

func (s *TodoService) AddTask(listID string, task string) error {
	t := todo.Task{ID: s.Stor.GetNewID(), Title: task, Done: false, Created: time.Now()}
	return s.Stor.AddTask(listID, t)
}

func (s *TodoService) GetList(id string) (*todo.List, error) {
	list, err := s.Stor.GetList(id)
	return list, err
}
