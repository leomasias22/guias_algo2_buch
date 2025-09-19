package divycon

func encontrarMinimo(arr []int, inicio, fin int) int {
	// Caso base: si hay un solo elemento
	if inicio == fin {
		return arr[inicio]
	}

	// División: encontrar el punto medio
	medio := inicio + (fin-inicio)/2

	// Conquista: resolver recursivamente ambas mitades
	minIzquierda := encontrarMinimo(arr, inicio, medio)
	minDerecha := encontrarMinimo(arr, medio+1, fin)

	// Combinación: devolver el menor de los dos mínimos
	if minIzquierda < minDerecha {
		return minIzquierda
	}
	return minDerecha
}

/*
ANÁLISIS DE COMPLEJIDAD MEDIANTE TEOREMA MAESTRO:
La recurrencia de este algoritmo es:
T(n) = 2T(n/2) + O(1)

Aplicando el Teorema Maestro:
- a = 2, b = 2, c = 0 (ya que f(n) = O(1) = n^0)
- LogB(a) = log_2(2) = 1
- Comparación: c = 0 < 1 → Caso 3

Por lo tanto, la complejidad es: T(n) = O(n)
*/
