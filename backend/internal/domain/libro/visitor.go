package libro

// Visitor permite agregar operaciones sobre Libro sin modificar su estructura.
type Visitor interface {
	VisitLibro(Libro) error
}

type InventarioVisitor struct {
	Cantidad int
	AnioMasAntiguo int
	AnioMasReciente int
}

func (v *InventarioVisitor) VisitLibro(l Libro) error {
	v.Cantidad++
	if v.AnioMasAntiguo == 0 || l.Anio < v.AnioMasAntiguo { v.AnioMasAntiguo = l.Anio }
	if l.Anio > v.AnioMasReciente { v.AnioMasReciente = l.Anio }
	return nil
}
