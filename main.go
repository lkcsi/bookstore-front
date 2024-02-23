package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lkcsi/bookstore-front/client"
	"github.com/lkcsi/bookstore-front/controller"
	"github.com/lkcsi/bookstore-front/service"
)

func main() {
	godotenv.Load()
	server := gin.Default()

	apiKey := os.Getenv("API_KEY")

	bookClient := client.NewBookClient(apiKey)
	userClient := client.NewUserClient(apiKey)
	userBookClient := client.NewUserBookClient(apiKey)

	bookView := controller.NewBookView(&bookClient, &userBookClient)
	myBookView := controller.NewMyBookView(&bookClient, &userBookClient)
	userView := controller.NewLoginView(&userClient)

	authService := service.CookieAuthService()

	view := server.Group("")

	view.Use(authService.Auth)
	view.GET("/", bookView.Get)
	view.GET("/books", bookView.Get)
	view.GET("/my-books", myBookView.Get)
	view.POST("/checkout-book/:id", bookView.Checkout)
	view.POST("/return-book/:id", myBookView.Return)
	server.POST("/login", userView.Login)
	server.POST("/logout", userView.Logout)
	server.GET("/login", userView.Get)
	server.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("BOOKS_FRONT_PORT")))
}
