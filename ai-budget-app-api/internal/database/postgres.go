package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// SSLモードのデフォルト設定
	if sslmode == "" {
		sslmode = "require"
	}

	// 必須環境変数の確認
	if host == "" || port == "" || user == "" || password == "" || dbname == "" || sslmode == "" {
		return nil, fmt.Errorf("データベース接続情報が環境変数に設定されていません")
	}

	// port番号が数字であることを確認
	portNum, err := strconv.Atoi(port)
	if err != nil || portNum <= 0 || portNum > 65535 {
		return nil, fmt.Errorf("無効なポート番号が指定されています: %s", port)
	}

	// 接続文字列構築（GORMのPostgreSQLドライバ用）
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// GORMでデータベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// 接続エラーの場合
	if err != nil {
		return nil, fmt.Errorf("データベース接続時にエラーが発生しました: %v", err)
	}

	// 接続確認
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("データベースインスタンスの取得に失敗しました: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("データベースへのpingに失敗しました: %v", err)
	}

	return db, nil
}