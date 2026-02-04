package app

import "github.com/gin-gonic/gin"

func (app *App) InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		purchases := v1.Group("/purchases")
		{
			purchases.POST("")
			purchases.GET("")
			purchases.GET("/:purchaseID")
			purchases.PUT("/:purchaseID")
			purchases.DELETE("/:purchaseID")
		}
		customs := v1.Group("/customs")
		{
			customs.POST("")
			customs.GET("")
			customs.GET("/:customID")
			customs.PUT("/:customID")
			customs.DELETE("/:customID")

			materials := customs.Group("/:customID/materials")
			{
				materials.POST("")
				materials.GET("")
				materials.GET("/:materialID")
				materials.PUT("/:materialID")
				materials.DELETE("/:materialID")
			}
		}
		clients := v1.Group("/clients")
		{
			clients.POST("")
			clients.GET("")
			clients.GET("/:clientID")
			clients.PUT("/:clientID")
			clients.DELETE("/:clientID")

			customsForClient := clients.Group("/:clientID/customs")
			{
				customsForClient.POST("")
				customsForClient.GET("")
				customsForClient.DELETE("/:customID")
			}
		}
	}

	return router
}
