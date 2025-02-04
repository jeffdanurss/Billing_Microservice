package models

import (
    "gorm.io/gorm"
)

type Factura struct {
    gorm.Model
    OrdenID uint   `json:"orden_id"`
    Total   float64 `json:"total"`
    Fecha   string  `json:"fecha"`
}
