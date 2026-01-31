package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// MailMessage 邮件消息信息
type MailMessage struct {
	MessageID      string   `json:"message_id"`       // 邮件ID
	Subject        string   `json:"subject"`          // 邮件主题
	From           string   `json:"from"`             // 发件人
	To             []string `json:"to"`               // 收件人列表
	Cc             []string `json:"cc,omitempty"`     // 抄送列表
	Bcc            []string `json:"bcc,omitempty"`    // 密送列表
	Body           string   `json:"body,omitempty"`   // 邮件正文
	BodyHTML       string   `json:"body_html,omitempty"` // 邮件正文HTML
	ReceivedTime   string   `json:"received_time"`    // 接收时间
	SentTime       string   `json:"sent_time"`        // 发送时间
	HasAttachments bool     `json:"has_attachments"`  // 是否有附件
	Attachments    []Attachment `json:"attachments,omitempty"` // 附件列表
	IsRead         bool     `json:"is_read"`          // 是否已读
	IsStarred      bool     `json:"is_starred"`       // 是否星标
	Labels         []string `json:"labels,omitempty"` // 标签列表
}

// Attachment 邮件附件信息
type Attachment struct {
	AttachmentID string `json:"attachment_id"` // 附件ID
	Name         string `json:"name"`          // 附件名称
	Size         int64  `json:"size"`         // 附件大小（字节）
	ContentType  string `json:"content_type"`  // 内容类型
}

// GetMessageRequest 获取邮件详情请求
type GetMessageRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	MessageID     string `json:"-"` // 邮件ID（路径参数）
	IncludeBody   bool   `json:"include_body,omitempty"` // 是否包含邮件正文（查询参数）
}

// GetMessageResponse 获取邮件详情响应
type GetMessageResponse struct {
	Message MailMessage `json:"message"` // 邮件信息
}

// GetMessage 获取邮件详情
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// messageID: 邮件ID
// includeBody: 是否包含邮件正文
// 返回: 邮件详情
func GetMessage(client *feishumail.Client, userMailboxID, messageID string, includeBody bool) (*MailMessage, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages/%s", userMailboxID, messageID)
	if includeBody {
		path += "?include_body=true"
	}
	var resp GetMessageResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Message, nil
}

// ListMessagesRequest 获取邮件列表请求
type ListMessagesRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	PageToken     string `json:"page_token,omitempty"` // 分页token（查询参数）
	PageSize      int    `json:"page_size,omitempty"`  // 每页数量（查询参数，默认20）
	FolderID      string `json:"folder_id,omitempty"`  // 文件夹ID（查询参数，如"inbox"）
	IsRead        *bool  `json:"is_read,omitempty"`    // 是否已读（查询参数）
	IsStarred     *bool  `json:"is_starred,omitempty"` // 是否星标（查询参数）
}

// ListMessagesResponse 获取邮件列表响应
type ListMessagesResponse struct {
	Items    []MailMessage      `json:"items"`     // 邮件列表
	PageInfo feishumail.PageInfo `json:"page_info"` // 分页信息
}

// ListMessages 获取邮件列表
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 查询请求
// 返回: 邮件列表和分页信息
func ListMessages(client *feishumail.Client, userMailboxID string, req ListMessagesRequest) (*ListMessagesResponse, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages", userMailboxID)
	
	// 构建查询参数
	params := make([]string, 0)
	if req.PageToken != "" {
		params = append(params, "page_token="+req.PageToken)
	}
	if req.PageSize > 0 {
		params = append(params, fmt.Sprintf("page_size=%d", req.PageSize))
	}
	if req.FolderID != "" {
		params = append(params, "folder_id="+req.FolderID)
	}
	if req.IsRead != nil {
		params = append(params, fmt.Sprintf("is_read=%v", *req.IsRead))
	}
	if req.IsStarred != nil {
		params = append(params, fmt.Sprintf("is_starred=%v", *req.IsStarred))
	}
	
	if len(params) > 0 {
		path += "?" + params[0]
		for i := 1; i < len(params); i++ {
			path += "&" + params[i]
		}
	}
	
	var resp ListMessagesResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
