package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"turan.com/web_demo/Middleware"
	"turan.com/web_demo/controller"

	"turan.com/web_demo/log"
	"turan.com/web_demo/utils"
)

var R *gin.Engine

func Init() {
	SetUp()
	User()
	JWT()
	Community()
	R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetUp() {
	R = gin.New()
	R.Use(log.GinLogger(), log.GinRecovery(true))
	utils.InitTrans("zh")
}

func User() {
	R.POST("/regist", controller.UserRegist)
	R.POST("login", controller.UserLogin)

}

func JWT() {
	R.GET("/ping", Middleware.JwtAuthMiddleware(), func(c *gin.Context) {
		value, _ := c.Get(controller.CtxAuthUserID)
		c.JSON(http.StatusOK, gin.H{"userid": value})
	})
}
func Community() {
	R.GET("/community", Middleware.JwtAuthMiddleware(), controller.GetCommunity)
	R.PUT("/community", Middleware.JwtAuthMiddleware(), controller.PutCommunity)
	R.POST("/community/like", Middleware.JwtAuthMiddleware(), controller.CommunityLike)
	R.POST("/community/dislike", Middleware.JwtAuthMiddleware(), controller.CommunityDisLike)
}
