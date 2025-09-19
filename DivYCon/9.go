package divycon

func estaOrdenado(arr []int, inicio, fin int) bool {
	// Caso base: si hay un solo elemento o menos, está ordenado
	if inicio >= fin {
		return true
	}

	// División: encontrar el punto medio
	medio := inicio + (fin-inicio)/2

	// Conquista: verificar recursivamente ambas mitades
	izquierdaOrdenada := estaOrdenado(arr, inicio, medio)
	derechaOrdenada := estaOrdenado(arr, medio+1, fin)

	// Combinación: verificar que ambas mitades estén ordenadas
	// y que el último elemento de la izquierda sea <= al primero de la derecha
	return izquierdaOrdenada && derechaOrdenada && (arr[medio] <= arr[medio+1])
}

/*
ANÁLISIS DE COMPLEJIDAD MEDIANTE TEOREMA MAESTRO:
La recurrencia de este algoritmo es:
T(n) = 2T(n/2) + O(1)

Aplicando el Teorema Maestro:
- a = 2, b = 2, c = 0 (ya que O(n^C) = O(1) = n^0)
- LogB(a) = log_2(2) = 1
- Comparación: c = 0 < 1 → Caso 3

Por lo tanto, la complejidad es: T(n) = O(n)
*/
