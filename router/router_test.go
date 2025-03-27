package router

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
