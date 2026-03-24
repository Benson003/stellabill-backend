package routes

import (
	"github.com/gin-gonic/gin"
	"stellarbill-backend/internal/handlers"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)
		api.GET("/subscriptions", handlers.ListSubscriptions)
		api.GET("/subscriptions/:id", handlers.GetSubscription)
		api.GET("/plans", handlers.ListPlans)
	}
}
