package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// ReceiveRule 收信规则信息
type ReceiveRule struct {
	RuleID      string   `json:"rule_id"`      // 规则ID
	Name        string   `json:"name"`         // 规则名称
	Conditions  []Condition `json:"conditions"` // 条件列表
	Actions     []Action    `json:"actions"`    // 动作列表
	IsEnabled   bool    `json:"is_enabled"`    // 是否启用
	Priority    int     `json:"priority"`      // 优先级
}

// Condition 收信规则条件
type Condition struct {
	Field    string `json:"field"`    // 字段：FROM, TO, SUBJECT等
	Operator string `json:"operator"` // 操作符：CONTAINS, EQUALS等
	Value    string `json:"value"`    // 值
}

// Action 收信规则动作
type Action struct {
	Type      string `json:"type"`       // 动作类型：MOVE_TO_FOLDER, MARK_AS_READ等
	Parameter string `json:"parameter"`  // 参数（如文件夹ID）
}

// CreateReceiveRuleRequest 创建收信规则请求
type CreateReceiveRuleRequest struct {
	UserMailboxID string     `json:"-"` // 用户邮箱ID（路径参数）
	Name          string     `json:"name"` // 规则名称（必填）
	Conditions    []Condition `json:"conditions"` // 条件列表（必填）
	Actions       []Action    `json:"actions"`    // 动作列表（必填）
	IsEnabled     bool       `json:"is_enabled,omitempty"` // 是否启用（可选，默认true）
	Priority      int        `json:"priority,omitempty"`  // 优先级（可选）
}

// CreateReceiveRuleResponse 创建收信规则响应
type CreateReceiveRuleResponse struct {
	Rule ReceiveRule `json:"rule"` // 规则信息
}

// CreateReceiveRule 创建收信规则
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 创建请求
// 返回: 创建的规则信息
func CreateReceiveRule(client *feishumail.Client, userMailboxID string, req CreateReceiveRuleRequest) (*ReceiveRule, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/receive_rules", userMailboxID)
	var resp CreateReceiveRuleResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Rule, nil
}

// ListReceiveRulesRequest 列出收信规则请求
type ListReceiveRulesRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	PageToken     string `json:"page_token,omitempty"` // 分页token（查询参数）
	PageSize      int    `json:"page_size,omitempty"`  // 每页数量（查询参数）
}

// ListReceiveRulesResponse 列出收信规则响应
type ListReceiveRulesResponse struct {
	Items    []ReceiveRule      `json:"items"`     // 规则列表
	PageInfo feishumail.PageInfo `json:"page_info"` // 分页信息
}

// ListReceiveRules 列出收信规则
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 查询请求
// 返回: 规则列表和分页信息
func ListReceiveRules(client *feishumail.Client, userMailboxID string, req ListReceiveRulesRequest) (*ListReceiveRulesResponse, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/receive_rules", userMailboxID)
	
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
	
	var resp ListReceiveRulesResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateReceiveRuleRequest 更新收信规则请求
type UpdateReceiveRuleRequest struct {
	Name      string     `json:"name,omitempty"`      // 规则名称（可选）
	Conditions []Condition `json:"conditions,omitempty"` // 条件列表（可选）
	Actions    []Action    `json:"actions,omitempty"`    // 动作列表（可选）
	IsEnabled  *bool      `json:"is_enabled,omitempty"` // 是否启用（可选）
	Priority   int        `json:"priority,omitempty"`   // 优先级（可选）
}

// UpdateReceiveRuleResponse 更新收信规则响应
type UpdateReceiveRuleResponse struct {
	Rule ReceiveRule `json:"rule"` // 规则信息
}

// UpdateReceiveRule 更新收信规则
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// ruleID: 规则ID
// req: 更新请求
// 返回: 更新后的规则信息
func UpdateReceiveRule(client *feishumail.Client, userMailboxID, ruleID string, req UpdateReceiveRuleRequest) (*ReceiveRule, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/receive_rules/%s", userMailboxID, ruleID)
	var resp UpdateReceiveRuleResponse
	if err := client.DoRequest("PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp.Rule, nil
}

// DeleteReceiveRule 删除收信规则
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// ruleID: 规则ID
func DeleteReceiveRule(client *feishumail.Client, userMailboxID, ruleID string) error {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/receive_rules/%s", userMailboxID, ruleID)
	return client.DoRequest("DELETE", path, nil, nil)
}
