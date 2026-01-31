package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ica1009/feishu-mail"
	"github.com/ica1009/feishu-mail/mailgroup"
	"github.com/ica1009/feishu-mail/publicmailbox"
	"github.com/ica1009/feishu-mail/usermailbox"
)

func main() {
	// 从环境变量获取访问令牌
	token := os.Getenv("FEISHU_ACCESS_TOKEN")
	if token == "" {
		log.Fatal("请设置环境变量 FEISHU_ACCESS_TOKEN")
	}

	// 创建客户端
	client := feishumail.NewClient(token)

	// 示例1: 创建邮件组
	fmt.Println("=== 创建邮件组 ===")
	createReq := mailgroup.CreateRequest{
		Email:          "test-group@example.com",
		Name:           "测试邮件组",
		Description:    "这是一个测试邮件组",
		WhoCanSendMail: "ALL_INTERNAL_USERS",
	}
	mailGroup, err := mailgroup.Create(client, createReq)
	if err != nil {
		log.Printf("创建邮件组失败: %v\n", err)
	} else {
		fmt.Printf("邮件组创建成功，ID: %s\n", mailGroup.MailgroupID)
		fmt.Printf("邮箱地址: %s\n", mailGroup.Email)
		fmt.Printf("名称: %s\n", mailGroup.Name)
	}

	// 示例2: 查询邮件组
	if mailGroup != nil {
		fmt.Println("\n=== 查询邮件组 ===")
		mg, err := mailgroup.Get(client, mailGroup.MailgroupID)
		if err != nil {
			log.Printf("查询邮件组失败: %v\n", err)
		} else {
			fmt.Printf("邮件组名称: %s\n", mg.Name)
			fmt.Printf("描述: %s\n", mg.Description)
		}
	}

	// 示例3: 创建公共邮箱
	fmt.Println("\n=== 创建公共邮箱 ===")
	publicReq := publicmailbox.CreateRequest{
		Email:       "public@example.com",
		Name:        "公共邮箱",
		Description: "公司公共邮箱",
	}
	publicMailbox, err := publicmailbox.Create(client, publicReq)
	if err != nil {
		log.Printf("创建公共邮箱失败: %v\n", err)
	} else {
		fmt.Printf("公共邮箱创建成功，ID: %s\n", publicMailbox.PublicMailboxID)
	}

	// 示例4: 创建用户邮箱别名
	fmt.Println("\n=== 创建用户邮箱别名 ===")
	aliasReq := usermailbox.CreateAliasRequest{
		Email: "alias@example.com",
	}
	// 注意：需要替换为实际的用户邮箱ID
	userMailboxID := "user_mailbox_id_here"
	alias, err := usermailbox.CreateAlias(client, userMailboxID, aliasReq)
	if err != nil {
		log.Printf("创建用户邮箱别名失败: %v\n", err)
	} else {
		fmt.Printf("别名创建成功，ID: %s\n", alias.AliasID)
		fmt.Printf("别名邮箱: %s\n", alias.Email)
	}

	fmt.Println("\n示例执行完成！")
}
