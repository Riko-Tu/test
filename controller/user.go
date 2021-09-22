package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"turan.com/web_demo/logic"
	"turan.com/web_demo/models"
)

//UserRegist godoc
//@Summary 注册一个用户
//@Tags 用户模块
// @Description get object by models.User
//@Accept json
//@Produce  json
// @Router /front/login [post]
func UserRegist(c *gin.Context) {
	//1,获取参数
	user := new(models.User)
	ParamsJson(c, user)
	//3.业务处理
	err := logic.UserRegist(user)
	if errors.Is(err, logic.ErrorUserExist) {
		ResWithMsg(c, CodeUserExist)
		return
	} else if errors.Is(err, logic.ErrorUserRegistFailed) {
		ResWithMsg(c, CodeRegisFailed)
		return
	} else if errors.Is(err, logic.ErrorUserRegistSuccess) {
		ResWithMsg(c, CodeRegisSuccess)
		return
	}

}

func UserLogin(c *gin.Context) {
	user := new(models.UserLogin)
	ParamsJson(c, user)
	token, err := logic.UserLogin(user)
	if errors.Is(err, logic.ErrorUserNoExist) {
		ResWithMsg(c, CodeUserNoExist)
	} else if errors.Is(err, logic.ErrorUserLoginFailed) {
		ResWithMsg(c, CodeUserPsdIncorrect)
	} else if errors.Is(err, logic.ErrorUserLoginSuccess) {
		ResWithToken(c, CodeLoginSuccess, token)
	} else if errors.Is(err, logic.ErrorUserTokenGetFailed) {
		ResWithMsg(c, CodeGetTokeErr)
	}
}
