package models

import "gorm.io/gorm"

type CompanyContact struct {
	gorm.Model
	Country      string `json:"country" gorm:"size:50;not null"`
	Phone        string `json:"phone" gorm:"size:15;not null"`
	Email        string `json:"email" gorm:"size:100"`
	WorkingHours string `json:"working_hours" gorm:"size:50"`
	Address      string `json:"address" gorm:"not null"`
}

