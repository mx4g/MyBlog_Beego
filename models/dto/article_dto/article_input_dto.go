package article_dto

type ArticleInputDto struct {

	Title string `json:"title"`
	Content   string `json:"content"`
	HtmlContent string `json:"htmlContent"`
	Tag string `json:"tag"`
	FirstImage string `json:"firstImg"`

	Remark string `json:"desc"`
}
