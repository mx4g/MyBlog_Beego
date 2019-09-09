package response

//返回查询数据
type Result struct {
	Data interface{} `json:"data"`
	Total int64 `json:"total"`
	Code int64 `json:"code"`

}