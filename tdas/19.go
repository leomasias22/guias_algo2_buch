package tdas

import "ejemplo.com/guias/typos/pila"

func Balanceado(texto string) bool {
	p := pila.CrearPilaDinamica[rune]()

	for _, char := range texto {
		if char == '(' || char == '[' || char == '{' {
			p.Apilar(char)
		} else if char == ')' || char == ']' || char == '}' {
			if p.EstaVacia() {
				return false
			}

			apertura := p.Desapilar()
			if (apertura == '(' && char != ')') ||
				(apertura == '[' && char != ']') ||
				(apertura == '{' && char != '}') {
				return false
			}
		}
	}

	return p.EstaVacia()
}
