package routes

import (
	"be-go-music-store/controllers"
	"be-go-music-store/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.GET("/products", controllers.GetProducts)

	// Protected routes
	auth := r.Group("/admin")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/products", controllers.CreateProduct)
		auth.PUT("/products/:id", controllers.UpdateProduct)
		auth.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
