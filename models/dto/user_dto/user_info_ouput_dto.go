package user_dto

import "time"

type UserInfoOutputDto struct {
	Username string `json:"username"`
	CreateDate time.Time `json:"createDate"`
	Photo string `json:"photo"`
	Remark string `json:"remark"`
}
