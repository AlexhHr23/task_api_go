package repositories

import (
	"gorm.io/gorm"
	"tast-list.com/models"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetById(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) DeleteTask(id uint) error {
	return r.db.Delete(models.Task{}, "id = ?", id).Error
}
