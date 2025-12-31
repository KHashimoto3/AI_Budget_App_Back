package repository

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
GetUserIDByFirebaseUID(firebaseUID string) (string, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// GetUserIDByFirebaseUID: FirebaseUIDからユーザーIDを取得する
func (r *userRepository) GetUserIDByFirebaseUID(firebaseUID string) (userID string, err error) {
	var user model.User
	err = r.db.First(&user, "firebase_uid = ?", firebaseUID).Error

	if err != nil {
		return "", err
	}

	return user.ID.String(), nil
}