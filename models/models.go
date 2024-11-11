package models

import (
    "time"
)

type Receipt struct {
    Retailer      string `json:"retailer" binding:"required"`
    PurchaseDate  string `json:"purchaseDate" binding:"required,datetime=2006-01-02"` 
    PurchaseTime  string `json:"purchaseTime" binding:"required,datetime=15:04"`      
    Items         []Item `json:"items" binding:"required,dive"`                      
    Total         string `json:"total" binding:"required,numeric"`                   
}


type Item struct {
    ShortDescription string `json:"shortDescription" binding:"required"`
    Price            string `json:"price" binding:"required"`
}

type ReceiptResponse struct {
    ID string `json:"id"`
}

type PointsResponse struct {
    Points int `json:"points"`
}

type ParsedReceiptData struct {
    Receipt      Receipt
    ParsedTotal  float64
    PurchaseDate time.Time
    PurchaseTime time.Time
    Valid        bool
}