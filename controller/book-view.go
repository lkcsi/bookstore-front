package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore-front/client"
	"github.com/lkcsi/bookstore-front/entity"
)

type BookView interface {
	Save(context *gin.Context)
	Checkout(context *gin.Context)
	Get(context *gin.Context)
}

type bookView struct {
	bookClient     client.BookClient
	userBookClient client.UserBookClient
}

func NewBookView(c *client.BookClient, uc *client.UserBookClient) BookView {
	return &bookView{bookClient: *c, userBookClient: *uc}
}

func checkoutButtonHtml(id string) string {
	return fmt.Sprintf("<button class='btn btn-success' hx-target='#book-%s' hx-post='/checkout-book/%s' hx-swap='outerHTML'> Checkout </button>", id, id)
}

func itemHtml(book *entity.Book) string {
	return fmt.Sprintf("<tr id='book-%s'><td>%s</td><td>%s</td><td>%d</td><td>%s</td></tr>",
		book.Id, book.Title, book.Author, book.Quantity, checkoutButtonHtml(book.Id))
}

func (b *bookView) Save(context *gin.Context) {
	var requestedBook entity.Book
	if err := context.ShouldBind(&requestedBook); err != nil {
		setViewError(context, err)
		return
	}

	newBook, err := b.bookClient.Save(&requestedBook)
	if err != nil {
		setViewError(context, err)
		return
	}

	tmpl, _ := template.New("t").Parse(itemHtml(newBook))
	tmpl.Execute(context.Writer, nil)
}

func (b *bookView) Get(context *gin.Context) {
	fmt.Println("BookView Called")
	tmpl, _ := template.New("").ParseFiles("template/index.html", "template/books.html")
	books, err := b.bookClient.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		books = make([]entity.Book, 0)
	}

	tmpl.ExecuteTemplate(context.Writer, "index", gin.H{
		"Books": books,
	})
}

func (b *bookView) Checkout(context *gin.Context) {
	id := context.Param("id")
	username, _ := context.Cookie("username")

	book, err := b.userBookClient.Checkout(username, id)
	if err != nil {
		fmt.Println(err.Error())
		setViewError(context, err)
		return
	}

	tmpl, _ := template.New("t").Parse(itemHtml(book))
	tmpl.Execute(context.Writer, nil)
}
