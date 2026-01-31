package publicmailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// AddMemberRequest 添加公共邮箱成员请求
type AddMemberRequest struct {
	UserID string `json:"user_id,omitempty"` // 用户ID（可选）
	Type   string `json:"type,omitempty"`     // 成员类型（可选）
}

// AddMemberResponse 添加公共邮箱成员响应
type AddMemberResponse struct {
	Member feishumail.PublicMailboxMember `json:"member"`
}

// AddMember 添加公共邮箱成员
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// req: 添加成员请求
// 返回: 添加的成员信息
func AddMember(client *feishumail.Client, publicMailboxID string, req AddMemberRequest) (*feishumail.PublicMailboxMember, error) {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/members", publicMailboxID)
	var resp AddMemberResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Member, nil
}

// DeleteMember 删除公共邮箱成员
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// memberID: 成员ID
func DeleteMember(client *feishumail.Client, publicMailboxID, memberID string) error {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/members/%s", publicMailboxID, memberID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// DeleteAllMembers 删除公共邮箱所有成员
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
func DeleteAllMembers(client *feishumail.Client, publicMailboxID string) error {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/members/clear", publicMailboxID)
	return client.DoRequest("POST", path, nil, nil)
}

// GetMemberResponse 获取公共邮箱成员信息响应
type GetMemberResponse struct {
	Member feishumail.PublicMailboxMember `json:"member"`
}

// GetMember 获取公共邮箱成员信息
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// memberID: 成员ID
// 返回: 成员信息
func GetMember(client *feishumail.Client, publicMailboxID, memberID string) (*feishumail.PublicMailboxMember, error) {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/members/%s", publicMailboxID, memberID)
	var resp GetMemberResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Member, nil
}

// ListMembersRequest 查询所有公共邮箱成员信息请求
type ListMembersRequest struct {
	PageToken string `json:"page_token,omitempty"` // 分页token（可选）
	PageSize  int    `json:"page_size,omitempty"`  // 每页数量（可选）
}

// ListMembersResponse 查询所有公共邮箱成员信息响应
type ListMembersResponse struct {
	Items    []feishumail.PublicMailboxMember `json:"items"`     // 成员列表
	PageInfo feishumail.PageInfo             `json:"page_info"` // 分页信息
}

// ListMembers 查询所有公共邮箱成员信息
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// req: 查询请求
// 返回: 成员列表和分页信息
func ListMembers(client *feishumail.Client, publicMailboxID string, req ListMembersRequest) (*ListMembersResponse, error) {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/members", publicMailboxID)
	var resp ListMembersResponse
	if err := client.DoRequest("GET", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchAddMembersRequest 批量添加公共邮箱成员请求
type BatchAddMembersRequest struct {
	Items []AddMemberRequest `json:"items"` // 成员列表
}

// BatchAddMembersResponse 批量添加公共邮箱成员响应
type BatchAddMembersResponse struct {
	Items []feishumail.PublicMailboxMember `json:"items"` // 添加的成员列表
}

// BatchAddMembers 批量添加公共邮箱成员
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// req: 批量添加请求
// 返回: 添加的成员列表
func BatchAddMembers(client *feishumail.Client, publicMailboxID string, req BatchAddMembersRequest) ([]feishumail.PublicMailboxMember, error) {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/members/batch_add", publicMailboxID)
	var resp BatchAddMembersResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return resp.Items, nil
}

// BatchDeleteMembersRequest 批量删除公共邮箱成员请求
type BatchDeleteMembersRequest struct {
	MemberIDs []string `json:"member_ids"` // 成员ID列表
}

// BatchDeleteMembers 批量删除公共邮箱成员
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// req: 批量删除请求
func BatchDeleteMembers(client *feishumail.Client, publicMailboxID string, req BatchDeleteMembersRequest) error {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/members/batch_delete", publicMailboxID)
	return client.DoRequest("POST", path, req, nil)
}
