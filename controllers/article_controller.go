package controllers

import (
	"blog.api/models"
	_ "blog.api/models/dto/article_dto"
	"blog.api/models/dto/response"
	"encoding/json"
	"strconv"
)

// 文章控制器
type ArticleController struct {
	BaseController
}

// @Title Get
// @Description 通过id获取一篇文章
// @Param	id		path 	string	true		"文章id参数"
// @Success 200 {object} article_dto.ArticleOutputDto
// @Failure 403 :id is empty
// @router /:id [get]
func (u *ArticleController) Get() {
	idStr := u.GetString(":id")
	if idStr != "" {
		//字符转整形
		id, _ := strconv.Atoi(idStr)
		article := models.GetArticle(id)
		u.Data["json"] = article

	}
	u.ServeJSON()
}

// @Title GetAll
// @Description 获取所有的文章
// @Param	page		query  int	true		"页码"
// @Param	title		query  string	false		"标题"
// @Param	tag		query  string	false		"标签"
// @Success 200 {object} article_dto.ArticleListOutputDto
// @router /all [get]
func (u *ArticleController) GetAll() {

	page,err := u.GetInt("page")
	title  := u.GetString("title")
	tag := u.GetString("tag")

	var result response.Result

	if err != nil {
		result.Code = 500
	}else{
        if page == 0 { page =1 }

		result.Total = models.GetArticleCount(title,tag)
		if result.Total==0 {
			result.Data = []int{}
		}else {

			result.Data = models.GetAllArticle(page,title,tag)
		}

		result.Code = 200
	}



	u.Data["json"] = result

	u.ServeJSON()
}

// @Title AddArticle
// @Description 创建一篇文章
// @Param	body		body 	article_dto.ArticleInputDto 	true		"文章具体内容"
// @Success 200 {int} models.Article.Id
// @router / [post]
func(u *ArticleController) Post(){
	var article models.Article

	json.Unmarshal(u.Ctx.Input.RequestBody, &article)
	id := models.AddArticle(article)
	u.Data["json"] = id
	u.ServeJSON()
}

// @Title UpdateArticle
// @Description 更新一篇文章
// @Param	body		body 	article_dto.ArticleInputDto 	true		"文章具体内容"
// @Success 200 {int}
// @router / [put]
func(u *ArticleController) Put(){
	var article models.Article
	json.Unmarshal(u.Ctx.Input.RequestBody, &article)

	num := models.UpdateArticle(article)
	u.Data["json"] = num
	u.ServeJSON()
}


// @Title DeleteArticle
// @Description 删除一篇文章
// @Param	id		path 	string	true		"文章id参数"
// @Success 200 {int}
// @router /:id [delete]
func(u *ArticleController) Delete(){
	idStr := u.GetString(":id")
	if idStr != "" {
		//字符转整形
		id, _ := strconv.Atoi(idStr)
		num := models.DeleteArticle(id)
		u.Data["json"] = num
		u.ServeJSON()
		return
	}
	u.Data["json"] = 0
	u.ServeJSON()
}