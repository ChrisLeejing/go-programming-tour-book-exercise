package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	// remember the json tag.
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

// implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s.", e.Code(), e.Msg())
}

var codes = make(map[int]string)

func NewError(code int, msg string) *Error {
	// remove duplicate error code.
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码: %d 已存在, 请重新更换一个. ", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

func (e *Error) SetDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
