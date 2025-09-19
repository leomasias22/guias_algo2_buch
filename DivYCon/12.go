package divycon

func ParteEnteraRaiz(n int) int {
	return parteEnteraRaiz(n, 0, n)
}

func parteEnteraRaiz(n, inicio, fin int) int {
	if n <= 1 {
		return n
	}

	medio := (inicio + fin) / 2

	if medio*medio <= n && (medio+1)*(medio+1) > n {
		return medio
	}

	if medio*medio > n {
		return parteEnteraRaiz(n, inicio, medio-1)
	}

	return parteEnteraRaiz(n, medio+1, fin)
}

/*
ANÁLISIS DE COMPLEJIDAD MEDIANTE TEOREMA MAESTRO:

La recurrencia de este algoritmo es:
T(n) = T(n/2) + O(1)

Aplicando el Teorema Maestro:
- a = 1, b = 2, c = 0 (ya que O(n^C) = O(1) = n^0)
- LogB(a) = log_2(1) = 0
- Comparación: c = 0 = 0 → Caso 2

Por lo tanto, la complejidad es: T(n) = O(log n)
*/
