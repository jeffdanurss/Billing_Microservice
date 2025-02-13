package middlewares

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

const authServiceURL = "http://auth-service:8080/validate-token"

// AuthMiddleware is a middleware to validate the authorization token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
            c.Abort()
            return
        }

        // Create a new request to validate the token
        req, err := http.NewRequest("GET", authServiceURL, nil)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
            c.Abort()
            return
        }

        // Set the Authorization header for the request
        req.Header.Set("Authorization", authHeader)

        // Send the request to the authentication service
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Check if the response status code is OK (200)
        if resp.StatusCode != http.StatusOK {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Proceed to the next middleware or handler
        c.Next()
    }
}