package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "version": "1.0.0"})
	})
	
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", response["status"])
	assert.Equal(t, "1.0.0", response["version"])
}

func TestAddition(t *testing.T) {
	result, err := performCalculation(10, 5, "+")
	assert.Equal(t, 15.0, result)
	assert.Equal(t, "", err)
}

func TestSubtraction(t *testing.T) {
	result, err := performCalculation(10, 3, "-")
	assert.Equal(t, 7.0, result)
	assert.Equal(t, "", err)
}

func TestMultiplication(t *testing.T) {
	result, err := performCalculation(6, 7, "*")
	assert.Equal(t, 42.0, result)
	assert.Equal(t, "", err)
}

func TestDivision(t *testing.T) {
	result, err := performCalculation(20, 4, "/")
	assert.Equal(t, 5.0, result)
	assert.Equal(t, "", err)
}

func TestDivisionByZero(t *testing.T) {
	result, err := performCalculation(10, 0, "/")
	assert.Equal(t, 0.0, result)
	assert.Equal(t, "Division by zero is not allowed", err)
}

func TestUnsupportedOperation(t *testing.T) {
	result, err := performCalculation(10, 5, "%")
	assert.Equal(t, 0.0, result)
	assert.Equal(t, "Unsupported operation: %", err)
}