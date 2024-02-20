package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/lkcsi/bookstore-front/custerror"
	"github.com/lkcsi/bookstore-front/entity"
)

type UserClient interface {
	Login(*entity.User) error
}

type userClient struct {
	apiKey string
}

func NewUserClient(apiKey string) UserClient {
	return &userClient{apiKey: apiKey}
}

func (u *userClient) Login(requestUser *entity.User) error {
	marshalled, _ := json.Marshal(requestUser)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8081/api/users/login", bytes.NewReader(marshalled))
	req.Header.Set("ApiKey", u.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return custerror.ApiError(resp.StatusCode, string(body))
	}

	return nil
}
