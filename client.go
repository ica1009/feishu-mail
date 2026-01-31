package feishumail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL 飞书API基础URL
	DefaultBaseURL = "https://open.feishu.cn/open-apis"
	// DefaultTimeout 默认请求超时时间
	DefaultTimeout = 30 * time.Second
)

// Client 飞书邮件API客户端
type Client struct {
	baseURL    string        // API基础URL
	token      string        // 访问令牌
	httpClient *http.Client  // HTTP客户端
}

// NewClient 创建新的客户端
// token: 访问令牌（tenant_access_token或app_access_token）
func NewClient(token string) *Client {
	return &Client{
		baseURL: DefaultBaseURL,
		token:   token,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

// NewClientWithOptions 使用自定义选项创建客户端
// token: 访问令牌
// baseURL: 自定义API基础URL（可选，用于企业版）
// timeout: 自定义超时时间
func NewClientWithOptions(token, baseURL string, timeout time.Duration) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	if timeout == 0 {
		timeout = DefaultTimeout
	}
	return &Client{
		baseURL: baseURL,
		token:   token,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// SetToken 设置访问令牌
func (c *Client) SetToken(token string) {
	c.token = token
}

// DoRequest 执行HTTP请求（公开方法供其他包使用）
// method: HTTP方法（GET, POST, PUT, PATCH, DELETE）
// path: API路径（相对于baseURL）
// body: 请求体（nil表示无请求体）
// result: 响应结果（需要是指针类型）
func (c *Client) DoRequest(method, path string, body interface{}, result interface{}) error {
	url := c.baseURL + path

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// 检查HTTP状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("http error: status=%d, body=%s", resp.StatusCode, string(respBody))
	}

	// 解析响应
	var baseResp BaseResponse
	if err := json.Unmarshal(respBody, &baseResp); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}

	// 检查业务错误码
	if !IsSuccess(baseResp.Code) {
		return NewError(baseResp.Code, baseResp.Msg)
	}

	// 如果指定了result，解析data字段
	if result != nil {
		var respData struct {
			Data json.RawMessage `json:"data"`
		}
		if err := json.Unmarshal(respBody, &respData); err != nil {
			return fmt.Errorf("unmarshal data: %w", err)
		}

		if len(respData.Data) > 0 {
			if err := json.Unmarshal(respData.Data, result); err != nil {
				return fmt.Errorf("unmarshal result: %w", err)
			}
		}
	}

	return nil
}
