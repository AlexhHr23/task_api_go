package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	ManagerID   uint   `gorm:"not null;index" json:"manager_id"`
	Manager     User   `json:"manager"`
	Tasks       []Task `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"tasks"`
	Team        []User `gorm:"many2many:project_members;" json:"team"`
}
