package service

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/repository"
	"github.com/google/uuid"
)

type UserService interface {
	IsExistingUser(firebaseUID string) (bool, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// IsExistingUser: 指定されたFirebaseUIDのユーザーが存在するか確認する
func (s *userService) IsExistingUser(firebaseUID string) (bool, error) {
	return s.repo.IsExistingUser(firebaseUID)
}

// CreateUserByFirebaseUID: FirebaseUIDで新しいユーザーを作成する
func (s *userService) CreateUserByFirebaseUID(user model.RegisteredGoogleLoginUser) (createdUser model.RegisteredGoogleLoginUser, err error) {
	newUser := model.User{
		ID:       uuid.New(),
		FirebaseUID: user.FirebaseUID,
		Name:     user.Name,
		DispName: user.DispName,
		Email:    user.Email,
	}

	createdUser, err = s.repo.CreateUserByFirebaseUID(newUser)
	if err != nil {
		return model.RegisteredGoogleLoginUser{}, err
	}

	return createdUser, nil
}

