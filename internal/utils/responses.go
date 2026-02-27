package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data"`
	Message string `json:"message,omitempty"`
}

func SendError(c *gin.Context, err error) {
	if appErr := GetAppError(err); appErr != nil {
		c.JSON(appErr.StatusCode, ErrorResponse{
			Success: false,
			Error:   getErrorCode(appErr.StatusCode),
			Message: appErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Success: false,
		Error:   "internal_error",
		Message: "Внутренняя ошибка сервера",
	})
}

func SendSuccess[T any](c *gin.Context, data T, message string) {
	c.JSON(http.StatusOK, SuccessResponse[T]{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func SendSuccessWithStatus[T any](c *gin.Context, status int, data T, message string) {
	c.JSON(status, SuccessResponse[T]{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func getErrorCode(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return "bad_request"
	case http.StatusUnauthorized:
		return "unauthorized"
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not_found"
	case http.StatusConflict:
		return "conflict"
	case http.StatusInternalServerError:
		return "internal_error"
	default:
		return "error"
	}
}
