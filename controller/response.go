package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"turan.com/web_demo/utils"
)

type ReCode int

var CtxAuthUserID string = "user_id"

const (
	CodeSuccess               ReCode = 10000 + iota //成功
	CodeUserExist                                   //用户名已存在
	CodeUserNoExist                                 //用户名不存在
	CodeUserPsdIncorrect                            //用户名或密码错误
	CodeLoginSuccess                                //登录成功
	CodeRegisFailed                                 //注册失败
	CodeRegisSuccess                                //注册成功
	CodeServicebusy                                 //服务繁忙
	CodeGetTokeErr                                  //获取token失败
	CodeResolvTokeErr                               //解析token错误
	CodeRedisUserNameNotExist                       //用户名不存在
	CodeTokeninvalid                                //token无效，请重新登录
	CodeLikeFailed                                  //点赞失败
	CodeLikeSuccess                                 //点赞成功
)

var CodeMap = map[ReCode]string{
	CodeSuccess:               "成功",
	CodeUserExist:             "用户名已存在",
	CodeUserNoExist:           "用户名不存在",
	CodeUserPsdIncorrect:      "用户名或密码错误",
	CodeRegisFailed:           "注册失败",
	CodeServicebusy:           "服务繁忙",
	CodeLoginSuccess:          "登录成功",
	CodeRegisSuccess:          "注册成功",
	CodeGetTokeErr:            "获取token失败",
	CodeResolvTokeErr:         "解析token错误",
	CodeRedisUserNameNotExist: "redis找不到该token",
	CodeTokeninvalid:          "token无效，请重新登录",

	CodeLikeFailed:  "点赞失败",
	CodeLikeSuccess: "点赞成功",
}

const (
	CommFriend = 1 + iota
	CommMachine
	CommERR
)

var CommunityID = map[int]string{
	CommFriend:  "朋友圈",
	CommMachine: "机器",
	CommERR:     "错误的社区",
}

func GetcommunityName(id int) string {
	value, ok := CommunityID[id]
	if !ok {
		return CommunityID[CommERR]
	}
	return value
}

func (r ReCode) GetMsg() string {
	value, ok := CodeMap[r]
	if !ok {
		return CodeMap[CodeServicebusy]
	}
	return value
}
func ResWithErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "err": err.Error()})
}
func ResWithData(c *gin.Context, code ReCode, data ...interface{}) {

	c.JSON(http.StatusOK, gin.H{"code": code, "msg": code.GetMsg(), "data": data})
}
func ResWithToken(c *gin.Context, code ReCode, token string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": code.GetMsg(), "token": token})
}
func ResWithMsg(c *gin.Context, code ReCode) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": code.GetMsg()})
}

func ParamsJson(c *gin.Context, obj interface{}) {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)), //返回时对去除结构体名称
			})
			return
		}
	}
}
