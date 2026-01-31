package publicmailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// CreateAliasRequest 创建公共邮箱别名请求
type CreateAliasRequest struct {
	Email string `json:"email"` // 别名邮箱地址（必填）
}

// CreateAliasResponse 创建公共邮箱别名响应
type CreateAliasResponse struct {
	Alias feishumail.MailGroupAlias `json:"alias"`
}

// CreateAlias 创建公共邮箱别名
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// req: 创建别名请求
// 返回: 创建的别名信息
func CreateAlias(client *feishumail.Client, publicMailboxID string, req CreateAliasRequest) (*feishumail.MailGroupAlias, error) {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/aliases", publicMailboxID)
	var resp CreateAliasResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Alias, nil
}

// DeleteAlias 删除公共邮箱别名
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// aliasID: 别名ID
func DeleteAlias(client *feishumail.Client, publicMailboxID, aliasID string) error {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/aliases/%s", publicMailboxID, aliasID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// ListAliasesResponse 查询公共邮箱的所有别名响应
type ListAliasesResponse struct {
	Items []feishumail.MailGroupAlias `json:"items"` // 别名列表
}

// ListAliases 查询公共邮箱的所有别名
// client: 飞书API客户端
// publicMailboxID: 公共邮箱ID
// 返回: 别名列表
func ListAliases(client *feishumail.Client, publicMailboxID string) ([]feishumail.MailGroupAlias, error) {
	path := fmt.Sprintf("/mail/v1/public_mailboxes/%s/aliases", publicMailboxID)
	var resp ListAliasesResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Items, nil
}
