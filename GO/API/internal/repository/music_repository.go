package repository

import (
	"errors"

	"github.com/lorelva/BackendRoadmap/GO/API/internal/domain"
)

// implementación en memoria del repositorio de música
// estructura que representa nuestro repositorio de música. Tiene:
// Music: Una lista (slice) de canciones.
// NewID: Para dar a cada canción un ID
type MusicRepository struct {
	Music []domain.Music
	NewID int
}

// Un puntero es una variable que almacena la dirección de memoria de otra variable.
func MusicRepositoryService() *MusicRepository {
	//Esta función crea y devuelve un nuevo MusicRepository.
	//Piensa en ello como la creación de una nueva lista de reproducción vacía.
	//Comienza con la lista vacía y el siguiente ID disponible es 1.
	return &MusicRepository{
		Music: []domain.Music{},
		NewID: 1,
	}
}

//Una función es un bloque de código que realiza una tarea específica.
//No está asociada con ningún tipo particular-->Llamarlas independiente
//->Las funciones trabajan con los datos que les pasas como argumentos y de retorno

// Un método es una función que está asociada con un tipo específico
// requieren una instancia de ese tipo para ser llamados
// tienen acceso directo a los datos del tipo al que están asociados (como los campos de una estructura)
func (m *MusicRepository) Create(music *domain.Music) error {
	//nil se refiere a la ausencia de valor o dirección,
	//mientras que "vacío" se refiere a la ausencia de elementos dentro de una estructura
	if music == nil {
		//el método devuelve un error indicando que music no puede ser nil
		return errors.New("music cannot be nil")
	}

	//Se crea una copia de la música recibida (*music) y se asigna a una nueva variable addMusic.
	addMusic := *music
	//Luego, se le asigna un ID único m.NewID a esta nueva instancia de música.
	addMusic.ID = m.NewID

	// Incrementar el ID para la siguiente música
	m.NewID++

	// Añadir la nueva música a la lista
	//se utiliza para agregar elementos a un slice existente
	m.Music = append(m.Music, addMusic)

	//devuelve nil, indicando si todo salió bien y que no hubo errores.
	return nil
}

func (m *MusicRepository) Update(id int, musicUpdate *domain.Music) error {
	//// Recorrer la lista de músicas en el repositorio
	//range proporciona una forma común de iterar sobre estructuras como slices, arrays, strings, mapas
	for i, music := range m.Music {
		// Buscar la música por su ID
		if music.ID == id {
			// Mantener el ID original
			musicUpdate.ID = id
			// Actualizar la entrada de música
			m.Music[i] = *musicUpdate
			//Retornar nil indicando que la operación fue exitosa
			return nil
		}
	}
	//Retornar nil si no se encontró ninguna música con el ID
	return nil
}

func (m *MusicRepository) Delete(id int) error {
	//// Iterar sobre el slice de música en el repositorio
	for i, music := range m.Music {
		//Verificar si el ID de la música actual coincide con el ID solicitado
		if music.ID == id {
			// Crear un nuevo slice omitiendo el elemento en el índice i
			/*3
			[1 2 3 5]
			[1 2 5][3]
			[1 2 5]
			*/
			before := m.Music[:i]  // Elementos antes del índice i
			after := m.Music[i+1:] // Elementos después del índice i

			// Concatenar los dos slices usando append
			m.Music = append(before, after...)

			//devuelve nil, indicando que la operación de eliminación se realizó correctamente
			return nil
		}
	}
	//Si no se encuentra ninguna música con el ID especificado
	//el método devuelve un error usando errors.New.
	return errors.New("music with requested ID not found")
}

func (m *MusicRepository) GetAll() ([]domain.Music, error) {
	//Devuelve el slice m.Music, que contiene todos los elementos de música almacenados en el repositorio
	return m.Music, nil
}
