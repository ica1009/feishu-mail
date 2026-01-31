package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// SendMailRequest 发送邮件请求
type SendMailRequest struct {
	UserMailboxID string   `json:"-"` // 用户邮箱ID（路径参数）
	To            []string `json:"to"` // 收件人列表（必填）
	Cc            []string `json:"cc,omitempty"` // 抄送列表（可选）
	Bcc           []string `json:"bcc,omitempty"` // 密送列表（可选）
	Subject       string   `json:"subject"` // 邮件主题（必填）
	Body          string   `json:"body,omitempty"` // 邮件正文（可选）
	BodyHTML      string   `json:"body_html,omitempty"` // 邮件正文HTML（可选）
	Attachments   []string `json:"attachments,omitempty"` // 附件ID列表（可选）
	ReplyTo       string   `json:"reply_to,omitempty"` // 回复地址（可选）
	InReplyTo     string   `json:"in_reply_to,omitempty"` // 回复的邮件ID（可选）
}

// SendMailResponse 发送邮件响应
type SendMailResponse struct {
	MessageID string `json:"message_id"` // 发送的邮件ID
}

// SendMail 发送邮件
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 发送邮件请求
// 返回: 发送的邮件ID
func SendMail(client *feishumail.Client, userMailboxID string, req SendMailRequest) (string, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages/send", userMailboxID)
	var resp SendMailResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return "", err
	}
	return resp.MessageID, nil
}
