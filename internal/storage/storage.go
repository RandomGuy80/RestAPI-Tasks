package storage

import (
	"RestAPITasks/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(connStr string) (*Storage, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Save(task domain.Task) error {
	_, err := s.db.Exec("INSERT INTO tasks (id, title, done) VALUES ($1, $2, $3)", task.Id, task.Title, task.Done)
	return err
}

func (s *Storage) GetById(id int) (domain.Task, error) {
	var task domain.Task
	err := s.db.Get(&task, "SELECT id, title, done FROM tasks WHERE id=$1", id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("task with id: %d not found", id)
	}
	return task, nil
}

func (s *Storage) DeleteById(id int) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("tasks with id: %d not exists", id)
	}
	return nil
}

func (s *Storage) UpdateById(id int, task domain.Task) error {
	_, err := s.db.Exec("UPDATE tasks SET title = $1, done = $2 WHERE id = $3", task.Title, task.Done, id)
	if err != nil {
		return fmt.Errorf("task with id: %d not found", id)
	}
	return nil
}
