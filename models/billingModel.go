package models

import "gorm.io/gorm"

//Payment model
type Payment struct {
    OrderID string   `json:"orderId"`
    Amount float64 `json:"amount"`
    Date   string `json:"date,omitempty`
    Email  string  `json:"email"`
    Currency string `json:"currency"`
}

// Invoice model
type Invoice struct {
    gorm.Model
    OrderID string    `json:"order_id"`
    Amount float64 `json:"amount"`
    Total float64 `json:"amount"`
    Date    string  `json:"date"`
    Email   string  `json:"email"` // Added email field
    Currency   string `json:"currency"`
    
    
}