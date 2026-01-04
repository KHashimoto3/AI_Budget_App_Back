package repository

import (
	"time"

	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpenses(expenses []model.Expense) ([]model.RegisteredExpense, error)
	GetAllExpenses(userID uuid.UUID, year int, month int) ([]model.RegisteredExpense, error)
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

// CreateExpenses: 複数の支出を一括作成する
func (r *expenseRepository) CreateExpenses(expenses []model.Expense) (resultExpenses []model.RegisteredExpense, err error) {
	err = r.db.Create(&expenses).Error
	if err != nil {
		return nil, err
	}

	var createdExpenses []model.RegisteredExpense
	for _, expense := range expenses {
		createdExpense := model.RegisteredExpense {
			ID: expense.ID,
			ExpenseDate: expense.ExpenseDate,
			Amount: expense.Amount,
			GenresID: expense.GenresID,
			ShopName: expense.ShopName,
			Memo: expense.Memo,
			InputType: expense.InputType,
			ImageID: expense.ImageID,
		}
		createdExpenses = append(createdExpenses, createdExpense)
	}

	return createdExpenses, nil
}

// GetAllExpenses: 指定ユーザー、指定年月の支出一覧を取得する
func (r *expenseRepository) GetAllExpenses(userID uuid.UUID, year int, month int) (expenses []model.RegisteredExpense, err error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)

	err = r.db.Model(&model.Expense{}).
		Where("user_id = ? AND expense_date >= ? AND expense_date < ?", userID, startDate, endDate).
		Find(&expenses).Error

	if err != nil {
		return nil, err
	}

	var resultExpenses = make([]model.RegisteredExpense, 0)
	for _, expense := range expenses {
		registeredExpense := model.RegisteredExpense {
			ID: expense.ID,
			ExpenseDate: expense.ExpenseDate,
			Amount: expense.Amount,
			GenresID: expense.GenresID,
			ShopName: expense.ShopName,
			Memo: expense.Memo,
			InputType: expense.InputType,
			ImageID: expense.ImageID,
		}
		resultExpenses = append(resultExpenses, registeredExpense)
	}

	return resultExpenses, nil
}