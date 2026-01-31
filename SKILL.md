---
name: feishu-mail
description: Comprehensive guide for working with Feishu (Lark) Mail APIs. Use when Claude needs to interact with Feishu mail functionality including: (1) Mail group management (creating, deleting, modifying mail groups), (2) Public mailbox operations (managing public mailboxes and members), (3) User mailbox operations (managing user mailboxes, aliases, and addresses), (4) Querying mail-related information, or (5) Any task involving Feishu mail API integration or development.
---

# Feishu Mail API Skill

## Overview

This skill provides comprehensive guidance for working with Feishu (Lark) Mail APIs. It includes complete API documentation, usage patterns, and reference materials for all mail-related operations in the Feishu platform.

## Quick Start

When working with Feishu Mail APIs:

1. **Identify the operation type**: Mail Group, Public Mailbox, or User Mailbox
2. **Check API reference**: See [feishu_mail_api_list.md](references/feishu_mail_api_list.md) for complete API list
3. **Review authentication**: Ensure proper access tokens are configured
4. **Follow API patterns**: Use the documented request/response formats

## API Categories

Feishu Mail APIs are organized into three main categories:

### 1. Mail Group (邮件组) APIs
- Mail group management (create, delete, modify, query)
- Member management (add, remove, query, batch operations)
- Alias management
- Permission member management

### 2. Public Mailbox (公共邮箱) APIs
- Public mailbox management (create, delete, modify, query)
- Member management (add, remove, query, batch operations)
- Alias management

### 3. User Mailbox (用户邮箱) APIs
- User mailbox address management
- Alias management
- Address status queries
- Password reset operations

## Reference Documentation

For complete API details, including:
- All 46 mail-related API endpoints
- API paths and URLs
- Function descriptions
- Usage guidelines

See: [feishu_mail_api_list.md](references/feishu_mail_api_list.md)

## Common Workflows

### Creating a Mail Group
1. Use the "创建邮件组" API
2. Configure group settings
3. Add members using member management APIs
4. Set up aliases if needed

### Managing Public Mailbox
1. Create or query public mailbox
2. Add/remove members
3. Configure aliases
4. Manage permissions

### User Mailbox Operations
1. Query mailbox status
2. Manage aliases
3. Handle address deletion (from recycle bin)
4. Reset passwords when needed

## Resources

### references/
- **feishu_mail_api_list.md**: Complete list of all 46 Feishu Mail APIs with paths, descriptions, and categorization

## Notes

- All APIs require proper authentication tokens (tenant_access_token, app_access_token, etc.)
- API permissions must be applied for in the Feishu Open Platform
- Use Feishu's official SDKs (Java, Golang, Python, NodeJS) when possible
- Refer to official documentation for detailed request/response formats and error codes
