package entity

type UserBook struct {
	Id       string `json:"book-id"`
	Title    string `json:"title" form:"title"`
	Author   string `json:"author" form:"author"`
	Username string `json:"username"`
}
