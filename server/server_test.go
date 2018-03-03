package server

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
