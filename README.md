# インストール


## Org,Project名変更
```
export ORG_NAME=yuuuutsk
export PROJECT_NAME=gobase-backend
find . -type f -name '*.go' -exec sed -i '' -e "s/yuuuutsk/${ORG_NAME}/g" {} +;
find . -type f -name 'go.*' -exec sed -i '' -e "s/yuuuutsk/${ORG_NAME}/g" {} +;
find . -type f -name '*.go' -exec sed -i '' -e "s/gobase-backend/${PROJECT_NAME}/g" {} +;
find . -type f -name '*.sql' -exec sed -i '' -e "s/gobase-backend/${PROJECT_NAME}/g" {} +;
find . -type f -name 'Makefile' -exec sed -i '' -e "s/gobase-backend/${PROJECT_NAME}/g" {} +;
find . -type f -name 'go.*' -exec sed -i '' -e "s/gobase-backend/${PROJECT_NAME}/g" {} +;
find . -type f -name 'Dockerfile' -exec sed -i '' -e "s/yuuuutsk/${ORG_NAME}/g" {} +;
find . -type f -name 'Dockerfile' -exec sed -i '' -e "s/gobase-backend/${PROJECT_NAME}/g" {} +;
find . -type f -name 'docker-compose.yml' -exec sed -i '' -e "s/gobase-backend/${PROJECT_NAME}/g" {} +;
find . -type f -name '.github/workflows/*.yml' -exec sed -i '' -e "s/gobase-backend/${PROJECT_NAME}/g" {} +;
find . -type f -name '.github/workflows/*.yml' -exec sed -i '' -e "s/yuuuutsk/${ORG_NAME}/g" {} +;
```

## 使用するツール群のインストール

```
brew install sqldef/sqldef/mysqldef
make install-tools
```

# DBセットアップ
```
$ docker-compose up -d
$ make db-migrate-all
```

# 自動生成
```
$ make gen
$ make wire
```

# コマンド実行
```
$ go run cmd/cli/main.go
```
