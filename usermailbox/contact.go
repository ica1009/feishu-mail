package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// MailContact 邮箱联系人信息
type MailContact struct {
	ContactID string `json:"contact_id"` // 联系人ID
	Name      string `json:"name"`        // 联系人名称
	Email     string `json:"email,omitempty"` // 邮箱地址（敏感字段，需要权限）
	Phone     string `json:"phone,omitempty"` // 手机号（敏感字段，需要权限）
	Company   string `json:"company,omitempty"` // 公司
	Title     string `json:"title,omitempty"` // 职位
}

// ListContactsRequest 列出邮箱联系人请求
type ListContactsRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	PageToken     string `json:"page_token,omitempty"` // 分页token（查询参数）
	PageSize      int    `json:"page_size,omitempty"`  // 每页数量（查询参数）
	Search        string `json:"search,omitempty"`     // 搜索关键词（查询参数）
}

// ListContactsResponse 列出邮箱联系人响应
type ListContactsResponse struct {
	Items    []MailContact       `json:"items"`     // 联系人列表
	PageInfo feishumail.PageInfo `json:"page_info"` // 分页信息
}

// ListContacts 列出邮箱联系人
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 查询请求
// 返回: 联系人列表和分页信息
func ListContacts(client *feishumail.Client, userMailboxID string, req ListContactsRequest) (*ListContactsResponse, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/mail_contacts", userMailboxID)
	
	// 构建查询参数
	params := make([]string, 0)
	if req.PageToken != "" {
		params = append(params, "page_token="+req.PageToken)
	}
	if req.PageSize > 0 {
		params = append(params, fmt.Sprintf("page_size=%d", req.PageSize))
	}
	if req.Search != "" {
		params = append(params, "search="+req.Search)
	}
	
	if len(params) > 0 {
		path += "?" + params[0]
		for i := 1; i < len(params); i++ {
			path += "&" + params[i]
		}
	}
	
	var resp ListContactsResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateContactRequest 创建邮箱联系人请求
type CreateContactRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	Name          string `json:"name"` // 联系人名称（必填）
	Email         string `json:"email,omitempty"` // 邮箱地址（可选）
	Phone         string `json:"phone,omitempty"` // 手机号（可选）
	Company       string `json:"company,omitempty"` // 公司（可选）
	Title         string `json:"title,omitempty"` // 职位（可选）
}

// CreateContactResponse 创建邮箱联系人响应
type CreateContactResponse struct {
	Contact MailContact `json:"contact"` // 联系人信息
}

// CreateContact 创建邮箱联系人
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 创建请求
// 返回: 创建的联系人信息
func CreateContact(client *feishumail.Client, userMailboxID string, req CreateContactRequest) (*MailContact, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/mail_contacts", userMailboxID)
	var resp CreateContactResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Contact, nil
}

// UpdateContactRequest 更新邮箱联系人请求
type UpdateContactRequest struct {
	Name    string `json:"name,omitempty"`    // 联系人名称（可选）
	Email   string `json:"email,omitempty"`   // 邮箱地址（可选）
	Phone   string `json:"phone,omitempty"`   // 手机号（可选）
	Company string `json:"company,omitempty"` // 公司（可选）
	Title   string `json:"title,omitempty"`   // 职位（可选）
}

// UpdateContactResponse 更新邮箱联系人响应
type UpdateContactResponse struct {
	Contact MailContact `json:"contact"` // 联系人信息
}

// UpdateContact 更新邮箱联系人
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// contactID: 联系人ID
// req: 更新请求
// 返回: 更新后的联系人信息
func UpdateContact(client *feishumail.Client, userMailboxID, contactID string, req UpdateContactRequest) (*MailContact, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/mail_contacts/%s", userMailboxID, contactID)
	var resp UpdateContactResponse
	if err := client.DoRequest("PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Contact, nil
}

// DeleteContact 删除邮箱联系人
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// contactID: 联系人ID
func DeleteContact(client *feishumail.Client, userMailboxID, contactID string) error {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/mail_contacts/%s", userMailboxID, contactID)
	return client.DoRequest("DELETE", path, nil, nil)
}
