package models

import "gorm.io/gorm"

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusCancelled  Status = "cancelled"
)

type Task struct {
	gorm.Model
	Title         string  `gorm:"not null" json:"title"`
	Description   string  `gorm:"not null" json:"description"`
	Status        Status  `gorm:"type:enum('pending','in_progress','completed','cancelled');default:'pending'" json:"status"`
	ProjectID     uint    `gorm:"not null;index" json:"project_id"`
	Project       Project `json:"-"`
	CompletedByID *uint   `gorm:"index" json:"completed_by_id"`
	CompletedBy   *User   `json:"completed_by"`
}
