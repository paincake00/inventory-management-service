package models

import "time"

type Custom struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	ImageURI    string `gorm:"type:varchar(255);"`
	SelfCost    float64
	Cost        float64
	Description string
	Used        int       `gorm:"default:0;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Materials []Material `gorm:"foreignKey:CustomID"` // has many
	// Clients   []Client   `gorm:"many2many:client_customs"` // many to many (back reference) - EXTRA
}
