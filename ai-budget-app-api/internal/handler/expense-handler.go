package handler

import (
	"strconv"

	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/service"
	"github.com/labstack/echo/v4"
)

type ExpenseHandler struct {
	service service.ExpenseService
}

func NewExpenseHandler(service service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

// RegisterExpenses: 複数の支出を一括登録するハンドラー
func (h *ExpenseHandler) RegisterExpenses(c echo.Context) error {
	req := new([]model.CreateExpenseRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(400, model.ErrorResponse{
			Error: "不正なリクエストです",
			Details: err.Error(),
		})
	}

	// バリデーション
	for _, expense := range *req {
		if err := c.Validate(expense); err != nil {
			return c.JSON(400, model.ErrorResponse{
				Error: "バリデーションエラーが発生しました",
				Details: err.Error(),
			})
		}
		// amountが0以下の場合のチェック（追加の安全策）
		if expense.Amount <= 0 {
			return c.JSON(400, model.ErrorResponse{
				Error: "バリデーションエラーが発生しました",
				Details: "amountは0より大きい値を指定してください",
			})
		}
	}

	// ユーザーIDをコンテキストから取得
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(401, model.ErrorResponse{
			Error: "認証に失敗しました",
			Details: "ユーザーIDの取得に失敗しました",
		})
	}

	createdExpenses, err := h.service.CreateExpenses(*req, userID)

	if err != nil {
		return c.JSON(500, model.ErrorResponse{
			Error: "サーバーエラーが発生しました。",
			Details: err.Error(),
		})
	}

	return c.JSON(201, model.CreateExpenseResponse{
		Expenses: createdExpenses,
	});
}

// GetAllExpenses: 指定ユーザー、指定年月の支出一覧を取得する
func (h *ExpenseHandler) GetAllExpenses(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	// バリデーション
	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 1900 || year > 2100 {
		return c.JSON(400, model.ErrorResponse{
			Error: "バリデーションエラーが発生しました",
			Details: "yearのパタメータが無効な値です",
		})
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return c.JSON(400, model.ErrorResponse{
			Error: "バリデーションエラーが発生しました",
			Details: "monthのパタメータが無効な値です",
		})
	}

	// ユーザーIDをコンテキストから取得
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(401, model.ErrorResponse{
			Error: "認証に失敗しました",
			Details: "ユーザーIDの取得に失敗しました",
		})
	}

	expenses, err := h.service.GetAllExpenses(userID, year, month)
	if err != nil {
		return c.JSON(500, model.ErrorResponse{
			Error: "サーバーエラーが発生しました。",
			Details: err.Error(),
		})
	}

	return c.JSON(200, model.GetExpensesResponse{
		Expenses: expenses,
	});
}