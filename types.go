package feishumail

// BaseResponse 飞书API通用响应结构
type BaseResponse struct {
	Code int    `json:"code"` // 错误码，0表示成功
	Msg  string `json:"msg"`  // 错误信息
}

// PageInfo 分页信息
type PageInfo struct {
	PageToken string `json:"page_token,omitempty"` // 分页token
	PageSize  int    `json:"page_size,omitempty"`  // 每页数量
	HasMore   bool   `json:"has_more,omitempty"`   // 是否还有更多数据
}

// MailGroup 邮件组信息
type MailGroup struct {
	MailgroupID              string `json:"mailgroup_id"`                // 邮件组ID
	Email                     string `json:"email"`                       // 邮件组邮箱地址
	Name                      string `json:"name"`                        // 邮件组名称
	Description               string `json:"description,omitempty"`         // 邮件组描述
	DirectMembersCount        string `json:"direct_members_count"`         // 直接成员数量
	IncludeExternalMember     bool   `json:"include_external_member"`     // 是否包含外部成员
	IncludeAllCompanyMember   bool   `json:"include_all_company_member"`   // 是否包含所有公司成员
	WhoCanSendMail            string `json:"who_can_send_mail"`          // 谁可以发送邮件：ALL_INTERNAL_USERS, ALL_GROUP_MEMBERS等
}

// MailGroupMember 邮件组成员信息
type MailGroupMember struct {
	MemberID    string `json:"member_id,omitempty"`    // 成员ID
	Email       string `json:"email"`                   // 成员邮箱
	Name        string `json:"name,omitempty"`          // 成员名称
	UserID      string `json:"user_id,omitempty"`       // 用户ID
	Type        string `json:"type,omitempty"`          // 成员类型：USER, DEPARTMENT等
}

// MailGroupAlias 邮件组别名
type MailGroupAlias struct {
	AliasID string `json:"alias_id,omitempty"` // 别名ID
	Email   string `json:"email"`               // 别名邮箱地址
}

// PublicMailbox 公共邮箱信息
type PublicMailbox struct {
	PublicMailboxID string `json:"public_mailbox_id"` // 公共邮箱ID
	Email           string `json:"email"`             // 公共邮箱地址
	Name            string `json:"name"`             // 公共邮箱名称
}

// PublicMailboxMember 公共邮箱成员信息
type PublicMailboxMember struct {
	MemberID string `json:"member_id,omitempty"` // 成员ID
	UserID   string `json:"user_id,omitempty"`   // 用户ID
	Type     string `json:"type,omitempty"`     // 成员类型
}

// UserMailboxAlias 用户邮箱别名
type UserMailboxAlias struct {
	AliasID string `json:"alias_id,omitempty"` // 别名ID
	Email   string `json:"email"`               // 别名邮箱地址
}

// MailAddressStatus 邮箱地址状态
type MailAddressStatus struct {
	Email     string `json:"email"`      // 邮箱地址
	Status    string `json:"status"`      // 状态：ACTIVE, DELETED等
	IsPrimary bool   `json:"is_primary"` // 是否为主邮箱
}
