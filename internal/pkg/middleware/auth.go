package middleware

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"strings"
)
import "github.com/golang-jwt/jwt/v5"

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.Next()
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims := &api.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return base64.URLEncoding.DecodeString(config.AppConfig.System.SecretKey)
		})
		if err != nil || !token.Valid {
			c.Next()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func GetClaims(c *gin.Context) *api.UserClaims {
	claims, exist := c.Get("claims")
	if !exist {
		return nil
	}
	return claims.(*api.UserClaims)
}
