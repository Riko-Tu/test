package mysql

import (
	"encoding/hex"
	"golang.org/x/crypto/md4"
	"turan.com/web_demo/models"
)

//查询用户名是否存在,true:存在，false：不存在
func QueryUsserName(username string) bool {
	user := new(models.User)
	mysql.Raw("select * from user where username=?", username).Scan(user)
	if len(user.UserName) == 0 {
		return false
	}
	return true

}

//注册用户
func UserRegist(user models.User) bool {
	sql := "INSERT INTO user(user_id,username,password,email,gender,create_time,update_time) VALUE(?,?,?,?,?,?,?)"
	affected := mysql.Exec(sql, user.UserID, user.UserName, OPassword(user.PassWord), user.Email, user.Gender, user.CreateTime, user.UpdateTime).RowsAffected
	if affected == 0 {
		return false
	}
	return true
}

//查询用户
func QuryUser(username string) *models.User {
	user := new(models.User)
	sql := "select *  from user where username=?"
	mysql.Raw(sql, username).Scan(user)
	return user
}

//密码加密
func OPassword(password string) string {
	h := md4.New()
	h.Write([]byte("turan"))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
