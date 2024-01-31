package ctl

import (
	"github.com/gin-gonic/gin"
	"micro-todolist/pkg/e"
)

type Response struct {
	Status int         `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func RespError(ctx *gin.Context, err error, data interface{}, code ...int) *Response {
	status := e.CodeError
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
		Error:  err.Error(),
	}
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := e.CodeSuccess
	if code != nil {
		status = code[0]
	}
	if data == nil {
		data = "操作成功"
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
}
