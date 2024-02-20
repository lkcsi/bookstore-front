package controller

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore-front/client"
	"github.com/lkcsi/bookstore-front/entity"
)

type MyBookView interface {
	Return(context *gin.Context)
	Get(context *gin.Context)
}

type myBookView struct {
	bookClient     client.BookClient
	userBookClient client.UserBookClient
}

func NewMyBookView(c *client.BookClient, uc *client.UserBookClient) MyBookView {
	return &myBookView{bookClient: *c, userBookClient: *uc}
}

func (b *myBookView) Get(context *gin.Context) {
	tmpl, _ := template.New("").ParseFiles("template/index.html", "template/my-books.html")
	username, _ := context.Cookie("username")
	books, err := b.userBookClient.FindAllByUsername(username)
	if err != nil {
		books = make([]entity.UserBook, 0)
	}

	tmpl.ExecuteTemplate(context.Writer, "index", gin.H{
		"Books": books,
	})
}

func (b *myBookView) Return(context *gin.Context) {
	id := context.Param("id")
	username, _ := context.Cookie("username")

	err := b.userBookClient.Return(username, id)
	if err != nil {
		setViewError(context, err)
	}
}
