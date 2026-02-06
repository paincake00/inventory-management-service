package models

import "time"

type Material struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Count     int       `gorm:"not null"`
	Measure   string    `gorm:"type:varchar(255);not null"`
	Cost      float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	PurchaseID uint
	Purchase   *Purchase `gorm:"constraint:OnDelete:RESTRICT;"` // belongs to
	CustomID   uint
	Custom     *Custom `gorm:"constraint:OnDelete:CASCADE;"` // belongs to
}
