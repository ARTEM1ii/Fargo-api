package models

import (
	"gorm.io/gorm"
)

type TrackStatus string

const (
	InTransit         TrackStatus = "📦 В пути"
	AtChinaWarehouse TrackStatus = "📦 На складе Китая"
	AtDushanbeWarehouse TrackStatus = "📦 На складе Душанбе"
	Delivered        TrackStatus = "✅ Доставлено"
	Cancelled        TrackStatus = "❌ Отменено"
)

type TrackCode struct {
	gorm.Model
	ClientID  string      `json:"client_id" gorm:"size:10;not null"`
	TrackCode string      `json:"track_code" gorm:"size:20;not null;unique"`
	Status    TrackStatus `json:"status" gorm:"type:varchar(50);not null;default:'📦 В ожидании'"`
}