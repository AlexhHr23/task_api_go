package repositories

import (
	"gorm.io/gorm"
	"tast-list.com/models"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *ProjectRepository) CreateProject(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *ProjectRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *ProjectRepository) GetById(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *ProjectRepository) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *ProjectRepository) DeleteTask(id uint) error {
	return r.db.Delete(models.Task{}, "id = ?", id).Error
}
