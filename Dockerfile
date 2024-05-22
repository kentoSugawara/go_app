# 使用するGoのイメージ
FROM golang:1.19.1-alpine

# 必要なパッケージのインストール
RUN apk update && apk add git

# アプリケーションの作業ディレクトリを設定
WORKDIR /app

# プロジェクトの全ファイルをコンテナにコピー
COPY . /app

# Goモジュールの依存関係を確認し、必要なパッケージをダウンロード
RUN go mod tidy

# アプリケーションのビルド
RUN go build -o /app/myapp ./cmd/main.go && ls -la /app

# コンテナが起動するときに実行されるコマンド
CMD ["/app/myapp"]

# コンテナのポート8080を開放
EXPOSE 8080
