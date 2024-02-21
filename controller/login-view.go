package controller

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/lkcsi/bookstore-front/client"
	"github.com/lkcsi/bookstore-front/entity"
)

type LoginView interface {
	Get(context *gin.Context)
	Login(context *gin.Context)
	Logout(context *gin.Context)
}

type loginView struct {
	userClient client.UserClient
}

func NewLoginView(c *client.UserClient) LoginView {
	return &loginView{userClient: *c}
}

func (l *loginView) Get(context *gin.Context) {
	tmpl, _ := template.New("").ParseFiles("template/login.html")
	tmpl.ExecuteTemplate(context.Writer, "login", nil)
}

func (l *loginView) Login(context *gin.Context) {
	var requestUser entity.User
	if err := context.ShouldBind(&requestUser); err != nil {
		setViewError(context, err)
		return
	}
	err := l.userClient.Login(&requestUser)
	if err != nil {
		setViewError(context, err)
		return
	}
	context.SetCookie("username", requestUser.Username, (10 * 60), "/", "localhost", false, true)
	context.Writer.Header().Add("HX-Redirect", "/books")
}

func (l *loginView) Logout(context *gin.Context) {
	context.SetCookie("username", "deleted", -1, "/", "localhost", false, true)
	context.Writer.Header().Add("HX-Redirect", "/")
}
