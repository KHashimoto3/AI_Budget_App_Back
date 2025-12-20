# 開発用PostgreSQLデータベース

このディレクトリには、ローカル開発用のPostgreSQLデータベースをDockerで起動するための設定が含まれています。

## 前提条件

- Docker
- Docker Compose

## セットアップ

1. 環境変数ファイルを作成:

```bash
cd dev-db
cp .env.example .env
```

2. 必要に応じて `.env` ファイルを編集してデータベースの設定を変更できます。

## 使用方法

### データベースの起動

```bash
cd dev-db
docker-compose up -d
```

### データベースの停止

```bash
docker-compose down
```

### データベースの停止とデータ削除

```bash
docker-compose down -v
```

### ログの確認

```bash
docker-compose logs -f
```

### データベースへの接続

```bash
docker-compose exec postgres psql -U aibudget -d ai_budget_db
```

または、ホストマシンから接続:

```bash
psql -h localhost -p 5432 -U aibudget -d ai_budget_db
```

## デフォルト設定

- **ホスト**: localhost
- **ポート**: 5432
- **データベース名**: ai_budget_db
- **ユーザー名**: aibudget
- **パスワード**: aibudget_dev_password

## 初期化スクリプト

`init/` ディレクトリに `.sql` ファイルを配置すると、コンテナの初回起動時に自動的に実行されます。
ファイルはアルファベット順に実行されるため、以下のような命名規則を推奨します:

```
init/
  01_create_tables.sql
  02_create_indexes.sql
  03_insert_seed_data.sql
```

## データの永続化

データは Docker ボリューム `postgres_data` に保存されます。
コンテナを削除してもデータは保持されますが、ボリュームも削除する場合は `-v` オプションを使用してください。

## トラブルシューティング

### ポートが既に使用されている場合

`.env` ファイルで `DB_PORT` を変更してください:

```
DB_PORT=5433
```

### データベースをリセットしたい場合

```bash
docker-compose down -v
docker-compose up -d
```

これによりボリュームも削除され、クリーンな状態から再起動します。
