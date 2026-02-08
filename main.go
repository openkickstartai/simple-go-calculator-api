package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CalcRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	Op string  `json:"op"`
}

type CalcResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "version": "1.0.0"})
	})
	
	r.POST("/calculate", calculateHandler)
	
	fmt.Println("Calculator API running on port 8080")
	log.Fatal(r.Run(":8080"))
}

func calculateHandler(c *gin.Context) {
	var req CalcRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CalcResponse{Error: "Invalid JSON format"})
		return
	}
	
	result, err := performCalculation(req.A, req.B, req.Op)
	if err != "" {
		c.JSON(http.StatusBadRequest, CalcResponse{Error: err})
		return
	}
	
	c.JSON(http.StatusOK, CalcResponse{Result: result})
}

func performCalculation(a, b float64, op string) (float64, string) {
	switch op {
	case "+":
		return a + b, ""
	case "-":
		return a - b, ""
	case "*":
		return a * b, ""
	case "/":
		if b == 0 {
			return 0, "Division by zero is not allowed"
		}
		return a / b, ""
	default:
		return 0, fmt.Sprintf("Unsupported operation: %s", op)
	}
}