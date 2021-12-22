package controllers

import "github.com/gin-gonic/gin"

const StatusOK string = "OK"
const StatusError string = "ERROR"

type Response struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    *gin.H         `json:"data"`
	Error   *ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Data    *gin.H `json:"data"`
}
