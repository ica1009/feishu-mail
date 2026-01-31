package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// GetAttachmentDownloadURLRequest 获取附件下载链接请求
type GetAttachmentDownloadURLRequest struct {
	UserMailboxID string   `json:"-"` // 用户邮箱ID（路径参数）
	MessageID     string   `json:"-"` // 邮件ID（路径参数）
	AttachmentIDs []string `json:"-"` // 附件ID列表（查询参数）
}

// AttachmentDownloadURL 附件下载链接信息
type AttachmentDownloadURL struct {
	AttachmentID string `json:"attachment_id"` // 附件ID
	DownloadURL  string `json:"download_url"`  // 下载链接
	ExpiresAt    string `json:"expires_at"`    // 过期时间
}

// GetAttachmentDownloadURLResponse 获取附件下载链接响应
type GetAttachmentDownloadURLResponse struct {
	Items []AttachmentDownloadURL `json:"items"` // 下载链接列表
}

// GetAttachmentDownloadURL 获取附件下载链接
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// messageID: 邮件ID
// attachmentIDs: 附件ID列表
// 返回: 附件下载链接列表
// 注意: 下载链接仅可使用两次，链接有效期两小时
func GetAttachmentDownloadURL(client *feishumail.Client, userMailboxID, messageID string, attachmentIDs []string) ([]AttachmentDownloadURL, error) {
	if len(attachmentIDs) == 0 {
		return nil, fmt.Errorf("attachment_ids cannot be empty")
	}
	
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages/%s/attachments/download_url", userMailboxID, messageID)
	
	// 构建查询参数
	params := "attachment_ids=" + attachmentIDs[0]
	for i := 1; i < len(attachmentIDs); i++ {
		params += "&attachment_ids=" + attachmentIDs[i]
	}
	path += "?" + params
	
	var resp GetAttachmentDownloadURLResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Items, nil
}

// GetAttachmentInfoRequest 获取附件信息请求
type GetAttachmentInfoRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	MessageID     string `json:"-"` // 邮件ID（路径参数）
	AttachmentID  string `json:"-"` // 附件ID（路径参数）
}

// GetAttachmentInfoResponse 获取附件信息响应
type GetAttachmentInfoResponse struct {
	Attachment Attachment `json:"attachment"` // 附件信息
}

// GetAttachmentInfo 获取附件信息
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// messageID: 邮件ID
// attachmentID: 附件ID
// 返回: 附件信息
func GetAttachmentInfo(client *feishumail.Client, userMailboxID, messageID, attachmentID string) (*Attachment, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/messages/%s/attachments/%s", userMailboxID, messageID, attachmentID)
	var resp GetAttachmentInfoResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Attachment, nil
}
