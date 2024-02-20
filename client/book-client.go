package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/lkcsi/bookstore-front/custerror"
	"github.com/lkcsi/bookstore-front/entity"
)

type BookClient interface {
	FindAll() ([]entity.Book, error)
	Save(*entity.Book) (*entity.Book, error)
}

type bookClient struct {
	apiKey string
}

func NewBookClient(apiKey string) BookClient {
	return &bookClient{apiKey: apiKey}
}

func (b *bookClient) FindAll() ([]entity.Book, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8081/api/books", nil)
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

	var books []entity.Book
	err = json.Unmarshal(body, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *bookClient) Save(requestBook *entity.Book) (*entity.Book, error) {
	marshalled, _ := json.Marshal(requestBook)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8081/api/books", bytes.NewReader(marshalled))
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

	var book entity.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
