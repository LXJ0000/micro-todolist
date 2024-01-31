package router

import (
	"github.com/gin-gonic/gin"
	http_ "micro-todolist/app/gateway/http"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "success")
		})
		v1.POST("/user/register", http_.UserRegisterHandler)
		v1.POST("/user/login", http_.UserLoginHandler)
	}
	return r
}
