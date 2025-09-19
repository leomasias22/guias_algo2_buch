package tdas

import "ejemplo.com/guias/typos/pila"

func MergePilasSinDuplicados(pila1, pila2 pila.Pila[int]) pila.Pila[int] {
	aux1 := pila.CrearPilaDinamica[int]()
	aux2 := pila.CrearPilaDinamica[int]()
	resultado := pila.CrearPilaDinamica[int]()

	for !pila1.EstaVacia() {
		val := aux1.Desapilar()
		aux1.Apilar(val)
	}

	for !pila2.EstaVacia() {
		val := aux2.Desapilar()
		aux2.Apilar(val)
	}

	for !aux1.EstaVacia() || !aux2.EstaVacia() {
		if aux1.EstaVacia() {
			val := aux2.Desapilar()
			pila2.Apilar(val)
			if resultado.EstaVacia() || resultado.VerTope() != val {
				resultado.Apilar(val)
			}
		} else if aux2.EstaVacia() {
			val := aux1.Desapilar()
			pila1.Apilar(val)
			if resultado.EstaVacia() || resultado.VerTope() != val {
				resultado.Apilar(val)
			}
		} else {
			val1 := aux1.Desapilar()
			val2 := aux2.Desapilar()

			if val1 >= val2 {
				pila2.Apilar(val2)
				if resultado.EstaVacia() || resultado.VerTope() != val2 {
					resultado.Apilar(val2)
				}

				if val1 != val2 {
					aux1.Apilar(val1)
				} else {
					pila1.Apilar(val1)
				}
			} else {
				pila1.Apilar(val1)
				if resultado.EstaVacia() || resultado.VerTope() != val1 {
					resultado.Apilar(val1)
				}
				aux2.Apilar(val2)
			}
		}
	}

	for !resultado.EstaVacia() {
		val := resultado.Desapilar()
		aux1.Apilar(val)
	}

	// Devolver la pila
	return aux1
}
