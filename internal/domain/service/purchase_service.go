package service

import (
	"context"

	"github.com/paincake00/inventory-management-service/internal/infrastructure/models"
)

type IPurchaseService interface {
	Create(ctx context.Context, purchase *models.Purchase) (*models.Purchase, error)
	GetAll(ctx context.Context, limit, offset int) ([]models.Purchase, error)
	GetByID(ctx context.Context, id uint) (*models.Purchase, error)
	UpdateByID(ctx context.Context, purchase map[string]interface{}) (*models.Purchase, error)
	PutToArchive(ctx context.Context, purchaseID uint, inArchiveStatus bool) (*models.Purchase, error)
	IncrementUsed(ctx context.Context, purchaseID uint) error
	DecrementUsed(ctx context.Context, purchaseID uint) error
	ChangeRemainder(ctx context.Context, purchaseID uint, diff float64) error
	DeleteByID(ctx context.Context, id uint) error
}
