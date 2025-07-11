package web

import "github.com/gin-gonic/gin"

type loginReq struct {
	Username string `form:"user_name"`
	Password string `form:"pwd"`
}

func Login(c *gin.Context) {
	req := loginReq{}
	c.Bind(&req)
	//c.BindQuery()
	//c.BindJSON()
	//err :=c.ShouldBind()
	c.JSON(200, req)
}
