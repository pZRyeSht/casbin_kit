package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	HttpStatus int         `json:"-"`
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
}

func (r *Response) Error() string {
	return r.Msg
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

// NewFailRes 创建响应错误
func NewFailRes(msg string, args ...interface{}) error {
	return NewRes(200, ERROR, msg, args...)
}

// NewRes 创建响应
func NewRes(httpStatus, code int, msg string, args ...interface{}) error {
	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}
	res := &Response{
		HttpStatus: httpStatus,
		Code:       code,
		Msg:        fmt.Sprintf(msg, args...),
	}
	return res
}

// SprintfResponseMsg 格式化错误信息
func SprintfResponseMsg(res *Response, args ...interface{}) *Response {
	if res.HttpStatus == 0 {
		res.HttpStatus = http.StatusOK
	}
	res.Msg = fmt.Sprintf(res.Msg, args...)
	return res
}
