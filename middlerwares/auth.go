package middlerwares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justmamadou/rest-api-golang/utils"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized !"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized !"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
