package autor

import "errors"

var ErrNombreRequerido = errors.New("el nombre del autor es obligatorio")
var ErrNoEncontrado = errors.New("autor no encontrado")

type Repository interface {
	Listar() ([]Autor, error)
	Obtener(id int64) (Autor, error)
	Crear(a *Autor) error
	Actualizar(a Autor) error
	Eliminar(id int64) error
}
