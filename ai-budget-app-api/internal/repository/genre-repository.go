package repository

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"gorm.io/gorm"
)

type GenreRepository interface {
	CreateGenres(genres []model.Genre) ([]model.RegisteredGenre, error)
	GetAllGenres() ([]model.RegisteredGenre, error)
}

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{db: db}
}

// CreateGenres: 複数のジャンルを一括作成する
func (r *genreRepository) CreateGenres(genres []model.Genre) (resultGenres []model.RegisteredGenre, err error) {
	err = r.db.Create(&genres).Error
	if err != nil {
		return nil, err
	}

	var createdGenres []model.RegisteredGenre
	for _, genre := range genres {
		createdGenre := model.RegisteredGenre{
			ID:   genre.ID,
			Name: genre.Name,
		}
		createdGenres = append(createdGenres, createdGenre)
	}

	return createdGenres, nil
}

// GetAllGenres: 全ジャンルの一覧を取得する
func (r *genreRepository) GetAllGenres() (genres []model.RegisteredGenre, err error) {
	var dbGenres []model.Genre
	err = r.db.Model(&model.Genre{}).
		Find(&dbGenres).Error

	if err != nil {
		return nil, err
	}

	var resultGenres = make([]model.RegisteredGenre, 0)
	for _, genre := range dbGenres {
		registeredGenre := model.RegisteredGenre{
			ID:   genre.ID,
			Name: genre.Name,
		}
		resultGenres = append(resultGenres, registeredGenre)
	}

	return resultGenres, nil
}
