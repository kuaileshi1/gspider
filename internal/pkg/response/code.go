// @Title 返回状态统一定义
// @Description 请求状态码定义
// @Author shigx 2021/11/26 5:29 下午
package response

const (
	SuccessCode      = 200
	RedirectCode     = 301 // 重定向
	FailureCode      = 400
	UnauthorizedCode = 401
	ParamsErrorCode  = 402
	ForbiddenCode    = 403
	NotFoundCode     = 404
	ServerErrorCode  = 500
)

var CodeMessage = map[int]string{
	SuccessCode:      "成功",
	RedirectCode:     "",
	FailureCode:      "失败",
	UnauthorizedCode: "用户未登录授权",
	ParamsErrorCode:  "请求参数有误",
	ForbiddenCode:    "请求被拒绝",
	NotFoundCode:     "资源未找到",
	ServerErrorCode:  "服务器错误",
}

// @Description 返回状态描述信息
// @Auth shigx
// @Date 2021/11/26 5:45 下午
// @param code int 状态
// @return string 状态描述
func GetMessage(code int) string {
	msg, ok := CodeMessage[code]
	if ok {
		return msg
	}

	return "未知"
}
