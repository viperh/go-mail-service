package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-mail-service/internal/pkg/config"
	"net/http"
	"time"
)

type JwtMiddleware struct {
	Cfg *config.Config
}

func NewMiddleware(cfg *config.Config) *JwtMiddleware {
	return &JwtMiddleware{Cfg: cfg}
}

func (f *JwtMiddleware) JwtMiddlewareFunc(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
	}

	token = token[7:]

	if !validateToken(token, f.Cfg) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
		c.Abort()
		return
	}

	c.Next()
}

func validateToken(token string, cfg *config.Config) bool {

	jwtSecret := cfg.JwtSecret

	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !tokenObj.Valid {
		return false
	}

	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	expiration, ok := claims["exp"].(float64)
	if !ok {
		return false
	}

	expirationTime := time.Unix(int64(expiration), 0)
	if expirationTime.Before(time.Now()) {
		return false
	}

	return true

}
