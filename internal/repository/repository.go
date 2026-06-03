package repository

import (
	"RestAPITasks/internal/domain"
	"RestAPITasks/internal/storage"
)

type UserRepository struct {
	db *storage.Storage
}

func NewUserRepository(db *storage.Storage) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(task domain.Task) {
	r.db.Save(task)
}

func (r *UserRepository) GetById(id int) (task domain.Task, err error) {
	return r.db.GetById(id)
}
