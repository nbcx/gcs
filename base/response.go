package base

const (
	OK                 = 200  // Success
	NotLoggedIn        = 1000 // 未登录
	ParameterIllegal   = 1001 // 参数不合法
	UnauthorizedUserId = 1002 // 非法的用户Id
	Unauthorized       = 1003 // 未授权
	ServerError        = 1004 // 系统错误
	NotData            = 1005 // 没有数据
	ModelAddError      = 1006 // 添加错误
	ModelDeleteError   = 1007 // 删除错误
	ModelStoreError    = 1008 // 存储错误
	OperationFailure   = 1009 // 操作失败
	RoutingNotExist    = 1010 // 路由不存在
)

type JsonResult struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(code uint32, message string, data interface{}) JsonResult {

	message = GetErrorMessage(code, message)
	jsonMap := grantMap(code, message, data)

	return jsonMap
}

// 按照接口格式生成原数据数组
func grantMap(code uint32, message string, data interface{}) JsonResult {

	jsonMap := JsonResult{
		Code: code,
		Msg:  message,
		Data: data,
	}
	return jsonMap
}

// 根据错误码 获取错误信息
func GetErrorMessage(code uint32, message string) string {
	var codeMessage string
	codeMap := map[uint32]string{
		OK:                 "Success",
		NotLoggedIn:        "未登录",
		ParameterIllegal:   "参数不合法",
		UnauthorizedUserId: "非法的用户Id",
		Unauthorized:       "未授权",
		NotData:            "没有数据",
		ServerError:        "系统错误",
		ModelAddError:      "添加错误",
		ModelDeleteError:   "删除错误",
		ModelStoreError:    "存储错误",
		OperationFailure:   "操作失败",
		RoutingNotExist:    "路由不存在",
	}

	if message == "" {
		if value, ok := codeMap[code]; ok {
			// 存在
			codeMessage = value
		} else {
			codeMessage = "未定义错误类型!"
		}
	} else {
		codeMessage = message
	}

	return codeMessage
}
