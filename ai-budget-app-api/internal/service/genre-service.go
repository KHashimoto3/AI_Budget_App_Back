package service

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/repository"
	"github.com/google/uuid"
)

type GenreService interface {
	CreateGenres(genreRequests []model.CreateGenreRequest) ([]model.RegisteredGenre, error)
	GetAllGenres() ([]model.RegisteredGenre, error)
}

type genreService struct {
	repo repository.GenreRepository
}

func NewGenreService(repo repository.GenreRepository) GenreService {
	return &genreService{repo: repo}
}

// CreateGenres: 複数のジャンルを一括作成する
func (s *genreService) CreateGenres(genreRequests []model.CreateGenreRequest) (resultGenres []model.RegisteredGenre, err error) {
	// リクエストをモデルに変換
	var genres []model.Genre
	for _, req := range genreRequests {
		genre := model.Genre{
			ID:   uuid.New(),
			Name: req.Name,
		}
		genres = append(genres, genre)
	}

	resultGenres, err = s.repo.CreateGenres(genres)
	if err != nil {
		return nil, err
	}

	return resultGenres, nil
}

// GetAllGenres: 全ジャンルの一覧を取得する
func (s *genreService) GetAllGenres() (genres []model.RegisteredGenre, err error) {
	genres, err = s.repo.GetAllGenres()
	if err != nil {
		return nil, err
	}

	return genres, nil
}
