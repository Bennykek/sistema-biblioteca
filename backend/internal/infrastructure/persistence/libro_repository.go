package persistence

import (
	"database/sql"
	"biblioteca/internal/domain/libro"
)

type LibroRepository struct { db *sql.DB }
func NewLibroRepository(db *sql.DB) *LibroRepository { return &LibroRepository{db:db} }

const libroSelect = `SELECT l.id,l.titulo,l.isbn,l.anio,l.autor_id,a.nombre FROM libros l JOIN autores a ON a.id=l.autor_id`
func (r *LibroRepository) Listar() ([]libro.Libro,error) {
	rows,err := r.db.Query(libroSelect+` ORDER BY l.id DESC`); if err != nil { return nil,err }; defer rows.Close()
	result := []libro.Libro{}
	for rows.Next() { var l libro.Libro; if err:=rows.Scan(&l.ID,&l.Titulo,&l.ISBN,&l.Anio,&l.AutorID,&l.Autor); err!=nil{return nil,err}; result=append(result,l) }
	return result,rows.Err()
}
func (r *LibroRepository) Obtener(id int64) (libro.Libro,error) {
	var l libro.Libro
	err := r.db.QueryRow(libroSelect+` WHERE l.id=$1`,id).Scan(&l.ID,&l.Titulo,&l.ISBN,&l.Anio,&l.AutorID,&l.Autor)
	if err==sql.ErrNoRows { return l,libro.ErrNoEncontrado }; return l,err
}
func (r *LibroRepository) Crear(l *libro.Libro) error {
	return r.db.QueryRow(`INSERT INTO libros(titulo,isbn,anio,autor_id) VALUES($1,$2,$3,$4) RETURNING id`,l.Titulo,l.ISBN,l.Anio,l.AutorID).Scan(&l.ID)
}
func (r *LibroRepository) Actualizar(l libro.Libro) error {
	res,err:=r.db.Exec(`UPDATE libros SET titulo=$1,isbn=$2,anio=$3,autor_id=$4 WHERE id=$5`,l.Titulo,l.ISBN,l.Anio,l.AutorID,l.ID); if err!=nil{return err}
	n,_:=res.RowsAffected(); if n==0{return libro.ErrNoEncontrado}; return nil
}
func (r *LibroRepository) Eliminar(id int64) error {
	res,err:=r.db.Exec(`DELETE FROM libros WHERE id=$1`,id); if err!=nil{return err}
	n,_:=res.RowsAffected(); if n==0{return libro.ErrNoEncontrado}; return nil
}
