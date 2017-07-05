package bean

const (
	CODE_Success               = 200
	CODE_Created               = 201
	CODE_Bad_Request           = 400 // 请求错误
	CODE_Unauthorized          = 401 // 没有登录
	CODE_Not_Found             = 404 // not found
	CODE_Forbidden             = 403 // 没有权限
	CODE_Method_Not_Allowed    = 405 // 方法不对 (POST,PUT,GET)
	CODE_Not_Acceptable        = 406 // 不能通过
	CODE_Internal_Server_Error = 500 // 服务错误

	CODE_Params_Err = 430 // 参数错误
)

func CodeString(code int) string {
	s := map[int]string{
		CODE_Success:                      "OK",
		CODE_Created:                      "Created",
		CODE_Bad_Request:                  "Bad_Request",
		CODE_Unauthorized:                 "Unauthorized",
		CODE_Not_Found:                    "Not_Found",
		CODE_Forbidden:                    "Forbidden",
		CODE_Method_Not_Allowed:           "Method_Not_Allowed",
		CODE_Not_Acceptable:               "Not_Acceptable",
		CODE_Internal_Server_Error:        "Server_Error",
		CODE_Params_Err:                   "Params_Error",
	}[code]
	return s
}

