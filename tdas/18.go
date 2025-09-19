package tdas

import "ejemplo.com/guias/typos/pila"

func ContarElementos[T any](p pila.Pila[T]) int {
	if p.EstaVacia() {
		return 0
	}

	elemento := p.Desapilar()
	contador := 1 + ContarElementos(p)
	p.Apilar(elemento)

	return contador
}
