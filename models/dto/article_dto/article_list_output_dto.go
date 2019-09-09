package article_dto

import "time"

type ArticleListOutputDto struct {
	Id       int  `json:"id"`
	Avatar string  `json:"avatar"`
	Title string `json:"title"`
	Remark   string `json:"remark"`
	FirstImage string `json:"firstImage"`

	CreateDate time.Time `json:"createDate"`
}
