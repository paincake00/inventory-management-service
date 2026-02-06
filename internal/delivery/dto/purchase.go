package dto

import "time"

type PurchaseCreateDTO struct {
	Name    string  `json:"name" binding:"required"`
	Count   float64 `json:"count" binding:"required"`
	Measure string  `json:"measure" binding:"required"`
	Market  string  `json:"market" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

type PurchaseUpdateDTO struct {
	ID      uint     `json:"-"`
	Name    *string  `json:"name"`
	Count   *float64 `json:"count"`
	Measure *string  `json:"measure"`
	Market  *string  `json:"market"`
	Price   *float64 `json:"price"`
	//Used       *int     `json:"used"`
	//Remainder  *int     `json:"remainder"`
}

type PutToArchivePurchase struct {
	InArchiveStatus bool `json:"is_archived"`
}

type PurchaseResponseDTO struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Count      float64   `json:"count"`
	Measure    string    `json:"measure"`
	Market     string    `json:"market"`
	Price      float64   `json:"price"`
	IsArchived bool      `json:"is_archived"`
	Used       int       `json:"used"`
	Remainder  int       `json:"remainder"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
