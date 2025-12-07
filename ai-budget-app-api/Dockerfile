# マルチステージビルド: ビルドステージ
FROM golang:1.24.5-alpine AS builder

# 必要なビルドツールをインストール
RUN apk add --no-cache git

# ワーキングディレクトリを設定
WORKDIR /app

# Go modulesのファイルをコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# バイナリをビルド（静的リンク）
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 実行ステージ
FROM alpine:latest

# セキュリティアップデートと証明書をインストール
RUN apk --no-cache add ca-certificates tzdata

# タイムゾーンを設定（必要に応じて変更）
ENV TZ=Asia/Tokyo

# 非rootユーザーを作成
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# ワーキングディレクトリを設定
WORKDIR /home/appuser

# ビルドステージからバイナリをコピー
COPY --from=builder /app/main .

# 所有権を変更
RUN chown -R appuser:appuser /home/appuser

# 非rootユーザーに切り替え
USER appuser

# ポートを公開
EXPOSE 8080

# ヘルスチェック
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# アプリケーションを起動
CMD ["./main", "serve"]
