package fiber_errors

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"international-trading-test/internal/config"
)

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Extra   []string `json:"extra"`
}

func (e *Error) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func BadRequest(err ...interface{}) *fiber.Error {
	return NewError(fiber.StatusBadRequest, err...)
}

func Unauthorized(err ...interface{}) *fiber.Error {
	return NewError(fiber.StatusUnauthorized, err...)
}

func Forbidden(err ...interface{}) *fiber.Error {
	return NewError(fiber.StatusForbidden, err...)
}

func InternalServerError(err ...interface{}) *fiber.Error {

	return NewError(fiber.StatusInternalServerError, err...)
}

func NewError(code int, err ...interface{}) *fiber.Error {
	var message []string
	if code == fiber.StatusInternalServerError {
		message = append(message, fiber.NewError(code).Message)
	}
	for _, e := range err {
		message = append(message, fmt.Sprintf("%v", e))
	}

	var firstMessage string
	var extra []string
	if len(message) > 0 {
		firstMessage = message[0]
		if config.GetConfig().IsDebug {
			extra = message[1:]
		}
	} else {
		firstMessage = fiber.NewError(code).Message
	}

	formattedErr := Error{
		Code:    code,
		Message: firstMessage,
		Extra:   extra,
	}

	return fiber.NewError(code, formattedErr.Error())
}
