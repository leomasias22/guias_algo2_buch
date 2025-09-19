package tdas

import "ejemplo.com/guias/typos/pila"

func OrdenarPila(p pila.Pila[int]) {
	// Ordena una pila de enteros en orden ascendente usando solo una pila auxiliar.
	// Complejidad : O(n²)

	pilaAuxiliar := pila.CrearPilaDinamica[int]()

	for !p.EstaVacia() {
		// Tomar el elemento actual de la pila original
		elementoActual := p.Desapilar()

		// Mover elementos mayores de la pila auxiliar a la original
		for !pilaAuxiliar.EstaVacia() && pilaAuxiliar.VerTope() > elementoActual {
			mayor := pilaAuxiliar.Desapilar()
			p.Apilar(mayor)
		}

		// Insertar el elemento actual en su posición correcta
		pilaAuxiliar.Apilar(elementoActual)
	}

	// Transferir todos los elementos ordenados de vuelta a la pila original
	for !pilaAuxiliar.EstaVacia() {
		elemento := pilaAuxiliar.Desapilar()
		p.Apilar(elemento)

	}
}
