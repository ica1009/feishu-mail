#!/bin/bash
# 推送飞书Mail Skill到GitHub的脚本
# 使用方法: ./push_to_github.sh <your-github-username> <repository-name>

if [ $# -lt 2 ]; then
    echo "使用方法: $0 <github-username> <repository-name>"
    echo "示例: $0 yourusername feishu-mail-skill"
    exit 1
fi

GITHUB_USER=$1
REPO_NAME=$2

echo "添加远程仓库..."
git remote add origin git@github.com:${GITHUB_USER}/${REPO_NAME}.git 2>/dev/null || git remote set-url origin git@github.com:${GITHUB_USER}/${REPO_NAME}.git

echo "推送到GitHub..."
git push -u origin main

if [ $? -eq 0 ]; then
    echo "✅ 成功推送到GitHub!"
    echo "仓库地址: https://github.com/${GITHUB_USER}/${REPO_NAME}"
else
    echo "❌ 推送失败，请检查："
    echo "1. GitHub仓库是否已创建"
    echo "2. SSH密钥是否已添加到GitHub"
    echo "3. 仓库名称是否正确"
fi
