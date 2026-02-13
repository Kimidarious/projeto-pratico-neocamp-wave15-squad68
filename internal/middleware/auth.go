package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/service"
)

// AuthMiddleware verifica se o token JWT é válido
func AuthMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log início da request
		startTime := time.Now()
		log.Printf("🔐 [AUTH START] %s %s", c.Request.Method, c.Request.URL.Path)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format. Expected: Bearer <token>"})
			c.Abort()
			return
		}

		token := parts[1]

		// 3. Validar o token
		claims, err := authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// 4. Adicionar informações do usuário no contexto para usar nos handlers
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.UserName)
		c.Set("user_type", claims.UserType)

		// 5. Continuar para o próximo handler
		c.Next()

		// Log fim da request com duração
		duration := time.Since(startTime)
		log.Printf("✅ [AUTH END] %s %s | Status: %d | Duration: %v | User: %s", 
			c.Request.Method, 
			c.Request.URL.Path, 
			c.Writer.Status(),
			duration,
			claims.UserName,
		)
	}
}
