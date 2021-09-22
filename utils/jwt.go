package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"turan.com/web_demo/models"
)

const TokenExpirDuration = time.Hour * 2 //定义过期时间

var MySercet = []byte("turan") //定义加密密钥

type Myclaims struct {
	UserID             int64  `json:"user_id"`
	UserName           string `json:"user_name"`
	jwt.StandardClaims        //包含官方字段
}

//生成token
func GetToken(user *models.User) (string, error) {
	c := Myclaims{
		UserID:   user.UserID,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpirDuration).Unix(), //指定token过期时间
			Issuer:    "turan",
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return Token.SignedString(MySercet)
}

//解析token
func ParamsToken(tokenString string) (*Myclaims, error) {
	var mc = new(Myclaims) //存放token解析后的数据
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return MySercet, nil
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if token.Valid { //校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
