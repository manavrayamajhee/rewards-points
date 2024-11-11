package main

import (
    "github.com/gin-gonic/gin"
    "gin-docker-app/handler"
)

func main() {
    router := gin.Default()
    router.POST("/receipts/process", handler.ProcessReceipt)
    router.GET("/receipts/:id/points", handler.GetReceiptPoints)
    router.Run(":8080") // Run on port 8080
}

