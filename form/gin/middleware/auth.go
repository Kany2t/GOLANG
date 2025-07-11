package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 权限检查
func AuthCheck(c *gin.Context) {
	userId, _ := c.Get("user_Id")
	UserName, _ := c.Get("user_name")
	fmt.Printf("auth check userID :%s, UserName:%s\n", userId, UserName)
	c.Next()

}

// 登录校验
var token = "123456"

func TokenCheck(c *gin.Context) {
	accessToken := c.Request.Header.Get("access_token")
	if accessToken != token {
		c.JSON(401, gin.H{
			"message": "token error",
		})
		c.AbortWithError(401, errors.New("token error"))
	}
	c.Set("user_name", "nick")
	c.Set("user_id", 1)
	c.Next()
}
