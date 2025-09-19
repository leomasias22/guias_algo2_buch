package divycon

func primerCero(arr []int, inicio, fin int) int {
	// Casos base
	if inicio > fin {
		return -1
	}

	if inicio == fin {
		if arr[inicio] == 0 {
			return inicio
		}
		return -1
	}

	// División
	medio := inicio + (fin-inicio)/2

	if arr[medio] == 1 {
		return primerCero(arr, medio+1, fin)
	}

	izq := primerCero(arr, inicio, medio-1)
	if izq != -1 {
		return izq
	}
	return medio
}

/*
ALGORITMO D - PRIMER CERO:

ANÁLISIS DE COMPLEJIDAD MEDIANTE TEOREMA MAESTRO:
La recurrencia de este algoritmo es:
T(n) = T(n/2) + O(1)

Aplicando el Teorema Maestro:
- a = 1, b = 2, c = 0 (ya que f(n) = O(1) = n^0)
- LogB(a) = log_2(1) = 0
- Comparación: c = 0 = 0 → Caso 2

Por lo tanto, la complejidad es: T(n) = O(log n)
*/
