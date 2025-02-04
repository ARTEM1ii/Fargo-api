package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	FullName    string `json:"full_name" gorm:"size:100;not null"`
	Phone       string `json:"phone" gorm:"size:15;unique;not null"`
	SecondPhone string `json:"second_phone" gorm:"size:15"`
	Email       string `json:"email" gorm:"size:100;unique"`
	UniqueCode  string `json:"unique_code" gorm:"size:10;unique;not null"`
}