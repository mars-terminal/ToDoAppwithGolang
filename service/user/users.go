package user

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
