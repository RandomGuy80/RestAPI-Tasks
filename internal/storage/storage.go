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

func (s *Storage) DeleteById(id int) error {
	if _, exist := s.Storage[id]; !exist {
		return fmt.Errorf("task with id: %d don't exists", id)
	}
	delete(s.Storage, id)
	return nil
}

func (s *Storage) UpdateById(id int, task domain.Task) error {
	taskFirst, exist := s.Storage[id]
	if !exist {
		return fmt.Errorf("task with id: %d don't exists", id)
	}
	if taskFirst.Id != task.Id {
		return fmt.Errorf("current task and previous task must have same id")
	}
	s.Storage[id] = task
	return nil
}
