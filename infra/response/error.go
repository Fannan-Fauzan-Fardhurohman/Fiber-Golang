package response

import (
	"errors"
	"net/http"
)

// Error general
var (
	ErrNotFound = errors.New("error not found")
)

var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 characters")
	ErrAuthIsNotExists       = errors.New("Auth is not exists")
	ErrEmailAlreadyUsed      = errors.New("Email already used")
	ErrPasswordNotMatch      = errors.New("Password not match")
)

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral    = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request error", "40000", http.StatusBadRequest)
)

var (
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorAuthIsNotExists       = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyUsed      = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch      = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
	ErrorNotFound              = NewError(ErrNotFound.Error(), "40401", http.StatusNotFound)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrAuthIsNotExists.Error():       ErrorAuthIsNotExists,
		ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
	}
)
