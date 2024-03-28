package apperr

import (
	"errors"
	"strings"
)

type errorCode int

type Error interface {
	error
	Unwrap() error
	SystemCode() errorCode
	Message() string
	Description() string
}

var _ Error = (*appError)(nil)

type appError struct {
	err         error
	code        errorCode
	message     string
	description string
}

func (e *appError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return text(e.code)
}

func (e *appError) Unwrap() error {
	return e.err
}

func (e *appError) Message() string {
	if e.message != "" {
		return text(e.code) + ": " + e.message
	}
	return text(e.code)
}

func (e *appError) Description() string {
	return e.description
}

func (e *appError) SystemCode() errorCode {
	return e.code
}

type appErrorOption func(*appError)

func WithError(err ...error) appErrorOption {
	return func(e *appError) {
		e.err = errors.Join(err...)
	}
}

func WithMessage(message ...string) appErrorOption {
	return func(e *appError) {
		e.message = strings.Join(message, " ")
	}
}

func WithDescription(desc ...string) appErrorOption {
	return func(e *appError) {
		e.description = strings.Join(desc, " ")
	}
}

func New(code errorCode, opts ...appErrorOption) *appError {
	e := &appError{code: code}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func GetErrorMessage(err error) string {
	var apperr Error
	if errors.As(err, &apperr) {
		return apperr.Message()
	}
	return unknownErrorText
}

func GetDescription(err error) string {
	var apperr Error
	if errors.As(err, &apperr) {
		return apperr.Description()
	}
	return ""
}

func GetSystemCode(err error) errorCode {
	var ae *appError
	if errors.As(err, &ae) {
		return ae.SystemCode()
	}
	return unknownErrorCode
}
