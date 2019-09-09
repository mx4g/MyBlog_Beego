package user_dto


type UserLoginInputDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RememberMe bool `json:"rememberMe"`
}