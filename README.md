# AI_Budget_App_Back

取引画面を AI で読み取り、簡単に記録できる家計簿 WEB アプリ

バックエンドリポジトリ

## 環境構築

### AWS CLI インストール

```sh
brew install awscli

# Apple Sillicon Macのため、`Cannot install under Rosetta 2 in ARM default prefix (/opt/homebrew)!`が出る場合

arch -arm64 brew install awscli
```

### configure の設定

#### 1. 管理者に IAM ユーザーを作成してもらう

#### 2. 以下のコマンドから設定する

```sh
aws configure --profile #${指定された名前}
```

## 起動方法

### `.env`ファイルを作成

```env
DB_HOST=*****
DB_PORT=*****
DB_USER=*****
DB_PASSWORD=*****
DB_NAME=*****
```

**\***部分は管理者に問い合わせること

### ローカルで起動

#### 起動コマンド

```sh
go run main.go serve
```

### Docker で起動

#### 起動コマンド

```sh
docker-compose build --no-cache
docker-compose up
```

### Open API プレビュー

#### 1. Open API CLI のインストール

```sh
npm install -g @redocly/cli
```

#### 2. yml ファイルのディレクトリに移動し、プレビューサーバを起動

```sh
cd open-api
openapi preview
```

## デプロイ手順

### 1. ビルド

```sh
# 通常のビルド
docker build -t ai-budget-api .

# m1などのarm mac環境からビルドする場合
docker buildx build --platform linux/amd64 -t ai-budget-api .
```

### 2. タグ付け

```sh
docker tag ai-budget-api:latest \
  ${使用するIAMユーザーのID}.dkr.ecr.ap-northeast-1.amazonaws.com/ai-budget-api:latest
```

### 3. push

```sh
docker push ${使用するIAMユーザーのID}.dkr.ecr.ap-northeast-1.amazonaws.com/ai-budget-api:latest
```

### 4. aws の App Runner コンソールからデプロイ操作の続きを行う

## DB マイグレーション

全て、`ai-budget-app-api/migrations`ディレクトリ内で実行する。

### マイグレーションファイル未作成の場合

```sh
atlas schema inspect -u "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=public&sslmode=disable" > schema.hcl
```

### マイグレーションファイル`schema.hcl`の編集

### マイグレーション実施

```sh
atlas schema apply \
  -u "mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" \
  --to file://schema.hcl
```
