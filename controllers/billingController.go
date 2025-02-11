package controllers

import (
	"billing-microservice/database"
	"billing-microservice/models"
    "billing-microservice/utils"
    "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
    "time"
)

// webhoook
func ReceivePayment(c *gin.Context) {
    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    // Validate oayment data
    if err := validatePayment(payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // today date
    if payment.Date == "" {
        payment.Date = time.Now().Format("2006-01-02") // Format YYYY-MM-DD
    }

    // Create an invoice associated with the payment
    invoice := models.Invoice{
        OrderID: payment.OrderID,
        Amount:   payment.Amount,
        Total: payment.Amount,
        Date:   payment.Date,
		Email: payment.Email,
        Currency: payment.Currency,
    }

    // Save the invoice to the database
    result := database.DB.Create(&invoice)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
        return
    }

    // Generate a PDF for the voice
    filename := fmt.Sprintf("invoice_%d.pdf", invoice.OrderID)
    if err := utils.GenerateInvoicePDF(invoice, filename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
        return
    }

    // Respuesta exitosa
    c.JSON(http.StatusOK, gin.H{"message": "Payment received and invocie created", "invoice": invoice})
}

// Function to validate payment data
func validatePayment(payment models.Payment) error {
    if payment.OrderID == "" || payment.Amount <= 0 || payment.Email == "" {
        return fmt.Errorf("Payment data is incomplete or invalid")
    }
    return nil
}
	