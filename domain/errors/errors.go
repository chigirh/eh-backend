package errors

import (
	"eh-backend-api/domain/models"
	"fmt"
)

type NotFoundError struct {
	Sources string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found:%v", e.Sources)
}

type AuthenticationError struct {
	UserName models.UserName
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("Authentication error:%v", e.UserName)
}

type SystemError struct {
	Message string
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("Internal server error. message:%v", e.Message)
}
