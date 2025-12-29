package service

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/repository"
	"github.com/google/uuid"
)

type ExpenseService interface {
	CreateExpenses(expenseRequests []model.CreateExpenseRequest, userID string) ([]model.Expense, error)
}

type expenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

// CreateExpenses: 複数の支出を一括作成する
func (s *expenseService) CreateExpenses(expenseRequests []model.CreateExpenseRequest, userID string) (resultExpenses []model.Expense, err error) {
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