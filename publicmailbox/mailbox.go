package publicmailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// CreateRequest 创建公共邮箱请求
type CreateRequest struct {
	Email       string `json:"email"`                 // 公共邮箱地址（必填）
	Name        string `json:"name"`                  // 公共邮箱名称（必填）
	Description string `json:"description,omitempty"` // 公共邮箱描述（可选）
}

// CreateResponse 创建公共邮箱响应
type CreateResponse struct {
	PublicMailbox feishumail.PublicMailbox `json:"public_mailbox"`
}

// Create 创建公共邮箱
// client: 飞书API客户端
// req: 创建请求
// 返回: 创建的公共邮箱信息
func Create(client *feishumail.Client, req CreateRequest) (*feishumail.PublicMailbox, error) {
	var resp CreateResponse
	if err := client.DoRequest("POST", "/mail/v1/public_mailboxes", req, &resp); err != nil {
		return nil, err
	}
	return &resp.PublicMailbox, nil
}

// Delete 删除公共邮箱
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
func Delete(client *feishumail.Client, publicMailboxID string) error {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s", publicMailboxID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// UpdateRequest 修改公共邮箱请求
type UpdateRequest struct {
	Name        string `json:"name,omitempty"`        // 公共邮箱名称（可选）
	Description string `json:"description,omitempty"` // 公共邮箱描述（可选）
}

// Update 修改公共邮箱
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// req: 更新请求
func Update(client *feishumail.Client, publicMailboxID string, req UpdateRequest) error {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s", publicMailboxID)
	return client.DoRequest("PATCH", path, req, nil)
}

// UpdateFullRequest 修改公共邮箱全部信息请求
type UpdateFullRequest struct {
	Name        string `json:"name"`                  // 公共邮箱名称（必填）
	Description string `json:"description,omitempty"` // 公共邮箱描述（可选）
}

// UpdateFull 修改公共邮箱全部信息
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// req: 更新请求
func UpdateFull(client *feishumail.Client, publicMailboxID string, req UpdateFullRequest) error {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s", publicMailboxID)
	return client.DoRequest("PUT", path, req, nil)
}

// ListRequest 查询所有公共邮箱请求
type ListRequest struct {
	PageToken string `json:"page_token,omitempty"` // 分页token（可选）
	PageSize  int    `json:"page_size,omitempty"`  // 每页数量（可选）
}

// ListResponse 查询所有公共邮箱响应
type ListResponse struct {
	Items    []feishumail.PublicMailbox `json:"items"`     // 公共邮箱列表
	PageInfo feishumail.PageInfo       `json:"page_info"` // 分页信息
}

// List 查询所有公共邮箱
// client: 飞书API客户端
// req: 查询请求
// 返回: 公共邮箱列表和分页信息
func List(client *feishumail.Client, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := client.DoRequest("GET", "/mail/v1/public_mailboxes", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResponse 查询指定公共邮箱响应
type GetResponse struct {
	PublicMailbox feishumail.PublicMailbox `json:"public_mailbox"`
}

// Get 查询指定公共邮箱
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// 返回: 公共邮箱信息
func Get(client *feishumail.Client, publicMailboxID string) (*feishumail.PublicMailbox, error) {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s", publicMailboxID)
	var resp GetResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.PublicMailbox, nil
}
