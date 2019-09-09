package models

import (
	"blog.api/models/dto/article_dto"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)


//文章表
type Article struct {
	Id       int `json:"id"`
	UserId string `json:"userId"`
	Title string  `json:"title"`
	Content  string `orm:"type(text)" json:"content"`
	HtmlContent string `orm:"type(text)" json:"htmlContent"`
	FirstImage string `json:"firstImage"`

	Tag string `json:"tag"`
	Remark string `json:"remark"`
	LikeCount int `json:"likeCount"`
	ViewCount int `json:"viewCount"`
	Comments    	[]*Comment `orm:"reverse(many)" json:"comments"` // 设置一对多的反向关系
	CreateDate time.Time `json:"createDate"`
	IsActive bool `json:"isActive"`

}

//获取所有文章
func GetAllArticle(pageNum int,title, tag string) []Article{

	var articles []Article

	//条件
	cond := orm.NewCondition()

	cond1 := cond

	if title !="" {
		cond1 = cond.Or("title__icontains", title)
	}
	if tag !="" {
		cond1 = cond1.Or("tag__icontains", tag)
	}



	qs := orm.NewOrm().QueryTable("article")
	qs = qs.SetCond(cond1)

	//每页记录条数
	limit, _ := beego.AppConfig.Int("PageSize")

	//上一页的条数
	 var offset = (pageNum - 1) * limit

	//查询
	_, err := qs.Limit(limit, offset).OrderBy("-id").All(&articles, "Id","Title","Remark", "FirstImage")

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	return articles

	//o := orm.NewOrm()
	//var articles []article_dto.ArticleListOutputDto
	//
	////每页记录条数
	//var pageSize = beego.AppConfig.String("PageSize")
	////上一页的条数
	//var previousPageSize = (pageNum - 1) * 10
	//
	//
	//_, err := o.Raw("SELECT * FROM article Where title like '%?%' or tag like '%?%' ORDER BY Id DESC  limit ?,?",title,tag,previousPageSize,pageSize).QueryRows(&articles)
	//
	//if err != nil {
	//	fmt.Println("Error: ", err.Error())
	//
	//}
	//
	//return articles
}

//获取所有统计数
func GetArticleCount(title, tag string) int64 {
	o := orm.NewOrm()

	//条件
	cond := orm.NewCondition()

	cond1 := cond

	if title !="" {
		cond1 = cond.Or("title__icontains", title)
	}
	if tag !="" {
		cond1 = cond1.Or("tag__icontains", tag)
	}
	qs := o.QueryTable("article")
	qs = qs.SetCond(cond1)
	num, err := qs.Count()

	if err == nil {
		fmt.Println("GetArticleCount : ", num)

	}

	return num
}

func GetArticle(id int) article_dto.ArticleOutputDto {
	o := orm.NewOrm()
	var article article_dto.ArticleOutputDto
	err := o.Raw("SELECT id, title,remark,first_image,content,html_content,tag,create_date FROM Article WHERE id = ?", id).QueryRow(&article)
	if err != nil {
		fmt.Println("Error: ", err.Error())

	}

	return article
}

func AddArticle(article Article) int64 {
	o := orm.NewOrm()

	article.CreateDate = time.Now()
    article.IsActive = true

	id, err := o.Insert(&article)
	if err != nil {
		return 0
	}
	return  id
}

func UpdateArticle(updateArticle Article) int64 {

	o := orm.NewOrm()
	article := Article{Id: updateArticle.Id}

	if o.Read(&article) == nil {

		if num, err := o.Update(&updateArticle, "Title", "Tag","FirstImage","Remark","HtmlContent","Content"); err == nil {

			fmt.Println("UpdateArticle : ", num)
			return  num
		}
	}
	return  0
}

func DeleteArticle(id int) int64{
	o := orm.NewOrm()
	num, err := o.Delete(&Article{Id: id});

	if  err != nil {
		fmt.Println("Error: ", err.Error())
	}

	return num

}
