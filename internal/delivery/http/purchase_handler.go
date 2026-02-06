package http

import (
	"github.com/gin-gonic/gin"
	"github.com/paincake00/inventory-management-service/internal/domain/service"
	"go.uber.org/zap"
)

type PurchaseHandler struct {
	purchaseService service.IPurchaseService
	logger          *zap.SugaredLogger
}

func NewPurchaseHandler(logger *zap.SugaredLogger, purchaseService service.IPurchaseService) *PurchaseHandler {
	return &PurchaseHandler{
		purchaseService: purchaseService,
		logger:          logger,
	}
}

func (h *PurchaseHandler) CreatePurchase(c *gin.Context) {

}

func (h *PurchaseHandler) GetAllPurchases(c *gin.Context) {}

func (h *PurchaseHandler) GetPurchaseByID(c *gin.Context) {}

func (h *PurchaseHandler) UpdatePurchaseByID(c *gin.Context) {}

func (h *PurchaseHandler) DeletePurchaseByID(c *gin.Context) {}
