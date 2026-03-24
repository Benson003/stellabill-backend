package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterExposesAPIEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	Register(router)

	cases := []struct {
		path string
		key  string
	}{
		{path: "/api/health", key: "status"},
		{path: "/api/plans", key: "plans"},
		{path: "/api/subscriptions", key: "subscriptions"},
		{path: "/api/subscriptions/sub_123", key: "id"},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("%s expected 200, got %d", tc.path, res.Code)
		}

		var body map[string]any
		if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
			t.Fatalf("%s decode body: %v", tc.path, err)
		}
		if _, ok := body[tc.key]; !ok {
			t.Fatalf("%s expected key %q in response body", tc.path, tc.key)
		}
	}
}
