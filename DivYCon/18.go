package divycon

func pcr(grupo []int) bool {
	// Ejemplo: la persona 7 está contagiada
	contagiado := 7
	for _, persona := range grupo {
		if persona == contagiado {
			return true
		}
	}
	return false
}

// Algoritmo de búsqueda del contagiado usando división y conquista
func buscarContagiado(personas []int) []int {
	if len(personas) == 0 {
		return []int{}
	}
	if len(personas) == 1 {
		if pcr(personas) {
			return personas
		}
		return []int{}
	}

	mitad := len(personas) / 2
	primeraMitad := personas[:mitad]
	segundaMitad := personas[mitad:]

	if pcr(primeraMitad) {
		return buscarContagiado(primeraMitad)
	} else {
		return buscarContagiado(segundaMitad)
	}
}

/*Explicacion
Si bien en el algoritmo se encuentra un for(el cual es necesario ya que no sabemos
quien es la persona contagiada)
Este algoritmo sera siempre O(n) ya que deberemos recorrer todo le arreglo de personas.
El caso donde sera O(log(n)) sera cuando se nos indique la posicion del infectado y debamos buscarlo.
En dicho caso una busquda binario o divide y conquista clasico sera suficiente.
*/
