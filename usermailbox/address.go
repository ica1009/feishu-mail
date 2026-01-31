package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// DeleteFromRecycleBinRequest 从回收站删除用户邮箱地址请求
type DeleteFromRecycleBinRequest struct {
	TransferMailbox string `json:"transfer_mailbox,omitempty"` // 转移邮箱地址（可选，用于转移邮件）
}

// DeleteFromRecycleBin 从回收站删除用户邮箱地址（永久删除，无法恢复）
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID
// req: 删除请求（可选，用于转移邮件）
func DeleteFromRecycleBin(client *feishumail.Client, userMailboxID string, req *DeleteFromRecycleBinRequest) error {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s", userMailboxID)
	if req == nil {
		return client.DoRequest("DELETE", path, nil, nil)
	}
	return client.DoRequest("DELETE", path, req, nil)
}

// QueryStatusRequest 查询邮箱地址状态请求
type QueryStatusRequest struct {
	Email string `json:"email,omitempty"` // 邮箱地址（可选）
}

// QueryStatusResponse 查询邮箱地址状态响应
type QueryStatusResponse struct {
	Status feishumail.MailAddressStatus `json:"status"` // 邮箱地址状态
}

// QueryStatus 查询邮箱地址状态
// client: 飞书API客户端
// req: 查询请求
// 返回: 邮箱地址状态
func QueryStatus(client *feishumail.Client, req QueryStatusRequest) (*feishumail.MailAddressStatus, error) {
	var resp QueryStatusResponse
	if err := client.DoRequest("GET", "/mail/v1/user_mailboxes/status", req, &resp); err != nil {
		return nil, err
	}
	return &resp.Status, nil
}
