# Feishu Mail API Go SDK

飞书邮箱API的Golang客户端库，提供完整的邮件组、公共邮箱和用户邮箱管理功能。

## 功能特性

- ✅ **邮件组管理**：创建、删除、修改、查询邮件组
- ✅ **邮件组成员管理**：添加、删除、查询、批量操作成员
- ✅ **邮件组别名管理**：创建、删除、查询别名
- ✅ **邮件组权限管理**：权限成员的完整管理
- ✅ **公共邮箱管理**：公共邮箱的完整生命周期管理
- ✅ **用户邮箱管理**：用户邮箱地址、别名、密码管理
- ✅ **邮件查询**：获取邮件列表、查询邮件详情
- ✅ **附件管理**：获取附件信息、获取附件下载链接
- ✅ **类型安全**：完整的Go类型定义
- ✅ **错误处理**：统一的错误处理机制

## 安装

```bash
go get github.com/ica1009/feishu-mail
```

## 快速开始

### 1. 创建客户端

```go
package main

import (
    "github.com/ica1009/feishu-mail"
    "github.com/ica1009/feishu-mail/mailgroup"
)

func main() {
    // 使用访问令牌创建客户端
    token := "your_access_token"
    client := feishumail.NewClient(token)
    
    // 或者使用自定义选项
    client = feishumail.NewClientWithOptions(
        token,
        "https://open.feishu.cn/open-apis", // 基础URL
        30*time.Second,                     // 超时时间
    )
}
```

### 2. 邮件组操作示例

```go
// 创建邮件组
req := mailgroup.CreateRequest{
    Email:          "test@example.com",
    Name:           "测试邮件组",
    Description:    "这是一个测试邮件组",
    WhoCanSendMail: "ALL_INTERNAL_USERS",
}
mailGroup, err := mailgroup.Create(client, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("创建的邮件组ID: %s\n", mailGroup.MailgroupID)

// 查询邮件组
mg, err := mailgroup.Get(client, mailGroup.MailgroupID)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("邮件组名称: %s\n", mg.Name)

// 添加成员
memberReq := mailgroup.CreateMemberRequest{
    UserID: "user_id_123",
    Type:   "USER",
}
member, err := mailgroup.CreateMember(client, mailGroup.MailgroupID, memberReq)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("添加的成员ID: %s\n", member.MemberID)

// 删除邮件组
err = mailgroup.Delete(client, mailGroup.MailgroupID)
if err != nil {
    log.Fatal(err)
}
```

### 3. 公共邮箱操作示例

```go
import "github.com/ica1009/feishu-mail/publicmailbox"

// 创建公共邮箱
req := publicmailbox.CreateRequest{
    Email:       "public@example.com",
    Name:        "公共邮箱",
    Description: "公司公共邮箱",
}
publicMailbox, err := publicmailbox.Create(client, req)
if err != nil {
    log.Fatal(err)
}

// 添加成员
memberReq := publicmailbox.AddMemberRequest{
    UserID: "user_id_123",
    Type:   "USER",
}
member, err := publicmailbox.AddMember(client, publicMailbox.PublicMailboxID, memberReq)
if err != nil {
    log.Fatal(err)
}
```

### 4. 用户邮箱操作示例

```go
import "github.com/ica1009/feishu-mail/usermailbox"

// 创建用户邮箱别名
req := usermailbox.CreateAliasRequest{
    Email: "alias@example.com",
}
alias, err := usermailbox.CreateAlias(client, "user_mailbox_id", req)
if err != nil {
    log.Fatal(err)
}

// 重置密码
resetReq := usermailbox.ResetPasswordRequest{
    UserID: "user_id_123",
}
resp, err := usermailbox.ResetPassword(client, resetReq)
if err != nil {
    log.Fatal(err)
}
```

### 5. 邮件查询和附件操作示例

```go
import "github.com/ica1009/feishu-mail/usermailbox"

// 获取邮件列表（收件箱）
listReq := usermailbox.ListMessagesRequest{
    FolderID: "inbox",  // 收件箱
    PageSize: 20,
}
messages, err := usermailbox.ListMessages(client, "me", listReq)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("找到 %d 封邮件\n", len(messages.Items))

// 获取邮件详情
message, err := usermailbox.GetMessage(client, "me", "message_id_123", true)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("邮件主题: %s\n", message.Subject)
fmt.Printf("发件人: %s\n", message.From)

// 获取附件下载链接
attachmentIDs := []string{"attachment_id_1", "attachment_id_2"}
downloadURLs, err := usermailbox.GetAttachmentDownloadURL(
    client, 
    "me", 
    "message_id_123", 
    attachmentIDs,
)
if err != nil {
    log.Fatal(err)
}
for _, item := range downloadURLs {
    fmt.Printf("附件 %s 下载链接: %s\n", item.AttachmentID, item.DownloadURL)
    // 注意：下载链接仅可使用两次，有效期两小时
}

// 获取附件信息
attachment, err := usermailbox.GetAttachmentInfo(
    client,
    "me",
    "message_id_123",
    "attachment_id_1",
)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("附件名称: %s, 大小: %d 字节\n", attachment.Name, attachment.Size)
```

## API模块

### 邮件组 (mailgroup)

- `group.go` - 邮件组管理（创建、删除、修改、查询）
- `member.go` - 成员管理（添加、删除、查询、批量操作）
- `alias.go` - 别名管理
- `permission.go` - 权限成员管理

### 公共邮箱 (publicmailbox)

- `mailbox.go` - 公共邮箱管理
- `member.go` - 成员管理
- `alias.go` - 别名管理

### 用户邮箱 (usermailbox)

- `address.go` - 地址管理
- `alias.go` - 别名管理
- `password.go` - 密码管理
- `message.go` - 邮件查询（获取邮件列表、邮件详情）
- `attachment.go` - 附件管理（获取附件信息、下载链接）

## 错误处理

所有API调用都会返回错误，使用统一的错误类型：

```go
result, err := mailgroup.Create(client, req)
if err != nil {
    if feishuErr, ok := err.(*feishumail.Error); ok {
        fmt.Printf("错误码: %d, 错误信息: %s\n", feishuErr.Code, feishuErr.Message)
    } else {
        fmt.Printf("其他错误: %v\n", err)
    }
    return
}
```

## 认证

使用前需要获取访问令牌：

1. **tenant_access_token** - 租户访问令牌
2. **app_access_token** - 应用访问令牌
3. **user_access_token** - 用户访问令牌

获取方式请参考[飞书开放平台文档](https://open.feishu.cn/document)。

## 完整API列表

本SDK实现了飞书邮箱相关的46个API接口：

- 邮件组相关API：24个
- 公共邮箱相关API：13个
- 用户邮箱相关API：5个
- 其他邮箱相关API：4个

详细API列表请参考 [references/feishu_mail_api_list.md](references/feishu_mail_api_list.md)

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！

## 参考文档

- [飞书开放平台](https://open.feishu.cn/document)
- [飞书API文档](https://feishu.apifox.cn/doc-1939273)
