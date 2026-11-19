package persistence

import (
	"database/sql"
	"biblioteca/internal/domain/autor"
)

type AutorRepository struct { db *sql.DB }
func NewAutorRepository(db *sql.DB) *AutorRepository { return &AutorRepository{db:db} }

func (r *AutorRepository) Listar() ([]autor.Autor,error) {
	rows,err := r.db.Query(`SELECT id,nombre,nacionalidad FROM autores ORDER BY id DESC`); if err != nil { return nil,err }; defer rows.Close()
	result := []autor.Autor{}
	for rows.Next() { var a autor.Autor; if err := rows.Scan(&a.ID,&a.Nombre,&a.Nacionalidad); err != nil { return nil,err }; result=append(result,a) }
	return result,rows.Err()
}
func (r *AutorRepository) Obtener(id int64) (autor.Autor,error) {
	var a autor.Autor
	err := r.db.QueryRow(`SELECT id,nombre,nacionalidad FROM autores WHERE id=$1`,id).Scan(&a.ID,&a.Nombre,&a.Nacionalidad)
	if err == sql.ErrNoRows { return a,autor.ErrNoEncontrado }; return a,err
}
func (r *AutorRepository) Crear(a *autor.Autor) error {
	return r.db.QueryRow(`INSERT INTO autores(nombre,nacionalidad) VALUES($1,$2) RETURNING id`,a.Nombre,a.Nacionalidad).Scan(&a.ID)
}
func (r *AutorRepository) Actualizar(a autor.Autor) error {
	res,err := r.db.Exec(`UPDATE autores SET nombre=$1,nacionalidad=$2 WHERE id=$3`,a.Nombre,a.Nacionalidad,a.ID); if err != nil { return err }
	n,_ := res.RowsAffected(); if n==0 { return autor.ErrNoEncontrado }; return nil
}
func (r *AutorRepository) Eliminar(id int64) error {
	res,err := r.db.Exec(`DELETE FROM autores WHERE id=$1`,id); if err != nil { return err }
	n,_ := res.RowsAffected(); if n==0 { return autor.ErrNoEncontrado }; return nil
}
