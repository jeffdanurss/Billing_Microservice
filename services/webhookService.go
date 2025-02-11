package services

import (
    "encoding/json"
    "bytes"
    "net/http"
    "billing-microservice/models"
)

func SendWebhookNotification(factura models.Factura) error {
    webhookURL := "http://localhost:5001/webhook" // URL microservices
    jsonData, err := json.Marshal(factura)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    if err != nil {
        return err
    }

    client := &http.Client{}
    _, err = client.Do(req)
    return err
}
