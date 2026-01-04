package model

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	FirebaseUID string `json:"firebase_uid"`
	Name string `json:"name"`
	DispName string `json:"disp_name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	AccountType int `json:"account_type"`
}

type CreateGoogleLoginUserRequest struct {
	FirebaseUID string `json:"firebase_uid" validate:"required"`
	Name string `json:"name" validate:"required"`
	DispName string `json:"disp_name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type RegisteredGoogleLoginUser struct {
	FirebaseUID string `json:"firebase_uid"`
	Name string `json:"name"`
	DispName string `json:"disp_name"`
	Email string `json:"email"`
}
type CreateGoogleLoginUserResponse struct {
	User RegisteredGoogleLoginUser `json:"user"`
}