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
DB_SSLMODE=*****
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

### 1. マイグレーションの元になるファイル`schema.sql`の作成

### 2. atlas.hcl の作成

- `atlas.hcl.sample`をもとに、`atlas.hcl`ファイルを作成し、
  `env "dev"`の中身の、`dev`と`url`に指定された値を設定する。
- マイグレーション対象の DB に、`atlas_dev`というデータベースを作成する。

### 3. マイグレーション実施

#### dev 環境

```sh
# 差分を見てマイグレーションファイル作成（initの部分は任意のマイグレーション名）
atlas migrate diff init --env dev
```

```sh
# マイグレーション実行
atlas migrate apply --env dev
```
