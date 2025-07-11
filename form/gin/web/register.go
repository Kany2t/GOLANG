package web

import "github.com/gin-gonic/gin"

// 参数校验
type registerReq struct {
	UserName string `form:"user_name" binding:"required"`    //必填
	Password string `form:"pwd" binding:"required"`          //必填
	Phone    string `form:"phone" binding:"required e164"`   //手机号格式
	Email    string `form:"email" binding:"omitempty,email"` //不必填，如果为空忽略
}

func Register(c *gin.Context) {
	req := &registerReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, req)
}
