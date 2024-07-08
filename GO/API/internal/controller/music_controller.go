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

type ControllerMessageResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// w http.ResponseWriter, r *http.Request
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
	})
}

func (m *MusicController) UpdateMusic(c echo.Context) error {
	//Se obtiene el parámetro de consulta id de la solicitud HTTP.
	//strconv.Atoi(musicID) convierte el ID de música de cadena de texto a entero
	musicID := c.QueryParam("id")
	musicIDConverted, err := strconv.Atoi(musicID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid music ID requested: %v", err),
		})
	}

	var music domain.Music
	if err := c.Bind(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body of the music: %v", err),
		})
	}

	//m.repo.Update(musicIDConverted, &music) intenta actualizar la música en el
	//repositorio con el ID especificado.
	if err := m.repo.Update(musicIDConverted, &music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to update music: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Music with the requested ID updated successfully",
	})
}

func (m *MusicController) DeleteMusic(c echo.Context) error {
	musicID := c.QueryParam("id")
	musicIDConverted, err := strconv.Atoi(musicID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid music ID requested: %v", err),
		})
	}

	if err := m.repo.Delete(musicIDConverted); err != nil {
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
	allMusic, err := m.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to retrieve all music: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved all music",
		Data:       allMusic,
	})
}
