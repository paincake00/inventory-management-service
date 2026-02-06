package models

import "time"

type Purchase struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Count      float64   `gorm:"not null"`
	Measure    string    `gorm:"type:varchar(255);not null"`
	Market     string    `gorm:"type:varchar(255);not null"`
	Price      float64   `gorm:"not null"`
	IsArchived bool      `gorm:"not null"`
	Used       int       `gorm:"default:0;not null"`
	Remainder  int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
