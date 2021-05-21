package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const UserKey = "user"

type Authenticator struct {
	Scopes        []string
	OAuth2Handler func(ctx *gin.Context)
	Key           string
	ApiKeyValid   func(string) bool
}

func (a *Authenticator) Auth(ctx *gin.Context) {
	if a.ApiKeyValid != nil {
		keyHeader := ctx.GetHeader(a.Key)

		if a.ApiKeyValid(keyHeader) {
			ctx.Next()
			return
		}

		if keyHeader != "" && !a.ApiKeyValid(keyHeader) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong api key"})
			return
		}
	}

	if a.OAuth2Handler != nil {
		a.OAuth2Handler(ctx)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
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
