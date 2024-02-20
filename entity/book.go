package entity

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title" form:"title"`
	Author   string `json:"author" form:"author"`
	Quantity int    `json:"quantity" form:"quantity"`
}
