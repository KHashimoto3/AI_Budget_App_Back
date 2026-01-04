package repository

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserIDByFirebaseUID(firebaseUID string) (string, error)
	IsExistingUser(firebaseUID string) (bool, error)
	CreateUserByFirebaseUID(user model.User) (createdUser model.RegisteredGoogleLoginUser, err error)
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

// IsExistingUser: 指定されたFirebaseUIDのユーザーが存在するか確認する
func (r *userRepository) IsExistingUser(firebaseUID string) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("firebase_uid = ?", firebaseUID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CreateUserByFirebaseUID: FirebaseUIDで新しいユーザーを作成する
func (r *userRepository) CreateUserByFirebaseUID(user model.User) (createdUser model.RegisteredGoogleLoginUser, err error) {
	err = r.db.Create(&user).Error
	if err != nil {
		return model.RegisteredGoogleLoginUser{}, err
	}

	createdUser = model.RegisteredGoogleLoginUser{
		FirebaseUID: user.FirebaseUID,
		Name: user.Name,
		DispName: user.DispName,
		Email: user.Email,
	}

	return createdUser, nil
}