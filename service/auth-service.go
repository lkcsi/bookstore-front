package service

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Auth(c *gin.Context)
}

type cookieAuthService struct {
}

type fakeAuthService struct {
}

func CookieAuthService() AuthService {
	return &cookieAuthService{}
}

func FakeAuthService() AuthService {
	return &fakeAuthService{}
}

func (a *fakeAuthService) Auth(c *gin.Context) {
	c.Next()
}

func (a *cookieAuthService) Auth(c *gin.Context) {
	if _, err := c.Cookie("username"); err == nil {
		c.Next()
		return
	}
	if h := c.GetHeader("Hx-Target"); h == "" {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		c.Writer.Header().Set("HX-Redirect", "/login")
		c.Abort()
	}
}
