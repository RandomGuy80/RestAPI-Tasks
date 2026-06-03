package storage

import (
	"RestAPITasks/internal/domain"
	"fmt"
)

type Storage struct {
	Storage map[int]domain.Task
}

func NewStorage() *Storage {
	return &Storage{
		Storage: map[int]domain.Task{},
	}
}

func (s *Storage) Save(task domain.Task) {
	s.Storage[task.Id] = task
}

func (s *Storage) GetById(id int) (domain.Task, error) {
	task, exist := s.Storage[id]
	if !exist {
		return domain.Task{}, fmt.Errorf("task undefined")
	}

	return task, nil
}
