// parser/parser.go

package parser

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "gin-docker-app/models"
    "strconv"
    "sync"
)

// ParseReceiptData standardizes the parsed receipt data and returns it as a struct
func ParseReceiptData(c *gin.Context, receiptID string, receiptDB *sync.Map) models.ParsedReceiptData {
    // Default response with Valid as false
    result := models.ParsedReceiptData{
        Receipt:      models.Receipt{},
        ParsedTotal:  0,
        PurchaseDate: time.Time{},
        PurchaseTime: time.Time{},
        Valid:        false,
    }

    // Load and validate the receipt
    receipt, exists := receiptDB.Load(receiptID)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
        return result
    }

    r, ok := receipt.(models.Receipt)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assert receipt type"})
        return result
    }

    // Parse purchase date
    purchaseDate, err := time.Parse("2006-01-02", r.PurchaseDate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid purchase date format"})
        return result
    }

    // Parse purchase time
    purchaseTime, err := time.Parse("15:04", r.PurchaseTime)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid purchase time format"})
        return result
    }

    parsedTotal, err := strconv.ParseFloat(r.Total, 64)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid total format"})
        return result
    }

    // Populate the fields in the struct and set Valid to true
    result = models.ParsedReceiptData{
        Receipt:      r,
        ParsedTotal:  parsedTotal,
        PurchaseDate: purchaseDate,
        PurchaseTime: purchaseTime,
        Valid:        true,
    }

    return result
}
