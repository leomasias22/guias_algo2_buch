package tdas

import "ejemplo.com/guias/typos/cola"

func FiltrarCola[K any](c cola.Cola[K], filtro func(K) bool) {
	// Cola auxiliar para almacenar temporalmente los elementos
	colaAux := cola.CrearColaEnlazada[K]()

	// Primera pasada: desencolamos todos los elementos y aplicamos el filtro
	for !c.EstaVacia() {
		elemento := c.Desencolar()

		// Si el elemento pasa el filtro, lo guardamos en la cola auxiliar
		if filtro(elemento) {
			colaAux.Encolar(elemento)
		}
		// Si no pasa el filtro, simplemente lo descartamos
	}

	// Segunda pasada: devolvemos los elementos filtrados a la cola original
	for !colaAux.EstaVacia() {
		c.Encolar(colaAux.Desencolar())
	}
}
