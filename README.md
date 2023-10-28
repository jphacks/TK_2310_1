# backend

### エンドポイント

prd: https://jphacks-backend-prd-twmyiymqla-an.a.run.app

stg: https://jphacks-backend-stg-twmyiymqla-an.a.run.app

dev1: https://jphacks-backend-dev1-twmyiymqla-an.a.run.app

##  データベースに接続する

```bash
make connect-to-db
```

## デプロイのタイミング

### prd 環境
prod ブランチに push されたとき。

### stg 環境

main ブランチに push がされたとき。

### dev 環境

#### dev1
dev1 ブランチに push がされたとき。

## dev 環境へのデプロイ方法

### dev1 環境

```bash
make deploy-dev1 branch=ブランチ名
```
