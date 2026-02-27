package utils

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code       int    `json:"code"`
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewAppError(statusCode int, message string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func NewAppErrorWithErr(statusCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

var (
	ErrUnauthorized     = &AppError{StatusCode: http.StatusUnauthorized, Message: "Не авторизован"}
	ErrForbidden        = &AppError{StatusCode: http.StatusForbidden, Message: "Доступ запрещен"}
	ErrNotFound         = &AppError{StatusCode: http.StatusNotFound, Message: "Ресурс не найден"}
	ErrBadRequest       = &AppError{StatusCode: http.StatusBadRequest, Message: "Неверный запрос"}
	ErrConflict         = &AppError{StatusCode: http.StatusConflict, Message: "Ресурс уже существует"}
	ErrInternalServer   = &AppError{StatusCode: http.StatusInternalServerError, Message: "Внутренняя ошибка сервера"}
	ErrValidationFailed = &AppError{StatusCode: http.StatusBadRequest, Message: "Ошибка валидации"}
)

func IsAppError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr)
}

func GetAppError(err error) *AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}
	return nil
}
