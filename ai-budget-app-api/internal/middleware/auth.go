package middleware

import (
	"context"
	"fmt"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

var authClient *auth.Client

// SDKの初期化
func InitializeFirebaseApp() (error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("Firebaseアプリの初期化に失敗しました: %v", err)
	}
	
	authClient, err = app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("Firebase認証クライアントの取得に失敗しました: %v", err)
	}

	return nil
}

// トークン検証ミドルウェア関数
func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// AuthorizationヘッダーからBearerトークンを取得
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(401, model.ErrorResponse{
					Error: "認証に失敗しました",
					Details: "認証に必要な情報が不足しています",
				})
			}

			idToken := strings.TrimPrefix(authHeader, "Bearer ")
			if idToken == authHeader {
				return echo.NewHTTPError(401, model.ErrorResponse{
					Error: "認証に失敗しました",
					Details: "認証に必要な情報が不足しています",
				})
			}

			// トークンの検証
			token, err := authClient.VerifyIDToken(c.Request().Context(), idToken)
			if err != nil {
				return echo.NewHTTPError(401, model.ErrorResponse{
					Error: "認証に失敗しました",
					Details: "有効な認証トークンが必要です",
				})
			}

			// ユーザー情報をコンテキストに保存
			c.Set("userID", token.UID)

			return next(c)
		}
	}
}