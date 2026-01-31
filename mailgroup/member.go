package mailgroup

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// CreateMemberRequest 创建邮件组成员请求
type CreateMemberRequest struct {
	UserID string `json:"user_id,omitempty"` // 用户ID（可选）
	Email  string `json:"email,omitempty"`    // 成员邮箱（可选）
	Type   string `json:"type,omitempty"`     // 成员类型：USER, DEPARTMENT等（可选）
}

// CreateMemberResponse 创建邮件组成员响应
type CreateMemberResponse struct {
	Member feishumail.MailGroupMember `json:"member"`
}

// CreateMember 创建邮件组成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 创建成员请求
// 返回: 创建的成员信息
func CreateMember(client *feishumail.Client, mailgroupID string, req CreateMemberRequest) (*feishumail.MailGroupMember, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/members", mailgroupID)
	var resp CreateMemberResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Member, nil
}

// DeleteMember 删除邮件组成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// memberID: 成员ID
func DeleteMember(client *feishumail.Client, mailgroupID, memberID string) error {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/members/%s", mailgroupID, memberID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// GetMemberRequest 查询邮件组成员请求
type GetMemberRequest struct {
	UserID string `json:"user_id,omitempty"` // 用户ID（可选）
	Email  string `json:"email,omitempty"`  // 成员邮箱（可选）
}

// GetMemberResponse 查询邮件组成员响应
type GetMemberResponse struct {
	Member feishumail.MailGroupMember `json:"member"`
}

// GetMember 查询指定邮件组成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 查询请求
// 返回: 成员信息
func GetMember(client *feishumail.Client, mailgroupID string, req GetMemberRequest) (*feishumail.MailGroupMember, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/members", mailgroupID)
	var resp GetMemberResponse
	if err := client.DoRequest("GET", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Member, nil
}

// ListMembersRequest 获取所有邮件组成员请求
type ListMembersRequest struct {
	PageToken string `json:"page_token,omitempty"` // 分页token（可选）
	PageSize  int    `json:"page_size,omitempty"`  // 每页数量（可选）
}

// ListMembersResponse 获取所有邮件组成员响应
type ListMembersResponse struct {
	Items    []feishumail.MailGroupMember `json:"items"`     // 成员列表
	PageInfo feishumail.PageInfo          `json:"page_info"` // 分页信息
}

// ListMembers 获取所有邮件组成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 查询请求
// 返回: 成员列表和分页信息
func ListMembers(client *feishumail.Client, mailgroupID string, req ListMembersRequest) (*ListMembersResponse, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/members", mailgroupID)
	var resp ListMembersResponse
	if err := client.DoRequest("GET", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchCreateMembersRequest 批量创建邮件组成员请求
type BatchCreateMembersRequest struct {
	Items []CreateMemberRequest `json:"items"` // 成员列表
}

// BatchCreateMembersResponse 批量创建邮件组成员响应
type BatchCreateMembersResponse struct {
	Items []feishumail.MailGroupMember `json:"items"` // 创建的成员列表
}

// BatchCreateMembers 批量创建邮件组成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 批量创建请求
// 返回: 创建的成员列表
func BatchCreateMembers(client *feishumail.Client, mailgroupID string, req BatchCreateMembersRequest) ([]feishumail.MailGroupMember, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/members/batch_create", mailgroupID)
	var resp BatchCreateMembersResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return resp.Items, nil
}

// BatchDeleteMembersRequest 批量删除邮件组成员请求
type BatchDeleteMembersRequest struct {
	MemberIDs []string `json:"member_ids"` // 成员ID列表
}

// BatchDeleteMembers 批量删除邮件组成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 批量删除请求
func BatchDeleteMembers(client *feishumail.Client, mailgroupID string, req BatchDeleteMembersRequest) error {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/members/batch_delete", mailgroupID)
	return client.DoRequest("POST", path, req, nil)
}
