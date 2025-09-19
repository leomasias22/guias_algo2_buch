package tdas

import "ejemplo.com/guias/typos/pila"

func OrdenarPila(p pila.Pila[int]) {
	// Ordena una pila de enteros en orden ascendente usando solo una pila auxiliar.
	// Utiliza únicamente las operaciones: Apilar, Desapilar, VerTope, EstaVacia, CrearPila
	//
	// Complejidad temporal: O(n²)
	// Complejidad espacial: O(n)

	pilaAuxiliar := pila.CrearPilaDinamica[int]()

	for !p.EstaVacia() {
		// Tomar el elemento actual de la pila original
		elementoActual := p.Desapilar()

		// Mover elementos mayores de la pila auxiliar a la original
		for !pilaAuxiliar.EstaVacia() && pilaAuxiliar.VerTope() > elementoActual {
			p.Apilar(pilaAuxiliar.Desapilar())
		}

		// Insertar el elemento actual en su posición correcta
		pilaAuxiliar.Apilar(elementoActual)
	}

	// Transferir todos los elementos ordenados de vuelta a la pila original
	for !pilaAuxiliar.EstaVacia() {
		p.Apilar(pilaAuxiliar.Desapilar())
	}
}
