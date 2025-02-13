package controllers

import (
    "billing-microservice/config"
    "billing-microservice/models"
    "billing-microservice/services"
    "billing-microservice/utils"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ReceivePayment handles incoming payment webhook notifications
func ReceivePayment(c *gin.Context) {
    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Aquí puedes procesar el pago (por ejemplo, guardar en la base de datos, enviar notificaciones, etc.)
    // Por ahora, simplemente devolvemos una respuesta exitosa
	// Guardar el pago en la base de datos
    result := config.DB.Create(&payment)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Payment received successfully", "payment": payment})
}
// Create a new billing
func CreateBilling(c *gin.Context) {
    var billing models.Invoice
    if err := c.ShouldBindJSON(&billing); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save billing into the database
    result := config.DB.Create(&billing)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    // Send webhook notification
    err := services.SendWebhookNotification(billing)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending notification"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Billing created successfully", "billing": billing})
}

// Download a billing PDF
func DownloadBilling(c *gin.Context) {
    billingID := c.Param("id")
    var billing models.Invoice

    // Find the billing by ID
    result := config.DB.First(&billing, billingID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Billing not found"})
        return
    }

    // Generate PDF file
    filename := fmt.Sprintf("billing_%d.pdf", billing.OrderID)
    err := utils.GenerateInvoicePDF(billing, filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating PDF"})
        return
    }

    // Return the PDF file
    c.File(filename)
}