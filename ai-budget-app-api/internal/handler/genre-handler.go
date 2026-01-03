package handler

import (
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/model"
	"github.com/KHashimoto3/AI_Budget_App_Back/ai-budget-app-api/internal/service"
	"github.com/labstack/echo/v4"
)

type GenreHandler struct {
	service service.GenreService
}

func NewGenreHandler(service service.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

// RegisterGenres: 複数のジャンルを一括登録するハンドラー
func (h *GenreHandler) RegisterGenres(c echo.Context) error {
	type RequestBody struct {
		Genres []model.CreateGenreRequest `json:"genres"`
	}

	req := new(RequestBody)

	if err := c.Bind(req); err != nil {
		return c.JSON(400, model.ErrorResponse{
			Error:   "不正なリクエストです",
			Details: err.Error(),
		})
	}

	// バリデーション
	for _, genre := range req.Genres {
		if err := c.Validate(genre); err != nil {
			return c.JSON(400, model.ErrorResponse{
				Error:   "バリデーションエラーが発生しました",
				Details: err.Error(),
			})
		}
	}

	createdGenres, err := h.service.CreateGenres(req.Genres)

	if err != nil {
		return c.JSON(500, model.ErrorResponse{
			Error:   "サーバーエラーが発生しました。",
			Details: err.Error(),
		})
	}

	return c.JSON(201, model.CreateGenreResponse{
		Genres: createdGenres,
	})
}

// GetAllGenres: 全ジャンルの一覧を取得する
func (h *GenreHandler) GetAllGenres(c echo.Context) error {
	genres, err := h.service.GetAllGenres()
	if err != nil {
		return c.JSON(500, model.ErrorResponse{
			Error:   "サーバーエラーが発生しました。",
			Details: err.Error(),
		})
	}

	return c.JSON(200, model.GetGenresResponse{
		Genres: genres,
	})
}
