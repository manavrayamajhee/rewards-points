package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"gin-docker-app/models"
)

func setupRouter() *gin.Engine {
	// Setup Gin router for testing
	router := gin.Default()
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getReceiptPoints)
	return router
}

// Test the POST /receipts/process endpoint
func TestProcessReceipt(t *testing.T) {
	router := setupRouter()

	// Create a test receipt
	receipt := models.Receipt{
		Retailer:     "Retailer1",
		PurchaseDate: "2024-11-10",
		PurchaseTime: "15:30",
		Items: []models.Item{
			{ShortDescription: "Item 1", Price: "10.5"},
			{ShortDescription: "Item 2", Price: "5.5"},
		},
		Total: "16",
	}

	// Convert receipt to JSON
	jsonValue, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("Failed to marshal receipt: %v", err)
	}

	// Send POST request
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the status code is OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response JSON to get the receipt ID
	var receiptResponse models.ReceiptResponse
	err = json.Unmarshal(w.Body.Bytes(), &receiptResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Assert that the ID is not empty
	assert.NotEmpty(t, receiptResponse.ID)
}

// Test the GET /receipts/:id/points endpoint
func TestGetReceiptPoints(t *testing.T) {
	router := setupRouter()

	// Create a test receipt
	receipt := models.Receipt{
		Retailer:     "Retailer1",
		PurchaseDate: "2024-11-10",
		PurchaseTime: "15:30",
		Items: []models.Item{
			{ShortDescription: "Item 1", Price: "10.5"},
			{ShortDescription: "Item 2", Price: "5.5"},
		},
		Total: "16",
	}

	// Simulate saving the receipt to the in-memory database
	receiptID := "some-unique-id"
	receiptDB.Store(receiptID, receipt)

	// Send GET request
	req, _ := http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the status code is OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response JSON to get the points
	var pointsResponse models.PointsResponse
	err := json.Unmarshal(w.Body.Bytes(), &pointsResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Assert that points are returned
	assert.Greater(t, pointsResponse.Points, 0)
}
