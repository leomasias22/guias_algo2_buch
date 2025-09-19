package cola

const (
	MENSAJE_PANIC_COLA = "La cola esta vacia"
)

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
	}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(MENSAJE_PANIC_COLA)
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elemento T) {
	nuevoNodo := &nodoCola[T]{
		dato: elemento,
		prox: nil,
	}

	if c.EstaVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else {
		c.ultimo.prox = nuevoNodo
		c.ultimo = nuevoNodo
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic(MENSAJE_PANIC_COLA)
	}

	elemento := c.primero.dato
	c.primero = c.primero.prox

	if c.primero == nil {
		c.ultimo = nil
	}

	return elemento
}

//ejercicios numero 7

func (c *colaEnlazada[T]) MultiPrimeros(k int) []T {
	// Validaciones
	if k <= 0 || c.primero == nil {
		return []T{}
	}

	var resultado []T
	actual := c.primero // Acceso directo al primer nodo

	// Recorrer hasta k elementos o hasta el final
	for i := 0; i < k && actual != nil; i++ {
		resultado = append(resultado, actual.dato)
		actual = actual.prox
	}

	return resultado
}
