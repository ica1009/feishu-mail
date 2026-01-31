package usermailbox

import (
	"fmt"
	"github.com/ica1009/feishu-mail"
)

// EventType 事件类型
const (
	EventTypeReceiveMail = "mail.receive" // 收信通知
)

// SubscribeEventRequest 订阅事件请求
type SubscribeEventRequest struct {
	UserMailboxID string   `json:"-"` // 用户邮箱ID（路径参数）
	EventTypes    []string `json:"event_types"` // 事件类型列表（必填）
	WebhookURL    string   `json:"webhook_url,omitempty"` // Webhook URL（可选）
}

// SubscribeEventResponse 订阅事件响应
type SubscribeEventResponse struct {
	SubscriptionID string `json:"subscription_id"` // 订阅ID
	Status         string `json:"status"`          // 订阅状态
}

// SubscribeEvent 订阅事件
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 订阅请求
// 返回: 订阅信息
func SubscribeEvent(client *feishumail.Client, userMailboxID string, req SubscribeEventRequest) (*SubscribeEventResponse, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/events/subscribe", userMailboxID)
	var resp SubscribeEventResponse
	if err := client.DoRequest("POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSubscriptionStatusRequest 获取订阅状态请求
type GetSubscriptionStatusRequest struct {
	UserMailboxID  string `json:"-"` // 用户邮箱ID（路径参数）
	SubscriptionID string `json:"-"` // 订阅ID（路径参数）
}

// GetSubscriptionStatusResponse 获取订阅状态响应
type GetSubscriptionStatusResponse struct {
	SubscriptionID string   `json:"subscription_id"` // 订阅ID
	EventTypes     []string `json:"event_types"`     // 事件类型列表
	Status         string   `json:"status"`          // 订阅状态
	WebhookURL     string   `json:"webhook_url,omitempty"` // Webhook URL
}

// GetSubscriptionStatus 获取订阅状态
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// subscriptionID: 订阅ID
// 返回: 订阅状态信息
func GetSubscriptionStatus(client *feishumail.Client, userMailboxID, subscriptionID string) (*GetSubscriptionStatusResponse, error) {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/events/subscriptions/%s", userMailboxID, subscriptionID)
	var resp GetSubscriptionStatusResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeEvent 取消订阅
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// subscriptionID: 订阅ID
func UnsubscribeEvent(client *feishumail.Client, userMailboxID, subscriptionID string) error {
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/events/subscriptions/%s", userMailboxID, subscriptionID)
	return client.DoRequest("DELETE", path, nil, nil)
}

// ListSubscriptionsRequest 列出所有订阅请求
type ListSubscriptionsRequest struct {
	UserMailboxID string `json:"-"` // 用户邮箱ID（路径参数）
	PageToken     string `json:"page_token,omitempty"` // 分页token（查询参数）
	PageSize      int    `json:"page_size,omitempty"`  // 每页数量（查询参数）
}

// ListSubscriptionsResponse 列出所有订阅响应
type ListSubscriptionsResponse struct {
	Items    []GetSubscriptionStatusResponse `json:"items"`     // 订阅列表
	PageInfo feishumail.PageInfo            `json:"page_info"` // 分页信息
}

// ListSubscriptions 列出所有订阅
// client: 飞书API客户端
// userMailboxID: 用户邮箱ID（可以是邮箱地址或"me"）
// req: 查询请求
// 返回: 订阅列表和分页信息
func ListSubscriptions(client *feishumail.Client, userMailboxID string, req ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	req.UserMailboxID = userMailboxID
	path := fmt.Sprintf("/mail/v1/user_mailboxes/%s/events/subscriptions", userMailboxID)
	
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
	
	var resp ListSubscriptionsResponse
	if err := client.DoRequest("GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
