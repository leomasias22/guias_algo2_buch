package tdas

type nodoLista[T any] struct {
	prox *nodoLista[T]
	dato T
}

type ListaEnlazada[T any] struct {
	prim *nodoLista[T]
}

// Primitiva que devuelve el elemento a k posiciones del final
func (lista *ListaEnlazada[T]) AnteKUltimo(k int) T {
	// Primer puntero avanza k posiciones
	adelantado := lista.prim
	for i := 0; i < k; i++ {
		adelantado = adelantado.prox
	}

	// Segundo puntero comienza desde el inicio
	atrasado := lista.prim

	// Ambos punteros avanzan juntos hasta que el adelantado llegue al final
	for adelantado != nil {
		adelantado = adelantado.prox
		atrasado = atrasado.prox
	}

	// El puntero atrasado está ahora en la posición k desde el final
	return atrasado.dato
}
