package facade

import (
	autorapp "biblioteca/internal/application/autor"
	libroapp "biblioteca/internal/application/libro"
	"biblioteca/internal/domain/autor"
	"biblioteca/internal/domain/libro"
)

// BibliotecaFacade ofrece una entrada unificada a los casos de uso.
type BibliotecaFacade struct { Autores *autorapp.Service; Libros *libroapp.Service }
func New(a *autorapp.Service,l *libroapp.Service) *BibliotecaFacade { return &BibliotecaFacade{Autores:a,Libros:l} }

func (f *BibliotecaFacade) ListarAutores()([]autor.Autor,error){return f.Autores.Listar()}
func (f *BibliotecaFacade) ObtenerAutor(id int64)(autor.Autor,error){return f.Autores.Obtener(id)}
func (f *BibliotecaFacade) CrearAutor(a *autor.Autor)error{return f.Autores.Crear(a)}
func (f *BibliotecaFacade) ActualizarAutor(a autor.Autor)error{return f.Autores.Actualizar(a)}
func (f *BibliotecaFacade) EliminarAutor(id int64)error{return f.Autores.Eliminar(id)}
func (f *BibliotecaFacade) ListarLibros()([]libro.Libro,error){return f.Libros.Listar()}
func (f *BibliotecaFacade) ObtenerLibro(id int64)(libro.Libro,error){return f.Libros.Obtener(id)}
func (f *BibliotecaFacade) CrearLibro(l *libro.Libro)error{return f.Libros.Crear(l)}
func (f *BibliotecaFacade) ActualizarLibro(l libro.Libro)error{return f.Libros.Actualizar(l)}
func (f *BibliotecaFacade) EliminarLibro(id int64)error{return f.Libros.Eliminar(id)}
func (f *BibliotecaFacade) ResumenInventario()(libro.InventarioVisitor,error){return f.Libros.Resumen()}
