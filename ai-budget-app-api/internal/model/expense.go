package model

import (
	"github.com/google/uuid"
)

// 支出
type Expense struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	ExpenseDate Date `json:"expense_date" validate:"required"`
	Amount int64 `json:"amount"`
	GenresID uuid.UUID `json:"genres_id"`
	ShopName string `json:"shop_name"`
	Memo string `json:"memo"`
	InputType string `json:"input_type"`
	ImageID *uuid.UUID `json:"image_id,omitempty"`
}

type CreateExpenseRequest struct {
	ExpenseDate Date `json:"expense_date" validate:"required"`
	Amount int64 `json:"amount" validate:"required"`
	GenresID uuid.UUID `json:"genres_id" validate:"required"`
	ShopName string `json:"shop_name"`
	Memo string `json:"memo"`
	InputType string `json:"input_type" validate:"required"`
	ImageID *uuid.UUID `json:"image_id,omitempty"`
}

type CreateExpenseResponse struct {
	Expenses []Expense `json:"expenses"`
}