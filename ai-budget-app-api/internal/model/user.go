package model

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	FirebaseUID string `json:"firebase_uid"`
	Name string `json:"name"`
	DispName string `json:"disp_name"`
	Email string `json:"email"`
	PasswordHash *string `json:"password_hash,omitempty"`
	AccountType string `json:"account_type"`
}