package infra

import (
	"ddd-go/domain/model"
	"ddd-go/domain/repository"

	"github.com/jinzhu/gorm"
)

// taskRepositoryの構造体
type TaskRepository struct {
	Conn *gorm.DB
}

func (t TaskRepository) Create(task *model.Task) (*model.Task, error) {
	panic("implement me")
}

func (t TaskRepository) FindByID(id int) (*model.Task, error) {
	panic("implement me")
}

func (t TaskRepository) Update(task *model.Task) (*model.Task, error) {
	panic("implement me")
}

func (t TaskRepository) Delete(id int) error {
	panic("implement me")
}

// コンストラクタ
func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}
