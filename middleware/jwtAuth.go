package middleware

import (
	"bioskop/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := utils.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			fmt.Println(err)
			context.Abort()
			return
		}
		context.Next()
	}
}
