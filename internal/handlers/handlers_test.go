package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Health(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body["status"] != "ok" || body["service"] != "stellarbill-backend" {
		t.Fatalf("unexpected body: %v", body)
	}
}

func TestListPlans(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ListPlans(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string][]Plan
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if plans := body["plans"]; len(plans) != 0 {
		t.Fatalf("expected empty plans, got %v", plans)
	}
}

func TestListSubscriptions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ListSubscriptions(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string][]Subscription
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if subscriptions := body["subscriptions"]; len(subscriptions) != 0 {
		t.Fatalf("expected empty subscriptions, got %v", subscriptions)
	}
}

func TestGetSubscription(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("missing id", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		GetSubscription(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", w.Code)
		}
	})

	t.Run("with id", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "sub_123"}}

		GetSubscription(c)

		if w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", w.Code)
		}

		var body map[string]string
		if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body["id"] != "sub_123" || body["status"] != "placeholder" {
			t.Fatalf("unexpected body: %v", body)
		}
	})
}
