package mailgroup

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// CreatePermissionMemberRequest 创建邮件组权限成员请求
type CreatePermissionMemberRequest struct {
	UserID string `json:"user_id,omitempty"` // 用户ID（可选）
	Email  string `json:"email,omitempty"`    // 成员邮箱（可选）
	Type   string `json:"type,omitempty"`     // 成员类型（可选）
}

// CreatePermissionMemberResponse 创建邮件组权限成员响应
type CreatePermissionMemberResponse struct {
	Member feishumail.MailGroupMember `json:"member"`
}

// CreatePermissionMember 创建邮件组权限成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 创建权限成员请求
// 返回: 创建的权限成员信息
func CreatePermissionMember(client *feishumail.Client, mailgroupID string, req CreatePermissionMemberRequest) (*feishumail.MailGroupMember, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/permission_members", mailgroupID)
	var resp CreatePermissionMemberResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Member, nil
}

// DeletePermissionMember 删除邮件组权限成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// permissionMemberID: 权限成员ID
func DeletePermissionMember(client *feishumail.Client, mailgroupID, permissionMemberID string) error {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/permission_members/%s", mailgroupID, permissionMemberID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// GetPermissionMemberResponse 获取邮件组权限成员响应
type GetPermissionMemberResponse struct {
	Member feishumail.MailGroupMember `json:"member"`
}

// GetPermissionMember 获取邮件组权限成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// permissionMemberID: 权限成员ID
// 返回: 权限成员信息
func GetPermissionMember(client *feishumail.Client, mailgroupID, permissionMemberID string) (*feishumail.MailGroupMember, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/permission_members/%s", mailgroupID, permissionMemberID)
	var resp GetPermissionMemberResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Member, nil
}

// ListPermissionMembersRequest 批量获取邮件组权限成员请求
type ListPermissionMembersRequest struct {
	PageToken string `json:"page_token,omitempty"` // 分页token（可选）
	PageSize  int    `json:"page_size,omitempty"`  // 每页数量（可选）
}

// ListPermissionMembersResponse 批量获取邮件组权限成员响应
type ListPermissionMembersResponse struct {
	Items    []feishumail.MailGroupMember `json:"items"`     // 权限成员列表
	PageInfo feishumail.PageInfo          `json:"page_info"` // 分页信息
}

// ListPermissionMembers 批量获取邮件组权限成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 查询请求
// 返回: 权限成员列表和分页信息
func ListPermissionMembers(client *feishumail.Client, mailgroupID string, req ListPermissionMembersRequest) (*ListPermissionMembersResponse, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/permission_members", mailgroupID)
	var resp ListPermissionMembersResponse
	if err := client.DoRequest("GET", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchCreatePermissionMembersRequest 批量创建邮件组权限成员请求
type BatchCreatePermissionMembersRequest struct {
	Items []CreatePermissionMemberRequest `json:"items"` // 权限成员列表
}

// BatchCreatePermissionMembersResponse 批量创建邮件组权限成员响应
type BatchCreatePermissionMembersResponse struct {
	Items []feishumail.MailGroupMember `json:"items"` // 创建的权限成员列表
}

// BatchCreatePermissionMembers 批量创建邮件组权限成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 批量创建请求
// 返回: 创建的权限成员列表
func BatchCreatePermissionMembers(client *feishumail.Client, mailgroupID string, req BatchCreatePermissionMembersRequest) ([]feishumail.MailGroupMember, error) {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/permission_members/batch_create", mailgroupID)
	var resp BatchCreatePermissionMembersResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return resp.Items, nil
}

// BatchDeletePermissionMembersRequest 批量删除邮件组权限成员请求
type BatchDeletePermissionMembersRequest struct {
	PermissionMemberIDs []string `json:"permission_member_ids"` // 权限成员ID列表
}

// BatchDeletePermissionMembers 批量删除邮件组权限成员
// client: 飞书API客户端
// mailgroupID: 邮件组ID
// req: 批量删除请求
func BatchDeletePermissionMembers(client *feishumail.Client, mailgroupID string, req BatchDeletePermissionMembersRequest) error {
	path := fmt.Sprintf("/mail/v1/mailgroups/%s/permission_members/batch_delete", mailgroupID)
	return client.DoRequest("POST", path, req, nil)
}
