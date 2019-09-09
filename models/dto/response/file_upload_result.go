package response

//返回查询数据
type FileUploadResult struct {
	Url string `json:"url"`
	Message string `json:"message"`
	Code int64 `json:"code"`

}