package repository

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpenses(expenses []model.Expense) ([]model.RegisteredExpense, error)
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
	
	// トランザクション内で複数の支出を作成
	/*err = r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&expenses).Error
	})

	if err != nil {
		return nil, err
	}

	return expenses, nil*/
}