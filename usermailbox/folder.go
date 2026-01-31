package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// MailFolder 邮箱文件夹信息
type MailFolder struct {
	FolderID   string `json:"folder_id"`   // 文件夹ID
	Name       string `json:"name"`        // 文件夹名称
	ParentID   string `json:"parent_id,omitempty"` // 父文件夹ID
	Type       string `json:"type,omitempty"`      // 文件夹类型：INBOX, SENT, DRAFT等
	MessageCount int  `json:"message_count,omitempty"` // 邮件数量
}

// CreateFolderRequest 创建邮箱文件夹请求
type CreateFolderRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	Name          string `json:"name"` // 文件夹名称（必填）
	ParentID      string `json:"parent_id,omitempty"` // 父文件夹ID（可选）
}

// CreateFolderResponse 创建邮箱文件夹响应
type CreateFolderResponse struct {
	Folder MailFolder `json:"folder"` // 文件夹信息
}

// CreateFolder 创建邮箱文件夹
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 创建请求
// 返回: 创建的文件夹信息
func CreateFolder(client *feishumail.Client, userMailboxID string, req CreateFolderRequest) (*MailFolder, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/folders", userMailboxID)
	var resp CreateFolderResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Folder, nil
}

// DeleteFolder 删除邮箱文件夹
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// folderID: 文件夹ID
func DeleteFolder(client *feishumail.Client, userMailboxID, folderID string) error {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/folders/%s", userMailboxID, folderID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// UpdateFolderRequest 修改邮箱文件夹请求
type UpdateFolderRequest struct {
	Name     string `json:"name,omitempty"`     // 文件夹名称（可选）
	ParentID string `json:"parent_id,omitempty"` // 父文件夹ID（可选）
}

// UpdateFolderResponse 修改邮箱文件夹响应
type UpdateFolderResponse struct {
	Folder MailFolder `json:"folder"` // 文件夹信息
}

// UpdateFolder 修改邮箱文件夹
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// folderID: 文件夹ID
// req: 更新请求
// 返回: 更新后的文件夹信息
func UpdateFolder(client *feishumail.Client, userMailboxID, folderID string, req UpdateFolderRequest) (*MailFolder, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/folders/%s", userMailboxID, folderID)
	var resp UpdateFolderResponse
	if err := client.DoRequest("PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Folder, nil
}

// ListFoldersRequest 列出邮箱文件夹请求
type ListFoldersRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	PageToken     string `json:"page_token,omitempty"` // 分页token（查询参数）
	PageSize      int    `json:"page_size,omitempty"`  // 每页数量（查询参数）
}

// ListFoldersResponse 列出邮箱文件夹响应
type ListFoldersResponse struct {
	Items    []MailFolder        `json:"items"`     // 文件夹列表
	PageInfo feishumail.PageInfo `json:"page_info"` // 分页信息
}

// ListFolders 列出邮箱文件夹
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 查询请求
// 返回: 文件夹列表和分页信息
func ListFolders(client *feishumail.Client, userMailboxID string, req ListFoldersRequest) (*ListFoldersResponse, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/folders", userMailboxID)
	
	// 构建查询参数
	params := make([]string, 0)
	if req.PageToken != "" {
		params = append(params, "page_token="+req.PageToken)
	}
	if req.PageSize > 0 {
		params = append(params, fmt.Sprintf("page_size=%d", req.PageSize))
	}
	
	if len(params) > 0 {
		path += "?" + params[0]
		for i := 1; i < len(params); i++ {
			path += "&" + params[i]
		}
	}
	
	var resp ListFoldersResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
