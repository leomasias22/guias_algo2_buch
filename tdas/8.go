package tdas

import "ejemplo.com/guias/typos/cola"

func funcMultiPrimeros[T any](c cola.Cola[T], k int) []T {
	if k <= 0 || c.EstaVacia() {
		return []T{}
	}

	var resultado []T
	var todos []T

	// Desencolar todos los elementos
	for !c.EstaVacia() {
		elemento := c.Desencolar()
		todos = append(todos, elemento)

		// Guardar solo los primeros k en resultado
		if len(resultado) < k {
			resultado = append(resultado, elemento)
		}
	}

	// Restaurar la cola
	for _, elemento := range todos {
		c.Encolar(elemento)
	}

	return resultado
}
