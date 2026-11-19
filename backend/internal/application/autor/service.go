package autor

import (
	domain "biblioteca/internal/domain/autor"
	"strings"
)

type Service struct { repo domain.Repository }
func NewService(r domain.Repository) *Service { return &Service{repo:r} }

func (s *Service) Listar() ([]domain.Autor,error) { return s.repo.Listar() }
func (s *Service) Obtener(id int64) (domain.Autor,error) { return s.repo.Obtener(id) }
func (s *Service) Crear(a *domain.Autor) error {
	a.Nombre = strings.TrimSpace(a.Nombre)
	a.Nacionalidad = strings.TrimSpace(a.Nacionalidad)
	if err := a.Validar(); err != nil { return err }
	return s.repo.Crear(a)
}
func (s *Service) Actualizar(a domain.Autor) error {
	a.Nombre = strings.TrimSpace(a.Nombre)
	a.Nacionalidad = strings.TrimSpace(a.Nacionalidad)
	if err := a.Validar(); err != nil { return err }
	return s.repo.Actualizar(a)
}
func (s *Service) Eliminar(id int64) error { return s.repo.Eliminar(id) }
