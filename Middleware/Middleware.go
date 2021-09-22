package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"turan.com/web_demo/controller"
	Re "turan.com/web_demo/dao/redis"
	"turan.com/web_demo/utils"
)

func IsLikeStutas() func(c *gin.Context) {
	return func(c *gin.Context) {
		getValue := Re.Client.Get(controller.CtxAuthUserID + "Like")
		i, _ := getValue.Int()
		if getValue.String() == "" || i == 0 {
			c.Next()
			return
		}
		c.Abort()
		c.JSON(http.StatusOK, gin.H{"mseg": "你已点赞"})
		return
	}
}

func DisLikeStutas() func(c *gin.Context) {
	return func(c *gin.Context) {
		getValue := Re.Client.Get(controller.CtxAuthUserID + "Like")
		i, _ := getValue.Int()
		if getValue.String() == "" || i == 0 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{"msg": "请先点赞"})
			return
		}
		c.Next()
		return
	}
}

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		Querytoken := c.Request.Header.Get("token")
		if Querytoken == "" { //token为空
			controller.ResWithMsg(c, controller.CodeGetTokeErr) //直接返回
			c.Abort()                                           //终止后续函数
			return
		}

		mc, err := utils.ParamsToken(Querytoken) //将token信息映射到mc中
		if err != nil {
			controller.ResWithMsg(c, controller.CodeResolvTokeErr)
			c.Abort()
			return
		}
		Redistoken, err := Re.Client.Get(mc.UserName).Result()
		if err == redis.Nil { //redis.nil:为不存在
			controller.ResWithMsg(c, controller.CodeRedisUserNameNotExist)
			c.Abort()
			return
		} else if err != nil { //再判断其他类型错误
			controller.ResWithErr(c, err)
			c.Abort()
			return
		} else if Querytoken != Redistoken {
			controller.ResWithMsg(c, controller.CodeTokeninvalid)
			c.Abort()
			return
		}
		c.Set(controller.CtxAuthUserID, mc.UserID) //将token解析出来的数据存在下上文中
		c.Next()
	}
}
