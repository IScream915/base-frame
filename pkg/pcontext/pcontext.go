package pcontext

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	CtxUserKey = "user" // 用户信息
)

// GetRequestToken 从请求头中获取token
func GetRequestToken(c *gin.Context) string {
	// 从请求头中获取Authorization参数
	authorization := c.Request.Header.Get("Authorization")
	// 如果Auth为空直接返回
	if authorization == "" {
		return ""
	}

	tokens := strings.SplitN(authorization, " ", 2)
	if len(tokens) == 2 || tokens[0] == "Bearer" {
		return tokens[1]
	}
	return ""
}
