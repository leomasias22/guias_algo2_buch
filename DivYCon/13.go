package divycon

func PosicionPico(v []int, ini, fin int) int {
	if ini == fin {
		return ini
	}

	med := (ini + fin) / 2

	// Comparamos solo con v[med+1], suficiente para arreglo de pico
	if v[med] < v[med+1] {
		return PosicionPico(v, med+1, fin)
	} else {
		return PosicionPico(v, ini, med)
	}
}

/*
ANÁLISIS DE COMPLEJIDAD MEDIANTE TEOREMA MAESTRO:

La recurrencia de este algoritmo es:
T(n) = T(n/2) + O(1)

Aplicando el Teorema Maestro:
- a = 1, b = 2, c = 0 (ya que f(n) = O(1) = n^0)
- LogB(a) = log_2(1) = 0
- Comparación: c = 0 = 0 → Caso 2

Por lo tanto, la complejidad es: T(n) = O(log n)

*/
