package routes

import (
	"github.com/gin-gonic/gin"
	"stellarbill-backend/internal/handlers"
)

func Register(r *gin.Engine) {
	// Apply global middleware
	r.Use(corsMiddleware())

	// Define the API version/group
	api := r.Group("/api")
	{
		// Health check for monitoring
		api.GET("/health", handlers.Health)

		// Subscription Management (Now with filtering support)
		api.GET("/subscriptions", handlers.ListSubscriptions)
		api.GET("/subscriptions/:id", handlers.GetSubscription)

		// Plan Management
		api.GET("/plans", handlers.ListPlans)
	}
}

// corsMiddleware handles Cross-Origin Resource Sharing for the frontend
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}