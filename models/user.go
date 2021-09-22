package models

import "time"

//binding参数校验器
type User struct {
	//Id         int       `json:"id" gorm:"column:id"`
	UserID     int64     `json:"user_id" gorm:"column:user_id"`
	UserName   string    `json:"username" binding:"required" gorm:"column:username"`
	PassWord   string    `json:"password" binding:"required" gorm:"column:password"`
	Email      string    `json:"email" binding:"required" gorm:"column:email"`
	Gender     bool      `json:"gender" binding:"required" gorm:"column:gender"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}
type UserLogin struct {
	UserName string `json:"username" binding:"required" gorm:"column:username"`
	PassWord string `json:"password" binding:"required" gorm:"column:password"`
	Token    string `json:"token"`
}
