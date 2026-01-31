package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// CreateAliasRequest 创建用户邮箱别名请求
type CreateAliasRequest struct {
	Email string `json:"email"` // 别名邮箱地址（必填）
}

// CreateAliasResponse 创建用户邮箱别名响应
type CreateAliasResponse struct {
	Alias feishumail.UserMailboxAlias `json:"alias"`
}

// CreateAlias 创建用户邮箱别名
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID
// req: 创建别名请求
// 返回: 创建的别名信息
func CreateAlias(client *feishumail.Client, userMailboxID string, req CreateAliasRequest) (*feishumail.UserMailboxAlias, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/aliases", userMailboxID)
	var resp CreateAliasResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Alias, nil
}

// DeleteAlias 删除用户邮箱别名
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID
// aliasID: 别名ID
func DeleteAlias(client *feishumail.Client, userMailboxID, aliasID string) error {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/aliases/%s", userMailboxID, aliasID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// ListAliasesResponse 获取用户邮箱所有别名响应
type ListAliasesResponse struct {
	Items []feishumail.UserMailboxAlias `json:"items"` // 别名列表
}

// ListAliases 获取用户邮箱所有别名
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID
// 返回: 别名列表
func ListAliases(client *feishumail.Client, userMailboxID string) ([]feishumail.UserMailboxAlias, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/aliases", userMailboxID)
	var resp ListAliasesResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Items, nil
}
