package divycon

func raizDC(f func(int) int, a, b int) int {
	if f(a) == 0 {
		return a
	}
	if f(b) == 0 {
		return b
	}

	if f(a)*f(b) > 0 {
		panic("No se cumple el teorema de Bolzano: f(a) y f(b) tienen el mismo signo")
	}

	// Caso base: intervalo de tamaño 1
	if a == b {
		return a
	}

	mid := (a + b) / 2
	val := f(mid)

	if val == 0 {
		return mid
	} else if val*f(a) < 0 {
		// La raíz está en la mitad izquierda
		return raizDC(f, a, mid)
	} else {
		// La raíz está en la mitad derecha
		return raizDC(f, mid+1, b)
	}
}

/*
----------------------------Planteo inicial----------------------------
La función f debe cumplir el teorema de Bolzano, es decir, f(a) y f(b) deben tener signos distintos
para asegurar que existe al menos una raíz en el intervalo.

----------------------------Implementación----------------------------
La función raizDC(f, a, b) se implementa recursivamente de la siguiente manera:

1. Comprobar si f(a) o f(b) es igual a 0, en cuyo caso se devuelve directamente la raíz.
2. Verificar que f(a) y f(b) tengan signos distintos. Si no, se detiene el programa.
3. Calcular el punto medio del intervalo: mid = (a + b)/2.
4. Evaluar f(mid):
   - Si f(mid) == 0, mid es la raíz.
   - Si f(a)*f(mid) < 0, la raíz está en [a, mid], se llama recursivamente.
   - Si no, la raíz está en [mid+1, b], se llama recursivamente.
5. Repetir hasta que el intervalo se reduzca a un solo punto.

----------------------------Aplicando el Teorema Maestro----------------------------
La recurrencia de la función es:
    T(n) = A * T(n/B) + O(n^C)

En nuestro caso:
- A = 1 (un subproblema por llamada)
- B = 2 (el tamaño del intervalo se divide a la mitad)
- C = 0 (trabajo constante en cada llamada, O(1))

Evaluamos log_B(A) = log_2(1) = 0

Como log_2(1) = 0 y C = 0, tenemos que log_B(A) = C → T(n) = O(log n)

----------------------------Respuesta----------------------------
como n es el intervarlo entre a y b --> O(log n)
*/
