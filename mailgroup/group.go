package mailgroup

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// CreateRequest 创建邮件组请求
type CreateRequest struct {
	Email             string `json:"email"`                        // 邮件组邮箱地址（必填）
	Name              string `json:"name"`                         // 邮件组名称（必填）
	Description       string `json:"description,omitempty"`        // 邮件组描述（可选）
	WhoCanSendMail    string `json:"who_can_send_mail,omitempty"` // 谁可以发送邮件：ALL_INTERNAL_USERS, ALL_GROUP_MEMBERS等（可选）
}

// CreateResponse 创建邮件组响应
type CreateResponse struct {
	MailGroup feishumail.MailGroup `json:"mailgroup"`
}

// Create 创建邮件组
// client: 飞书API客户端
// req: 创建请求
// 返回: 创建的邮件组信息
func Create(client *feishumail.Client, req CreateRequest) (*feishumail.MailGroup, error) {
	var resp CreateResponse
	if err := client.DoRequest("POST", "/mail/v1/mailgroups", req, &resp); err != nil {
		return nil, err
	}
	return &resp.MailGroup, nil
}

// Delete 删除邮件组
// client: 飞书API客户端
// mailgroupID: 邮件组ID
func Delete(client *feishumail.Client, mailgroupID string) error {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s", mailgroupID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// UpdatePartialRequest 修改邮件组部分信息请求
type UpdatePartialRequest struct {
	Name              string `json:"name,omitempty"`              // 邮件组名称（可选）
	Description       string `json:"description,omitempty"`      // 邮件组描述（可选）
	WhoCanSendMail    string `json:"who_can_send_mail,omitempty"` // 谁可以发送邮件（可选）
}

// UpdatePartial 修改邮件组部分信息
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 更新请求
func UpdatePartial(client *feishumail.Client, mailgroupID string, req UpdatePartialRequest) error {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s", mailgroupID)
	return client.DoRequest("PATCH", path, req, nil)
}

// UpdateFullRequest 修改邮件组全部信息请求
type UpdateFullRequest struct {
	Name              string `json:"name"`                        // 邮件组名称（必填）
	Description       string `json:"description,omitempty"`      // 邮件组描述（可选）
	WhoCanSendMail    string `json:"who_can_send_mail"`          // 谁可以发送邮件（必填）
}

// UpdateFull 修改邮件组全部信息
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 更新请求
func UpdateFull(client *feishumail.Client, mailgroupID string, req UpdateFullRequest) error {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s", mailgroupID)
	return client.DoRequest("PUT", path, req, nil)
}

// GetResponse 查询邮件组响应
type GetResponse struct {
	MailGroup feishumail.MailGroup `json:"mailgroup"`
}

// Get 查询指定邮件组
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// 返回: 邮件组信息
func Get(client *feishumail.Client, mailgroupID string) (*feishumail.MailGroup, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s", mailgroupID)
	var resp GetResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.MailGroup, nil
}

// ListRequest 批量获取邮件组请求
type ListRequest struct {
	PageToken string `json:"page_token,omitempty"` // 分页token（可选）
	PageSize  int    `json:"page_size,omitempty"`  // 每页数量（可选，默认20）
}

// ListResponse 批量获取邮件组响应
type ListResponse struct {
	Items    []feishumail.MailGroup `json:"items"`     // 邮件组列表
	PageInfo feishumail.PageInfo    `json:"page_info"` // 分页信息
}

// List 批量获取邮件组
// client: 飞书API客户端
// req: 查询请求
// 返回: 邮件组列表和分页信息
func List(client *feishumail.Client, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := client.DoRequest("GET", "/mail/v1/mailgroups", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
