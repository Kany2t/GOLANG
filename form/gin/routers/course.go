package routers

import (
	"form/gin/middleware"
	"form/gin/web"
	"github.com/gin-gonic/gin"
)

func initCourse(r *gin.Engine) {
	v1 := r.Group("/v1", middleware.TokenCheck, middleware.AuthCheck)
	v1.POST("/course", web.CreateCourse)
	v1.GET("/course", web.GetCourse)
	v1.PUT("/course", web.EditCourse)
	v1.DELETE("/course", web.DeleteCourse)

	v2 := r.Group("/v2")
	v2.POST("/course", web.CreateCourseV2)
	v2.PUT("/course", web.EditCourseV2)
}
