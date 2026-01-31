package feishumail

import "fmt"

// Error 飞书API错误
type Error struct {
	Code    int    // 错误码
	Message string // 错误信息
}

// Error 实现error接口
func (e *Error) Error() string {
	return fmt.Sprintf("feishu api error: code=%d, msg=%s", e.Code, e.Message)
}

// NewError 创建新的错误
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// IsSuccess 检查响应是否成功
func IsSuccess(code int) bool {
	return code == 0
}
