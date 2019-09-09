package models

import "time"

type Comment struct {
	Id       int
	UserId string
	Content   string
	ParentCommentId string
	CreateDate time.Time
	Article    *Article   `orm:"rel(fk)"`
	IsActive bool
}