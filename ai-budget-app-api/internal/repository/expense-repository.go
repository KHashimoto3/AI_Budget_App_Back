package repository

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpenses(expenses []model.Expense) ([]model.Expense, error)
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

// CreateExpenses: 複数の支出を一括作成する
func (r *expenseRepository) CreateExpenses(expenses []model.Expense) (resultExpenses []model.Expense, err error) {
	// トランザクション内で複数の支出を作成
	err = r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&expenses).Error
	})

	if err != nil {
		return nil, err
	}

	return expenses, nil
}