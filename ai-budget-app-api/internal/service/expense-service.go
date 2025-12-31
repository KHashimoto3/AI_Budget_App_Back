package service

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/repository"
	"github.com/google/uuid"
)

type ExpenseService interface {
	CreateExpenses(expenseRequests []model.CreateExpenseRequest, userID string) ([]model.RegisteredExpense, error)
	GetAllExpenses(userID string, year int, month int) ([]model.RegisteredExpense, error)
}

type expenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

// CreateExpenses: 複数の支出を一括作成する
func (s *expenseService) CreateExpenses(expenseRequests []model.CreateExpenseRequest, userID string) (resultExpenses []model.RegisteredExpense, err error) {
	// リクエストをモデルに変換
	var expenses []model.Expense
	for _, req := range expenseRequests {
		expense := model.Expense {
			ID: uuid.New(),
			UserID: uuid.MustParse(userID),
			ExpenseDate: req.ExpenseDate,
			Amount: req.Amount,
			GenresID: req.GenresID,
			ShopName: req.ShopName,
			Memo: req.Memo,
			InputType: req.InputType,
			ImageID: req.ImageID,
		}
		expenses = append(expenses, expense)
	}

	resultExpenses, err = s.repo.CreateExpenses(expenses)
	if err != nil {
		return nil, err
	}

	return resultExpenses, nil
}

// GetAllExpenses: 指定ユーザー、指定年月の支出一覧を取得する
func (s *expenseService) GetAllExpenses(userID string, year int, month int) (expenses []model.RegisteredExpense, err error) {
	expenses, err = s.repo.GetAllExpenses(uuid.MustParse(userID), year, month)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}