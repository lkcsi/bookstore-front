package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lkcsi/bookstore-front/custerror"
	"github.com/lkcsi/bookstore-front/entity"
)

type UserBookClient interface {
	FindAllByUsername(string) ([]entity.UserBook, error)
	Checkout(string, string) (*entity.Book, error)
	Return(string, string) error
}

type userBookClient struct {
	apiKey string
}

func NewUserBookClient(apiKey string) UserBookClient {
	return &userBookClient{apiKey: apiKey}
}

func (b *userBookClient) FindAllByUsername(username string) ([]entity.UserBook, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/api/user-books/%s", username), nil)
	req.Header.Set("ApiKey", b.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, custerror.ApiError(resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var books []entity.UserBook
	err = json.Unmarshal(body, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *userBookClient) Checkout(username, id string) (*entity.Book, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("http://localhost:8081/api/user-books/%s/checkout/%s", username, id), nil)
	req.Header.Set("ApiKey", b.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if resp.StatusCode != 202 {
		body, _ := io.ReadAll(resp.Body)
		var apiError entity.ApiError
		json.Unmarshal(body, &apiError)
		return nil, custerror.ApiError(resp.StatusCode, apiError.Error)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var book entity.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *userBookClient) Return(username, id string) error {
	client := &http.Client{}
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("http://localhost:8081/api/user-books/%s/return/%s", username, id), nil)
	req.Header.Set("ApiKey", b.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if resp.StatusCode != 202 {
		body, _ := io.ReadAll(resp.Body)
		var apiError entity.ApiError
		json.Unmarshal(body, &apiError)
		return custerror.ApiError(resp.StatusCode, apiError.Error)
	}

	return nil
}
