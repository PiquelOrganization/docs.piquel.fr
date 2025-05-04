package utils

import (
	"fmt"
	"net/http"
)

type Error struct {
	status int
	error
}

func NewError(message string, status int) Error {
	err := fmt.Errorf(message).(Error)
	err.status = status

	return err
}

func (e *Error) Handle(w http.ResponseWriter) {
	http.Error(w, e.Error(), e.status)
}
