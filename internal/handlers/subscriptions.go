package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Subscription struct {
	ID          string `json:"id"`
	PlanID      string `json:"plan_id"`
	Customer    string `json:"customer"`
	Status      string `json:"status"`
	Amount      string `json:"amount"`
	Interval    string `json:"interval"`
	NextBilling string `json:"next_billing,omitempty"`
}

// ListSubscriptions handles GET /api/subscriptions with dynamic filtering
func ListSubscriptions(c *gin.Context) {
	// 1. Capture Query Parameters for filtering
	status := c.Query("status")      // e.g., ?status=active
	planID := c.Query("plan_id")     // e.g., ?plan_id=premium_01
	customer := c.Query("customer")   // e.g., ?customer=cust_123

	// 2. Mock data for demonstration (In production, this comes from your DB)
	allSubscriptions := []Subscription{
		{ID: "1", PlanID: "basic", Customer: "Alice", Status: "active", Amount: "10", Interval: "month"},
		{ID: "2", PlanID: "pro", Customer: "Bob", Status: "expired", Amount: "50", Interval: "month"},
	}

	// 3. Apply filtering logic
	var filtered []Subscription
	for _, s := range allSubscriptions {
		match := true

		if status != "" && s.Status != status {
			match = false
		}
		if planID != "" && s.PlanID != planID {
			match = false
		}
		if customer != "" && s.Customer != customer {
			match = false
		}

		if match {
			filtered = append(filtered, s)
		}
	}

	// 4. Return the filtered list
	c.JSON(http.StatusOK, gin.H{
		"subscriptions": filtered,
		"meta": gin.H{
			"total":  len(filtered),
			"status_filter": status,
		},
	})
}

func GetSubscription(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subscription id required"})
		return
	}
	
	// TODO: Replace with DB lookup: db.Where("id = ?", id).First(&subscription)
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"status": "placeholder",
	})
}