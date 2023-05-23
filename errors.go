package utils

import (
	"fmt"
	"github.com/pkg/errors"
)

// WrapErr Wraps an error with formatted string. If the error is nil,
// then nil will be returned.
func WrapErr(err error, fmtString string, items ...interface{}) error {
	if err == nil {
		return nil
	}

	if len(items) == 0 {
		return errors.Wrap(err, fmtString)
	}

	return errors.Wrap(err, fmt.Sprintf(fmtString, items...))
}

// HttpError Simply makes it cleaner to send an error response back
// in a http handler after returning from another function.
type HttpError struct {
	UserFacingMessage string
	StatusCode        int
	Err               error
}

func (e *HttpError) Error() string {
	var msg = e.UserFacingMessage
	if msg == "" && e.Err != nil {
		msg = e.Err.Error()
	}
	return msg
}

func (e *HttpError) WithError(err error) *HttpError {
	e.Err = err
	return e
}

func (e *HttpError) WithMessage(msg string) *HttpError {
	e.UserFacingMessage = msg
	return e
}

func (e *HttpError) WithStatus(statusCode int) *HttpError {
	e.StatusCode = statusCode
	return e
}

func (e *HttpError) Unwrap() error {
	return e.Err
}
