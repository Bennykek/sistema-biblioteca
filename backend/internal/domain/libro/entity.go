package libro

type Libro struct {
	ID      int64  `json:"id"`
	Titulo  string `json:"titulo"`
	ISBN    string `json:"isbn"`
	Anio    int    `json:"anio"`
	AutorID int64  `json:"autor_id"`
	Autor   string `json:"autor,omitempty"`
}

func (l Libro) Accept(v Visitor) error {
	return v.VisitLibro(l)
}
