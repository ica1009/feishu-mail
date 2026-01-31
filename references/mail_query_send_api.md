# 飞书邮件查询和发送API检索结果

## 检索时间
2024年

## API文档位置

### 1. 主文档入口
- **飞书API文档中心**: https://feishu.apifox.cn/doc-1939273
- **飞书开放平台**: https://open.feishu.cn/document

### 2. 邮件消息相关API
- **邮件消息文档**: https://open.feishu.cn/document/server-docs/mail-v1/mail-message
  - 可能包含邮件查询、发送等功能

### 3. 附件相关API
- **获取附件信息**: https://feishu.apifox.cn/api-11267633
  - API路径: `/mail/v1/attachments/{attachment_id}`
  - 功能: 获取邮件附件的详细信息

## 已实现的API（46个）

### 邮件组管理（24个）
- 创建、删除、修改、查询邮件组
- 成员管理、别名管理、权限管理
- 所有API路径: `/mail/v1/mailgroups/*`

### 公共邮箱管理（13个）
- 创建、删除、修改、查询公共邮箱
- 成员管理、别名管理
- 所有API路径: `/mail/v1/public_mailboxes/*`

### 用户邮箱管理（5个）
- 地址管理、别名管理、密码管理
- 所有API路径: `/mail/v1/user_mailboxes/*`

### 其他邮箱相关API（4个）
- 通过邮箱获取用户ID
- 获取客服邮箱等

## 需要进一步确认的API

### 邮件查询相关
1. **获取收件箱邮件列表**
   - 可能路径: `/mail/v1/mail-messages` 或 `/mail/v1/inbox`
   - 状态: 需要访问文档确认

2. **获取邮件内容**
   - 可能路径: `/mail/v1/mail-messages/{message_id}`
   - 状态: 需要访问文档确认

3. **搜索邮件**
   - 可能路径: `/mail/v1/mail-messages/search`
   - 状态: 需要访问文档确认

### 邮件发送相关
1. **发送邮件**
   - 可能路径: `/mail/v1/mail-messages/send` 或 `/mail/v1/send`
   - 状态: 需要访问文档确认

2. **草稿管理**
   - 可能路径: `/mail/v1/drafts/*`
   - 状态: 需要访问文档确认

### 附件相关
1. **下载附件**
   - 可能路径: `/mail/v1/attachments/{attachment_id}/download`
   - 状态: 需要访问文档确认

## 建议的下一步操作

1. 访问邮件消息文档页面，查看完整的API列表
2. 确认邮件查询和发送的具体API路径和参数
3. 补充实现这些API到Golang SDK中

## 参考链接

- [飞书API主文档](https://feishu.apifox.cn/doc-1939273)
- [邮件消息文档](https://open.feishu.cn/document/server-docs/mail-v1/mail-message)
- [获取附件信息API](https://feishu.apifox.cn/api-11267633)
