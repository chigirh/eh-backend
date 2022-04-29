package errors

import (
	"fmt"
)

type NotFoundError struct {
	Sources string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found:%s", e.Sources)
}

type AlreadyExistsError struct {
	Sources string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("Already exists:%s", e.Sources)
}

type AuthenticationError struct {
	Sources string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("Authentication error:%s", e.Sources)
}

type SystemError struct {
	Message string
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("Internal server error. message:%s", e.Message)
}
