package pkg

var (
	ErrInternalServer  = &Response{Code: 500, Msg: "服务器发生错误"}
	UnauthorizedAuthFail = &Response{Code: 401, Msg: "权限验证不通过"}
)


