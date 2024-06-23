package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lorelva/BackendRoadmap/GO/API/internal/domain"
	"github.com/lorelva/BackendRoadmap/GO/API/internal/repository"
)

// se utilizan para agrupar datos relacionados y métodos asociados.
// almacenamiento de datos relacionados con la música.
// el controlador puede acceder y modificar
type MusicController struct {
	repo *repository.MusicRepository
}

// para inicializar una nueva instancia de MusicController
// repo permite que el controlador tenga acceso a la lógica de acceso a datos
func InitMusicController(repo *repository.MusicRepository) *MusicController {
	// una sintaxis de literales de estructura para crear una nueva instancia
	//de MusicController y asignar el campo repo.
	return &MusicController{
		repo: repo,
	}
}

// ControllerMessageResponse define una estructura para las respuestas del controlador
// Incluye código de estado (StatusCode), mensaje (Message) y datos (Data).
// permite que cualquier tipo de dato sea incluido en la respuesta,
// En Go  interface{} es como una caja que puede contener cualquier cosa. No importa si es un número,
// una cadena de texto, una lista, etc. Puedes poner cualquier tipo de valor dentro de esta caja.
// type MusicRepo interface{
// Es como un contrato que define lo que una cosa debe poder hacer y hacer una recopilación.
// }
type ControllerMessageResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Toma un parámetro c de tipo echo.Context, que representa el contexto de la solicitud HTTP.
// w http.ResponseWriter, r *http.Request
func (m *MusicController) CreateMusic(c echo.Context) error {
	//Se crea una nueva instancia de domain.Music
	//Declaración de la estructura de música
	music := domain.Music{}
	//// Unmarshal: Convertir el cuerpo de la solicitud JSON a la estructura de Go
	//if err := json.Unmarshal(body, &music)
	//c.Bind(&music) intenta llenar la estructura music con los datos del cuerpo de la solicitud HTTP
	if err := c.Bind(&music); err != nil {
		//Si ocurre un error durante la vinculación, se devuelve una respuesta JSON con un mensaje de error
		//y un código de estado 500 Internal Server Error.
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body of the music: %v", err),
		})
	}

	//se devuelve una respuesta JSON con un mensaje de error y un código de estado 500 Internal Server Error.
	//m.repo.Create(&music) intenta insertar la nueva música en el repositorio.
	if err := m.repo.Create(&music); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert music: %v", err),
		})
	}

	//cómo puedes manejar manualmente el proceso de marshall
	//y unmarshall en Go sin utilizar un framework como Echo.
	//responseBody, _ := json.Marshal(response)
	//Si la creación es exitosa, se devuelve una respuesta JSON
	//con un mensaje de éxito y un código de estado 201 Created
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
