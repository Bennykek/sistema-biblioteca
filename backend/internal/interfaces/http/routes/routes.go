package routes

import (
	"net/http"
	"biblioteca/internal/interfaces/http/facade"
	"biblioteca/internal/interfaces/http/handlers"
)

func Setup(f *facade.BibliotecaFacade) http.Handler {
	mux:=http.NewServeMux();h:=handlers.New(f)
	mux.HandleFunc("GET /api/health",func(w http.ResponseWriter,r *http.Request){w.Header().Set("Content-Type","application/json");w.Write([]byte(`{"status":"ok"}`))})
	mux.HandleFunc("/api/autores",h.Autores)
	mux.HandleFunc("/api/autores/{id}",h.Autores)
	mux.HandleFunc("/api/libros",h.Libros)
	mux.HandleFunc("/api/libros/{id}",h.Libros)
	mux.HandleFunc("GET /api/reportes/inventario",h.Resumen)
	return cors(mux)
}
func cors(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	if r.Method==http.MethodOptions{w.WriteHeader(http.StatusNoContent);return};next.ServeHTTP(w,r)
})}
