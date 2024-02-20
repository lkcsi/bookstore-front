package custerror

import (
	"fmt"
)

type CustError interface {
	Code() int
	Error() string
}

type custError struct {
	code    int
	message string
}

func (c *custError) Code() int {
	return c.code
}

func (err *custError) Error() string {
	return err.message
}

func BookOutOfStockError(id string) CustError {
	return &custError{400, fmt.Sprintf("book with id: %s is out of stock", id)}
}

func BookNotFoundError(id string) CustError {
	return &custError{404, fmt.Sprintf("book with id: %s is not found", id)}
}

func UserNotFoundError(username string) CustError {
	return &custError{404, fmt.Sprintf("user with username: %s is not found", username)}
}

func OccupiedUsernameError(username string) CustError {
	return &custError{400, fmt.Sprintf("username: %s has already been in use", username)}
}

func InvalidPasswordError() CustError {
	return &custError{401, "invalid password"}
}

func ApiError(statusCode int, message string) CustError {
	return &custError{statusCode, message}
}
