package errs

import (
	"fmt"
)

type TgError struct {
	Err    error
	ChatID int64
}

type InternalError struct {
	Cause string
}

func (e InternalError) Error() string {
	return fmt.Sprintf("internal server error: %s", e.Cause)
}

type NotFoundError struct {
	What string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.What)
}

type BadRequestError struct {
	Cause string
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("bad request: %s", e.Cause)
}

type ForbiddenError struct {
	Cause string
}

func (e ForbiddenError) Error() string {
	return fmt.Sprintf("forbidden: %s", e.Cause)
}

type Unauthorized struct {
	Cause string
}

func (e Unauthorized) Error() string {
	return fmt.Sprintf("forbidden: %s", e.Cause)
}
