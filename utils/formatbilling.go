package utils

import (
	"encoding/json"
	"fmt"
	"billing-microservice/models"
	"github.com/jung-kurt/gofpdf"
)


func FormatFactura(factura models.Factura) (string, error) {
	jsonData, err := json.MarshalIndent(factura, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func GenerateFacturaPDF(factura models.Factura, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Factura Electrónica")

	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Orden ID: %d", factura.OrdenID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Total: $%.2f", factura.Total))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Fecha: %s", factura.Fecha))

	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return err
	}
	return nil
}
