package middleware

import (
	"base_frame/internal/repo"
	"github.com/gin-gonic/gin"
)

func Auth(tokenRepo repo.UserToken) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
