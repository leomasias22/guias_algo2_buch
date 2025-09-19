package divycon

func elementoFueraDelugar(arr []int, inicio, fin int) int {
	// Caso base: un solo elemento
	if inicio == fin {
		// Verificar si está fuera de lugar comparando con vecinos
		if (inicio > 0 && arr[inicio] < arr[inicio-1]) ||
			(inicio < len(arr)-1 && arr[inicio] > arr[inicio+1]) {
			return arr[inicio]
		}
		return -1
	}

	// División
	medio := inicio + (fin-inicio)/2

	// Conquista: buscar en ambas mitades
	izq := elementoFueraDelugar(arr, inicio, medio)
	if izq != -1 {
		return izq
	}
	return elementoFueraDelugar(arr, medio+1, fin)
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
// si bien pareciera que solo llama a una recursion en
//	izq := elementoFueraDelugar(arr, inicio, medio)
// realmente solo se efectua una cuando se cumple el if
// por lo tanto en todos los demas casos se efectuara tambien el 2do llamado
// en el ultimo return
