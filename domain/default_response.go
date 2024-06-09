package domain

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func NewDefaultResponse(code int, messages ...string) DefaultResponse {
	//使用变参来实现省略msg参数的功能
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	} else {
		switch code {
		case 200:
			message = "请求成功"
		case 400:
			message = "请求异常"
		case 500:
			message = "服务器异常"
		case 404:
			message = "资源不存在"
		case 403:
			message = "权限不足"
		default:
			message = "未知错误"
		}
	}

	return DefaultResponse{
		Code:    code,
		Message: message,
	}
}
