package utils

import (
	"encoding/json"
	"fmt"
	"billing-microservice/models"
	"github.com/jung-kurt/gofpdf"
)

// Formatinvoice formats the invoice as json
func FormatInvoice(invoice models.Invoice) (string, error) {
	jsonData, err := json.MarshalIndent(invoice,""," " )
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
// Generate invoice pdf generates a pdf fot he invoice 
func GenerateInvoicePDF(invoice models.Invoice, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Electronic Invoice")

	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Order ID: %d", invoice.OrderID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Total: $%.2f", invoice.Amount))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Date: %s", invoice.Date))
    pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", invoice.Email))
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return err
	}
	return nil
}
