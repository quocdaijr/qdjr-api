package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"qdjr-api/helpers"
	"qdjr-api/services"
	"strconv"
	"strings"
)

type AuthMiddleware struct{}

func (_ AuthMiddleware) VerifiedToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		verifiedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(helpers.BaseHelper{}.GetEnv("SECRET_KEY", "")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthenticated"})
			return
		}

		sub, _ := verifiedToken.Claims.GetSubject()
		userId, _ := strconv.Atoi(sub)
		userService := services.UserService{}
		user := userService.Detail(userId)
		if user.Id == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User not found or not active"})
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
