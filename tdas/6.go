package tdas

import "ejemplo.com/guias/typos/pila"

func EstaOrdenadaAscendente(p pila.Pila[int]) bool {
	// si la pila está vacía, se considera ordenada
	if p.EstaVacia() {
		return true
	}
	pilaAuxiliar1 := pila.CrearPilaDinamica[int]()
	pilaAuxiliar2 := pila.CrearPilaDinamica[int]()
	elementoAnterior := p.Desapilar()
	pilaAuxiliar1.Apilar(elementoAnterior)
	// Recorrer el resto de elementos
	for !p.EstaVacia() {
		elementoActual := p.Desapilar()
		pilaAuxiliar1.Apilar(elementoActual)
		if elementoAnterior >= elementoActual {
			return false
		}
		elementoAnterior = elementoActual
	}
	for !pilaAuxiliar1.EstaVacia() {
		elemento := pilaAuxiliar1.Desapilar()
		pilaAuxiliar2.Apilar(elemento)
	}
	for !pilaAuxiliar2.EstaVacia() {
		elemento := pilaAuxiliar2.Desapilar()
		p.Apilar(elemento)
	}
	return true
}
