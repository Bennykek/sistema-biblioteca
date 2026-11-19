package autor

type Autor struct {
	ID           int64  `json:"id"`
	Nombre       string `json:"nombre"`
	Nacionalidad string `json:"nacionalidad"`
}

func (a Autor) Validar() error {
	if a.Nombre == "" { return ErrNombreRequerido }
	return nil
}
