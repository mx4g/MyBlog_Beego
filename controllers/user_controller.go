package controllers

import (
	"blog.api/models"
	"blog.api/models/dto/user_dto"

	_ "blog.api/models/dto/user_dto"
	"blog.api/utils"
	"encoding/json"
	"strconv"
)

// 用户控制器
type UserController struct {
	BaseController
}


// @Title Get
// @Description 通过id用户基本信息
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} user_dto.UserInfoOutputDto
// @Failure 403 :id is empty
// @router /:id [get]
func (u *UserController) Get() {
	id := u.GetString(":id")
	if id != "" {
		//字符转整形
		num, _ := strconv.Atoi(id)
		user := models.GetUserInfo(num)
		u.Data["json"] = user

	}
	u.ServeJSON()
}


// @Title Login
// @Description 用户登录
// @Param	body	 body 	user_dto.UserLoginInputDto 	true		"用户登录内容"
// @Success 200 {bool} true
// @router /login [post]
func (u *UserController) Login() {
	var loginDto user_dto.UserLoginInputDto
	json.Unmarshal(u.Ctx.Input.RequestBody, &loginDto)
	username := loginDto.Username
	password := loginDto.Password

	user := models.Login(username,password)
	result := make(map[string]interface{})
    strId:= strconv.Itoa(user.Id)
	if user.Id ==0 {
		result["state"]=false
	}else{
		result["token"] = utils.GenerateToken(strId,user.Username)
		result["state"]=true
	}
	u.Data["json"] =result
	u.ServeJSON()
}

// @Title logout
// @Description 用户注销
// @Success 200 {bool} true
// @router /logout [get]
func (u *UserController) Logout() {
	u.DelSession("loginUser")
	u.Data["json"] = true
	u.ServeJSON()
}
