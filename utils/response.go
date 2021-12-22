package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const StatusOK string = "OK"
const StatusError string = "ERROR"

const ErrorCodeUnauthorized string = "01"
const ErrorCodeValidation string = "10"
const ErrorCodeDBConnection string = "30"
const ErrorCodeUnknown string = "99"

const ErrorMessageUnauthorized = "UNAUTHORIZED"
const ErrorMessageValidation = "VALIDATION ERROR"
const ErrorMessageDatabase = "DATABASE ERROR"
const ErrorMessageUnknown = "GENERAL ERROR"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   *Error      `json:"error"`
}

type Error struct {
	Code   string   `json:"code"`
	Errors []string `json:"errors"`
}

func SuccessResponse(c *gin.Context, h interface{}) {
	c.JSON(http.StatusOK, &Response{
		Status:  StatusOK,
		Message: "SUCCESS",
		Data:    h,
		Error:   nil,
	})
}

func BuildErrorResponse(c *gin.Context, errorCode string, errors []string) {
	c.AbortWithStatusJSON(parseHttpStatus(errorCode), &Response{
		Status:  StatusError,
		Message: parseErrorMessage(errorCode),
		Data:    nil,
		Error: &Error{
			Code:   errorCode,
			Errors: errors,
		},
	})
}

func ErrorResponse(c *gin.Context, code string, err error) {
	var errs []string
	errs = append(errs, err.Error())
	BuildErrorResponse(c, code, errs)
	c.Abort()
}

func parseErrorMessage(errorCode string) string {
	switch errorCode {
	case ErrorCodeUnauthorized:
		return ErrorMessageUnauthorized
	case ErrorCodeValidation:
		return ErrorMessageValidation
	case ErrorCodeDBConnection:
		return ErrorMessageDatabase
	case ErrorCodeUnknown:
	default:
		return ErrorMessageUnknown
	}
	return ErrorMessageUnknown
}

func parseHttpStatus(errorCode string) int {
	switch errorCode {
	case ErrorCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrorCodeValidation:
		return http.StatusBadRequest
	case ErrorCodeDBConnection:
		return http.StatusUnprocessableEntity
	case ErrorCodeUnknown:
	default:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}

func ByteToInterface(jsonData []byte) map[string]interface{} {
	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})

	return data
}

type RData struct {
	Code    int         `json:"-"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   *Error      `json:"error"`
}

type resp struct {
	c *gin.Context
}

func R(c *gin.Context) *resp {
	return &resp{c}
}

//const CodeSuccess = 0
//const CodeUnknowError = 99999999

func (r resp) Response(data *RData) {
	c := r.c
	c.JSON(data.Code, data)
}

func (r resp) Parse(v interface{}) *RData {
	data := &RData{
		Code:    http.StatusOK,
		Message: StatusOK,
		Status:  StatusOK,
	}
	switch v.(type) {
	case error:
		res := v.(error)
		data.Code = http.StatusInternalServerError
		data.Message = ErrorMessageUnknown
		data.Status = StatusError
		data.Data = nil
		data.Error = &Error{
			Code:   ErrorCodeUnknown,
			Errors: []string{res.Error()},
		}

	case string:
		data.Data = v.(string)

	case gin.H, *gin.H, map[string]interface{}:
		var e gin.H
		if b, ok := v.(gin.H); ok {
			e = b
		} else if b, ok := v.(map[string]interface{}); ok {
			e = gin.H(b)
		} else if b, ok := v.(*gin.H); ok {
			e = *b
		}

		resCode := e["code"]
		if resCode == nil {
			resCode = http.StatusOK
		}

		resStatus := e["status"]
		if resStatus == nil {
			resStatus = StatusOK
		}

		resMsg := e["msg"]
		if resMsg == nil {
			resMsg = "ok"
		} else if errmsgError, ok := resMsg.(error); ok {
			resMsg = errmsgError.Error()
		}

		resData := e["data"]
		if resData == nil {
			resData = gin.H{}
		}

		data = &RData{
			Code:    resCode.(int),
			Status:  resStatus.(string),
			Message: resMsg.(string),
			Data:    resData,
		}

	case RData, *RData:
		if b, ok := v.(RData); ok {
			data = &b
		} else {
			data = v.(*RData)
		}
	default:
		data.Data = v
	}

	return data
}

func (r resp) OK(v interface{}) {
	r.Response(&RData{
		Code:    http.StatusOK,
		Message: "ok",
		Status:  StatusOK,
		Data:    v,
	})
}

func (r resp) Res(v interface{}) {
	r.Response(r.Parse(v))
}

func (r resp) Err(v error) {
	r.Res(v)
}

func (r resp) Error(v interface{}) {
	data := r.Parse(v)
	if data.Code == 0 {
		data.Code = 500
	}
	if data.Code == http.StatusOK {
		data.Code = http.StatusInternalServerError
	}
	if data.Message == "ok" {
		if msg, ok := data.Data.(string); ok {
			data.Message = msg
			data.Data = gin.H{}
		} else {
			data.Message = "unknown error"
		}
	}
	if data.Status == StatusOK {
		data.Status = StatusError
	}
	r.Response(data)
}

func (r resp) Forbidden(v error) {
	data := r.Parse(v)
	data.Code = http.StatusForbidden
	r.Response(data)
}
