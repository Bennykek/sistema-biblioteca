package libro

import (
	"strings"
	"biblioteca/internal/domain/autor"
	domain "biblioteca/internal/domain/libro"
)

type Service struct { repo domain.Repository; autores autor.Repository }
func NewService(r domain.Repository, a autor.Repository) *Service { return &Service{repo:r, autores:a} }

func (s *Service) Listar() ([]domain.Libro,error) { return s.repo.Listar() }
func (s *Service) Obtener(id int64) (domain.Libro,error) { return s.repo.Obtener(id) }
func (s *Service) Crear(l *domain.Libro) error {
	l.Titulo = strings.TrimSpace(l.Titulo); l.ISBN = strings.TrimSpace(l.ISBN)
	if l.Titulo == "" || l.ISBN == "" || l.Anio < 1000 || l.Anio > 2100 || l.AutorID <= 0 { return domain.ErrDatosInvalidos }
	if _,err := s.autores.Obtener(l.AutorID); err != nil { return err }
	return s.repo.Crear(l)
}
func (s *Service) Actualizar(l domain.Libro) error {
	l.Titulo = strings.TrimSpace(l.Titulo); l.ISBN = strings.TrimSpace(l.ISBN)
	if l.Titulo == "" || l.ISBN == "" || l.Anio < 1000 || l.Anio > 2100 || l.AutorID <= 0 { return domain.ErrDatosInvalidos }
	if _,err := s.autores.Obtener(l.AutorID); err != nil { return err }
	return s.repo.Actualizar(l)
}
func (s *Service) Eliminar(id int64) error { return s.repo.Eliminar(id) }

// Resumen usa Visitor para ejecutar una operación transversal sobre la colección.
func (s *Service) Resumen() (domain.InventarioVisitor,error) {
	libros,err := s.repo.Listar()
	if err != nil { return domain.InventarioVisitor{},err }
	visitor := domain.InventarioVisitor{}
	for _, l := range libros { if err := l.Accept(&visitor); err != nil { return visitor,err } }
	return visitor,nil
}
