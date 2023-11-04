package repository

import "github.com/backendengineerark/routines-app/internal/domain/task/entity"

type ITaskRepository interface {
	Create(task *entity.Task) error
	Update(task *entity.Task) error
	ListsBy(isArchived bool) ([]entity.Task, error)
}
