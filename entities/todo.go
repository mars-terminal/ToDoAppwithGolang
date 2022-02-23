package entities

type TodoList struct {
	Id          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId string
	ListId string
}

type TodoItem struct {
	Id          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
