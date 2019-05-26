package common

import (
	"net/http"
	"runtime"
	"strconv"
)

//値オブジェクトではないが、ファイルが散らかるので...

type ErrorType string

const (
	NotFound            ErrorType = "not_found"
	UnAuthorized        ErrorType = "unauthorized"
	BadRequest          ErrorType = "bad_request"
	InternalServerError ErrorType = "internal_server_error"
)

const (
	ErrorParseError   ErrorType = "parse_error"
	ErrorBindingError ErrorType = "binding_error"
)

var callerSkipCount = 3

type CommonError struct {
	StatusCode int
	Msg        string
	StackTrace string
	ErrorType
}

type ErrorResponse struct {
	Error     string
	ErrorType ErrorType
}

func ErrorByStatusCode(statusCode int, msg string, errorType ErrorType) CommonError {
	switch statusCode {
	case http.StatusNotFound:
		return ErrorNotFound(msg, errorType)
	case http.StatusBadRequest:
		return ErrorBadRequest(msg, errorType)
	default:
		return ErrorInternalServerError("", "")
	}
}

func NewCommonError(statusCode int, msg string, errorType ErrorType) CommonError {
	var s = ""
	for i := 2; i >= 0; i-- {
		_, file, line, _ := runtime.Caller(callerSkipCount + i)
		s = s + file + ":" + strconv.Itoa(line) + " "
	}
	return CommonError{
		StatusCode: statusCode,
		Msg:        msg,
		StackTrace: s,
		ErrorType:  errorType,
	}
}

func ErrorNotFound(msg string, errType ErrorType) CommonError {
	return NewCommonError(http.StatusNotFound, msg, errType)
}

func ErrorUnauthorized(msg string, errType ErrorType) CommonError {
	return NewCommonError(http.StatusUnauthorized, msg, errType)
}

func ErrorBadRequest(msg string, errType ErrorType) CommonError {
	return NewCommonError(http.StatusBadRequest, msg, errType)
}

func ErrorInternalServerError(msg string, errType ErrorType) CommonError {
	return NewCommonError(http.StatusInternalServerError, msg, errType)
}

func BindingError(msg string) CommonError {
	return ErrorBadRequest(msg, ErrorBindingError)
}

func ParseError(msg string) CommonError {
	return ErrorBadRequest(msg, ErrorParseError)
}

func (e CommonError) Error() string {
	return e.Msg
}

// 内部向けのスタックトレースとかを表示する
func (e CommonError) InternalErrorJson() map[string]interface{} {
	json := map[string]interface{}{}
	json["type"] = string(e.ErrorType)
	json["msg"] = e.Msg
	json["stackTrace"] = e.StackTrace
	return json
}
