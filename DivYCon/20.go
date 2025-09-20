package divycon

type Joya int

// si bien esta funcion tiene dos for, por enunciado es de tiempo constante, osea O(1)
func balanza(grupo1, grupo2 []Joya) int {
	peso1, peso2 := 0, 0
	for _, j := range grupo1 {
		peso1 += int(j)
	}
	for _, j := range grupo2 {
		peso2 += int(j)
	}
	return peso1 - peso2
}

func EncontrarJoyaVerdadera(joyas []Joya) Joya {
	n := len(joyas)
	if n == 1 {
		return joyas[0]
	}

	mitad := n / 2
	grupo1 := joyas[0:mitad]
	grupo2 := joyas[mitad:n]

	resultado := balanza(grupo1, grupo2)
	if resultado > 0 {
		return EncontrarJoyaVerdadera(grupo1)
	} else if resultado < 0 {
		return EncontrarJoyaVerdadera(grupo2)
	} else {
		return joyas[0]
	}
}

/*
----------------------------Justificación de complejidad----------------------------

Se divide el conjunto de joyas en dos grupos iguales y se pesa uno contra otro.
El grupo más pesado contiene la joya verdadera.
Cada paso reduce el número de joyas aproximadamente a la mitad.

Recurrencia: T(n) = T(n/2) + 1
Resolviendo: T(n) = O(log_2 n)

El algoritmo encuentra la joya verdadera en tiempo logarítmico respecto al número total de joyas.

*/
