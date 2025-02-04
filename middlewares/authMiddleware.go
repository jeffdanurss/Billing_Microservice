package middlewares

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

const authServiceURL = "http://auth-service:8080/validate-token"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		req, err := http.NewRequest("GET", authServiceURL, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear solicitud"})
			c.Abort()
			return
		}
		req.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}
		if resp.StatusCode != http.StatusOK {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
            c.Abort()
            return
        }

		c.Next()
	}
}
