package logic

import (
	"errors"
	"time"
	"turan.com/web_demo/dao/mysql"
	"turan.com/web_demo/dao/redis"
	"turan.com/web_demo/models"
	"turan.com/web_demo/utils"
)

var (
	ErrorUserExist           = errors.New("用户已存在")
	ErrorUserNoExist         = errors.New("用户不存在")
	ErrorUserRegistFailed    = errors.New("注册失败")
	ErrorUserRegistSuccess   = errors.New("注册成功")
	ErrorUserLoginFailed     = errors.New("登陆失败")
	ErrorUserLoginSuccess    = errors.New("登录成功")
	ErrorUserTokenGetFailed  = errors.New("获取token失败")
	ErrorRedisSetTokenfailed = errors.New("保存token失败")
)

func UserRegist(user *models.User) (err error) {
	//1,查询用户名是否存在
	if mysql.QueryUsserName(user.UserName) {
		return ErrorUserExist
	}
	//2，生成uid
	id := utils.GetId()
	user.UserID = id
	//3.注册用户
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	if mysql.UserRegist(*user) {
		return ErrorUserRegistSuccess
	}
	return ErrorUserRegistFailed
}

func UserLogin(user *models.UserLogin) (token string, err error) {
	if !mysql.QueryUsserName(user.UserName) {
		return "", ErrorUserNoExist
	}
	quryUser := mysql.QuryUser(user.UserName)
	if mysql.OPassword(user.PassWord) != quryUser.PassWord {
		return "", ErrorUserLoginFailed
	}
	//生成token
	token, err = utils.GetToken(quryUser)
	if err != nil {
		return "", ErrorUserTokenGetFailed
	}
	err = redis.SetToken(quryUser, token)
	if err != nil {
		return token, ErrorRedisSetTokenfailed
	}
	return token, ErrorUserLoginSuccess
}
