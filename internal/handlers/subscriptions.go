package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"stellabill-backend/internal/subscriptions"
)

type Subscription struct {
	ID        string `json:"id"`
	PlanID    string `json:"plan_id"`
	Customer  string `json:"customer"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	Interval  string `json:"interval"`
	NextBilling string `json:"next_billing,omitempty"`
}

func ListSubscriptions(c *gin.Context) {
	// TODO: load from DB, filter by merchant from JWT/API key
	subscriptions := []Subscription{}
	c.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
}

func GetSubscription(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subscription id required"})
		return
	}
	// TODO: load from DB by id
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"status": "placeholder",
	})
}

// UpdateSubscriptionStatus handles status updates with validation
func UpdateSubscriptionStatus(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subscription id required"})
		return
	}

	var payload struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: fetch current subscription from DB
	currentStatus := "active" // placeholder

	if err := subscriptions.CanTransition(currentStatus, payload.Status); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: persist update

	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"status": payload.Status,
	})
}
