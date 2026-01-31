package usermailbox

import (
	"github.com/ica1009/feishu-mail"
)

// ResetPasswordRequest 重置用户的企业邮箱密码请求
type ResetPasswordRequest struct {
	UserID  string `json:"user_id"`  // 用户ID（必填）
	NewPassword string `json:"new_password,omitempty"` // 新密码（可选，不填则系统自动生成）
}

// ResetPasswordResponse 重置用户的企业邮箱密码响应
type ResetPasswordResponse struct {
	Password string `json:"password,omitempty"` // 新密码（如果系统自动生成）
}

// ResetPassword 重置用户的企业邮箱密码
// client: 飞书API客户端
// req: 重置密码请求
// 返回: 重置密码响应（如果系统自动生成密码）
func ResetPassword(client *feishumail.Client, req ResetPasswordRequest) (*ResetPasswordResponse, error) {
	var resp ResetPasswordResponse
	if err := client.DoRequest("POST", "/mail/v1/user_mailboxes/reset_password", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
