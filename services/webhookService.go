package services

import (
    "bytes"
    "encoding/json"
    "net/http"
    "billing-microservice/models"
)

// SendWebhookNotification sends a webhook notification with the billing data
func SendWebhookNotification(billing models.Invoice) error {
    webhookURL := "http://localhost:5001/webhook" // URL of the microservice
    jsonData, err := json.Marshal(billing)
    if err != nil {
        return err
    }

    // Create a new HTTP POST request with the JSON data
    req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }

    // Set the Content-Type header to application/json
    req.Header.Set("Content-Type", "application/json")

    // Send the request using an HTTP client
    client := &http.Client{}
    _, err = client.Do(req)
    return err
}