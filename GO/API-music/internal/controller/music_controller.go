package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/domain"
	"github.com/lorelva/BackendRoadmap/GO/API-music/internal/repository"
)

type MusicController struct {
	repo *repository.MusicRepository
}

func InitMusicController(repo *repository.MusicRepository) *MusicController {
	return &MusicController{
		repo: repo,
	}
}

type ControllerMessageResponse struct {
	StatusCode int         `json:"status_code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (m *MusicController) CreateMusic(c echo.Context) error {
	music := domain.Music{}
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened  trying to bind the body of music %v", err),
		})
	}

	if err := m.repo.Create(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened while trying to insert music %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Music added succesfully",
	})
}

func (m *MusicController) UpdateMusic(c echo.Context) error {
	musicID := c.QueryParam("id")
	musicIDConverted, err := strconv.Atoi(musicID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("ID could not convert %v", err),
		})
	}

	music := domain.Music{}
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened  trying to bind the body of music %v", err),
		})
	}

	if err := m.repo.Update(musicIDConverted, &music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened while trying to update music %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Music updated succesfully",
	})
}

func (m *MusicController) DeleteMusic(c echo.Context) error {
	musicID := c.QueryParam("id")
	musicIDConverted, err := strconv.Atoi(musicID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("ID could not convert %v", err),
		})
	}

	if err := m.repo.Delete(musicIDConverted); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened while trying to delete music %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Music deleted succesfully",
	})
}

func (m *MusicController) GetMusic(c echo.Context) error {
	allMusic, err := m.repo.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to retrieved music information %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Music updated succesfully",
		Data:       allMusic,
	})
}
