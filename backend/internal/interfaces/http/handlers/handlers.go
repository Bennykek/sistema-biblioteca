package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"biblioteca/internal/domain/autor"
	"biblioteca/internal/domain/libro"
	"biblioteca/internal/interfaces/http/facade"
)

type Handler struct { f *facade.BibliotecaFacade }
func New(f *facade.BibliotecaFacade)*Handler{return &Handler{f:f}}

func writeJSON(w http.ResponseWriter,status int,v any){w.Header().Set("Content-Type","application/json");w.WriteHeader(status);_ = json.NewEncoder(w).Encode(v)}
func fail(w http.ResponseWriter,err error){
	status:=http.StatusInternalServerError
	if errors.Is(err,autor.ErrNoEncontrado)||errors.Is(err,libro.ErrNoEncontrado){status=http.StatusNotFound}
	if errors.Is(err,autor.ErrNombreRequerido)||errors.Is(err,libro.ErrDatosInvalidos){status=http.StatusBadRequest}
	if strings.Contains(strings.ToLower(err.Error()),"duplicate key"){status=http.StatusConflict}
	if strings.Contains(strings.ToLower(err.Error()),"foreign key"){status=http.StatusConflict}
	writeJSON(w,status,map[string]string{"error":err.Error()})
}
func idFrom(r *http.Request)(int64,error){return strconv.ParseInt(r.PathValue("id"),10,64)}

func (h *Handler) Autores(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case http.MethodGet:
		if r.PathValue("id")!="" { id,e:=idFrom(r);if e!=nil{fail(w,e);return};v,e:=h.f.ObtenerAutor(id);if e!=nil{fail(w,e);return};writeJSON(w,200,v);return }
		v,e:=h.f.ListarAutores();if e!=nil{fail(w,e);return};writeJSON(w,200,v)
	case http.MethodPost:
		var v autor.Autor;if e:=json.NewDecoder(r.Body).Decode(&v);e!=nil{fail(w,e);return};if e:=h.f.CrearAutor(&v);e!=nil{fail(w,e);return};writeJSON(w,201,v)
	case http.MethodPut:
		id,e:=idFrom(r);if e!=nil{fail(w,e);return};var v autor.Autor;if e=json.NewDecoder(r.Body).Decode(&v);e!=nil{fail(w,e);return};v.ID=id;if e=h.f.ActualizarAutor(v);e!=nil{fail(w,e);return};writeJSON(w,200,v)
	case http.MethodDelete:
		id,e:=idFrom(r);if e!=nil{fail(w,e);return};if e=h.f.EliminarAutor(id);e!=nil{fail(w,e);return};writeJSON(w,200,map[string]string{"mensaje":"Autor eliminado"})
	default:w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (h *Handler) Libros(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case http.MethodGet:
		if r.PathValue("id")!="" { id,e:=idFrom(r);if e!=nil{fail(w,e);return};v,e:=h.f.ObtenerLibro(id);if e!=nil{fail(w,e);return};writeJSON(w,200,v);return }
		v,e:=h.f.ListarLibros();if e!=nil{fail(w,e);return};writeJSON(w,200,v)
	case http.MethodPost:
		var v libro.Libro;if e:=json.NewDecoder(r.Body).Decode(&v);e!=nil{fail(w,e);return};if e:=h.f.CrearLibro(&v);e!=nil{fail(w,e);return};writeJSON(w,201,v)
	case http.MethodPut:
		id,e:=idFrom(r);if e!=nil{fail(w,e);return};var v libro.Libro;if e=json.NewDecoder(r.Body).Decode(&v);e!=nil{fail(w,e);return};v.ID=id;if e=h.f.ActualizarLibro(v);e!=nil{fail(w,e);return};writeJSON(w,200,v)
	case http.MethodDelete:
		id,e:=idFrom(r);if e!=nil{fail(w,e);return};if e=h.f.EliminarLibro(id);e!=nil{fail(w,e);return};writeJSON(w,200,map[string]string{"mensaje":"Libro eliminado"})
	default:w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (h *Handler) Resumen(w http.ResponseWriter,r *http.Request){v,e:=h.f.ResumenInventario();if e!=nil{fail(w,e);return};writeJSON(w,200,v)}
