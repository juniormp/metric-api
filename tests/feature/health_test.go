package feature_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/juniormp/metric-api/src/web"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestHealthConnection(t *testing.T) {
	router := web.SetupRouter()

	response := performRequest(router, "GET", "/health")
	var body map[string]string
	err := json.Unmarshal([]byte(response.Body.String()), &body)
	value, exists := body["status"]

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "live!", value)
}
