package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "sync"
    "unicode"
    "strconv"
    "strings"
    "math"
    "gin-docker-app/parser"
    "gin-docker-app/models"
)

// In-memory database to store receipts using sync.Map for concurrency safety
var receiptDB sync.Map

// processReceipt handles the submission of a new receipt
func ProcessReceipt(c *gin.Context) {
    var receipt models.Receipt

    // Bind JSON payload to the Receipt struct and validate required fields
    if err := c.ShouldBindJSON(&receipt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
        return
    }

    // Generate a unique ID for the receipt
    receiptID := uuid.New().String()

    // Store the receipt in the sync.Map in-memory database
    receiptDB.Store(receiptID, receipt)

    // Return the ID in a JSON response
    c.JSON(http.StatusOK, models.ReceiptResponse{ID: receiptID})
}

// getReceiptPoints retrieves the points for a receipt by ID
func GetReceiptPoints(c *gin.Context) {
    receiptID := c.Param("id")

    // Parse receipt data and handle any errors
    parsedData := parser.ParseReceiptData(c, receiptID, &receiptDB)
    if !parsedData.Valid {
        return // Exit if parsing fails, error handling has already been done in the function
    }

    
    receipt := parsedData.Receipt
    parsedTotal := parsedData.ParsedTotal
    purchaseDate := parsedData.PurchaseDate
    purchaseTime := parsedData.PurchaseTime
    points := 0

    // Calculate points based on retailer characters
    for _, char := range receipt.Retailer {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            points++
        }
    }

    // Calculate points based on items
    for _, item := range receipt.Items {
        trimmedDescription := strings.TrimSpace(item.ShortDescription)
        if len(trimmedDescription)%3 == 0 {
            price, _ := strconv.ParseFloat(item.Price, 64) // Already validated, safe to ignore error
            pointsEarned := math.Ceil(price * 0.2)
            points += int(pointsEarned)
        }
    }

    // Additional point calculations
    if parsedTotal == math.Floor(parsedTotal) {
        points += 50 // Add 50 points if total is a round dollar amount
    }
    if math.Mod(parsedTotal, 0.25) == 0 {
        points += 25 // Add 25 points if total is a multiple of 0.25
    }
    points += (len(receipt.Items) / 2) * 5 // 5 points for every 2 items

    if purchaseDate.Day()%2 != 0 {
        points += 6 // 6 points if purchase date is an odd day
    }
    if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
        points += 10 // 10 points if purchase time is between 2:00 PM and 4:00 PM
    }

    // Return points in the response
    c.JSON(http.StatusOK, models.PointsResponse{Points: points})
}