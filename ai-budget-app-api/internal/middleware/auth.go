package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

var authClient *auth.Client

type AuthMiddleware struct {
	userRepo repository.UserRepository
}

func NewAuthMiddleware(userRepo repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{userRepo: userRepo}
}

// SDKの初期化
func (m *AuthMiddleware) InitializeFirebaseApp() (error) {
	ctx := context.Background()

	credJSON := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	if credJSON == "" {
		return fmt.Errorf("Firebase credentials not found in env")
	}

	opt := option.WithCredentialsJSON([]byte(credJSON))

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
func (m *AuthMiddleware) FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ★ preflight は認証をスキップ
			if c.Request().Method == echo.OPTIONS {
				return c.NoContent(204)
			}

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

			// ログインしたFirebaseUIDが存在するかを確認
			exists, err := m.userRepo.IsExistingUser(token.UID)
			if err != nil {
				return echo.NewHTTPError(500, model.ErrorResponse{
					Error: "サーバーエラーが発生しました",
					Details: err.Error(),
				})
			}
			// DBに、今のログインユーザーが存在しなければ、名前とメールアドレスを取得して新規ユーザーを作成
			if !exists {
				// Firebaseからユーザー情報を取得
				email, _ := token.Claims["email"].(string)
    			name, _ := token.Claims["name"].(string)

				newUser := model.User{
					ID: uuid.New(),
					FirebaseUID: token.UID,
					Name:        name,
					DispName:    name,
					Email:       email,
					PasswordHash: "none",
					AccountType: 1,
				}

				_, err = m.userRepo.CreateUserByFirebaseUID(newUser)
				if err != nil {
					return echo.NewHTTPError(500, model.ErrorResponse{
						Error: "サーバーエラーが発生しました",
						Details: err.Error(),
					})
				}
			}

			// FirebaseUIDを使ってユーザーIDを取得
			userID, err := m.userRepo.GetUserIDByFirebaseUID(token.UID)
			if err != nil {
				return echo.NewHTTPError(401, model.ErrorResponse{
					Error: "認証に失敗しました",
					Details: "ユーザー情報の取得に失敗しました",
				})
			}

			// ユーザーIDをコンテキストに保存
			c.Set("userID", userID)

			return next(c)
		}
	}
}