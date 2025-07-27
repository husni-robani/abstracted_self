package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := utils.GetToken(c.GetHeader("Authorization"))

		if err != nil || tokenString != os.Getenv("AUTH_TOKEN") {
			if err != nil {
				logger.Error.Println("GetToken error: ", err)
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Access denied",
			})
			return
		}
		
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method: %v", t.Header["alg"])
			}

			return []byte(os.Getenv("AUTH_KEY")), nil
		})

		if err != nil {
			logger.Info.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Access denied",
			})
			return
		}

		if _, ok := token.Claims.(*jwt.RegisteredClaims); !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Access denied",
			})
			return
		}


		c.Next()
	}
}