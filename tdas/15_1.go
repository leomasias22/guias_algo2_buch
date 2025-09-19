package tdas

type Color int

const (
	Rojo Color = iota
	Azul
	Verde
	Amarillo
	Violeta
)

// Estructura del TDA Mamushka
type Mamushka struct {
	tam      int
	color    Color
	guardada *Mamushka
}

// CrearMamushka crea una mamushka con un tamaño y color definido
// Orden: O(1)
func CrearMamushka(tam int, color Color) *Mamushka {
	return &Mamushka{
		tam:      tam,
		color:    color,
		guardada: nil,
	}
}

// ObtenerColor obtiene el color de la Mamushka
// Orden: O(1)
func (m *Mamushka) ObtenerColor() Color {
	return m.color
}

// Guardar intenta guardar la segunda mamushka en la primera
// Orden: O(n) donde n es la profundidad de mamushkas anidadas
func (m *Mamushka) Guardar(aGuardar *Mamushka) bool {
	// Si no hay mamushka guardada
	if m.guardada == nil {
		// Solo se puede guardar si la mamushka a guardar es más pequeña
		if aGuardar.tam < m.tam {
			m.guardada = aGuardar
			return true
		}
		return false
	}

	// Si ya hay una mamushka guardada, intentar guardar en ella recursivamente
	return m.guardada.Guardar(aGuardar)
}

// ObtenerGuardada devuelve un puntero a la mamushka guardada
// Orden: O(1)
func (m *Mamushka) ObtenerGuardada() *Mamushka {
	return m.guardada
}
