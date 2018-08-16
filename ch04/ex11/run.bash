#!/usr/bin/env bash

# https://github.com/settings/tokens
export github_token=xxxxxxxxxx

# 一覧
go run issue.go -action=list

# 作成
#go run issue.go -action=create -title="test" -body="this is body"

# コメント表示
#go run issue.go -action=show -number=1

# クローズ
#go run issue.go -action=close -number=1
