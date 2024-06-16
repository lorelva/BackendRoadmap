package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lorelva/BackendRoadmap/GO/API/internal/domain"
	"github.com/lorelva/BackendRoadmap/GO/API/internal/repository"
)

type MusicController struct {
	repo *repository.MusicRepository
}

func InitMusicController(repo *repository.MusicRepository) *MusicController {
	return &MusicController{
		repo: repo,
	}
}

// ControllerMessageResponse define una estructura estandarizada para las respuestas del controlador
type ControllerMessageResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (m *MusicController) CreateMusic(c echo.Context) error {
	music := domain.Music{}
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body of the music: %v", err),
		})
	}

	if err := m.repo.Create(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert music: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Music added successfully",
		Data:       music,
	})
}

func (m *MusicController) UpdateMusic(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid music ID",
		})
	}

	var music domain.Music
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid request payload: %v", err),
		})
	}

	if err := m.repo.Update(id, &music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to update music: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Music with the ID requested updated successfully",
		Data:       music,
	})
}

func (m *MusicController) DeleteMusic(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid music ID",
		})
	}

	if err := m.repo.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to delete music: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Music with the requested ID was deleted successfully",
	})
}

func (m *MusicController) GetAllMusic(c echo.Context) error {
	idParam := c.QueryParam("id")
	var id int
	var err error
	if idParam != "" {
		id, err = strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid music ID",
			})
		}
	}

	music, err := m.repo.GetAllOrByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to retrieve music: %v", err),
		})
	}
	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved music",
		Data:       music,
	})
}
