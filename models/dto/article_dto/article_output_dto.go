package article_dto

import "time"

type ArticleOutputDto struct {

	Id       int  `json:"id"`
	Title string `json:"title"`
	Remark   string `json:"remark"`
	FirstImage string `json:"firstImage"`

	Content   string `json:"content"`
	HtmlContent string `json:"htmlContent"`
	Tag string `json:"tag"`
	CreateDate time.Time `json:"createDate"`

}
