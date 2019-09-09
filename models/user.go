package models

import (
	"blog.api/models/dto/user_dto"
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type User struct {
	Id       int
	Username string
	Password string
	Email   string
	CreateDate time.Time
	IsActive bool
	Photo string
	Remark string
	Ip string
}

//登录
func Login(username, password string)  User{
	o := orm.NewOrm()
	var user User
    o.Raw("SELECT id, username,photo FROM user WHERE username = ? and password =?", username,password).QueryRow(&user)

	return user
}

//获取用户基本信息
func GetUserInfo(id int) user_dto.UserInfoOutputDto {
	o := orm.NewOrm()
	var user user_dto.UserInfoOutputDto
	o.Raw("SELECT id, username,photo FROM user WHERE id = ?", id).QueryRow(&user)

	return user
}





