package usecase

import (
	"RestAPITasks/internal/domain"
	"fmt"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Save(task domain.Task) error {
	if task.Title == "" {
		return fmt.Errorf("Title can't be empty ")
	}
	if task.Id <= 0 {
		return fmt.Errorf("Id must be positive ")
	}
	u.repo.Save(task)
	return nil
}

func (u *UserUsecase) GetTaskById(id int) (domain.Task, error) {
	return u.repo.GetById(id)
}
