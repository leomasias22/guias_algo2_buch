package tdas

type ColaAcotada struct {
	elementos []int
	frente    int
	final     int
	cantidad  int
	capacidad int
}

func NuevaColaAcotada(k int) *ColaAcotada {
	return &ColaAcotada{
		elementos: make([]int, k),
		frente:    0,
		final:     0,
		cantidad:  0,
		capacidad: k,
	}
}

func (c *ColaAcotada) Encolar(elemento int) {
	c.elementos[c.final] = elemento
	c.final = (c.final + 1) % c.capacidad
	c.cantidad++
}

func (c *ColaAcotada) Desencolar() int {
	elemento := c.elementos[c.frente]
	c.frente = (c.frente + 1) % c.capacidad
	c.cantidad--
	return elemento
}
