package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mode := gin.Mode()
	assert.Equal(t, mode, gin.TestMode)
}

func TestPingRoute(t *testing.T) {
	router := Create()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	res, _ := json.Marshal(map[string]string{
		"message": "pong",
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(res), w.Body.String())
}
