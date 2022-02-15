package user

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	NickName string `json:"nickName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
