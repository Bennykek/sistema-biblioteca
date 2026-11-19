package libro

import "errors"

var ErrNoEncontrado = errors.New("libro no encontrado")
var ErrDatosInvalidos = errors.New("titulo, ISBN, año y autor son obligatorios")

type Repository interface {
	Listar() ([]Libro, error)
	Obtener(id int64) (Libro, error)
	Crear(l *Libro) error
	Actualizar(l Libro) error
	Eliminar(id int64) error
}
