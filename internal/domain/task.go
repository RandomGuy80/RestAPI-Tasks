package domain

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type UserRepository interface {
	Save(Task) error
	GetById(id int) (Task, error)
	DeleteById(id int) error
	UpdateById(id int, task Task) error
}

type UserUsecase interface {
	Save(Task) error
	GetTaskById(id int) (Task, error)
	DeleteById(id int) error
	UpdateById(id int, task Task) error
}
