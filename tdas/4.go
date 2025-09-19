package tdas

type ComposicionFunciones struct {
	funciones []func(float64) float64
}

// CrearComposicion crea una nueva instancia de ComposicionFunciones
func CrearComposicion() ComposicionFunciones {
	return ComposicionFunciones{
		funciones: make([]func(float64) float64, 0),
	}
}

// AgregarFuncion agrega una función a la composición
// Las funciones se agregan en el orden de lectura: f, g, h
// Pero se aplicarán en orden inverso: f(g(h(x)))
func (c *ComposicionFunciones) AgregarFuncion(f func(float64) float64) {
	c.funciones = append(c.funciones, f)
}

// Aplicar aplica todas las funciones compuestas al valor dado
// Aplica las funciones en orden inverso al que fueron agregadas
// Si se agregaron f, g, h entonces aplica f(g(h(x)))
func (c *ComposicionFunciones) Aplicar(x float64) float64 {
	resultado := x

	// Aplicamos las funciones en orden inverso (desde la última hacia la primera)
	for i := len(c.funciones) - 1; i >= 0; i-- {
		resultado = c.funciones[i](resultado)
	}

	return resultado
}
