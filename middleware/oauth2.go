package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const UserKey = "user"

type Authenticator interface {
	Auth() []gin.HandlerFunc

	SetOAuth2(flow, authURL, tokenURL string, scopes []string)
	SetAPIKey(name, in string)
	Use(handlerFunc gin.HandlerFunc)
}

func UserInfoHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get(UserKey)
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
