package pila

const (
	CONSTANTE_REDIMENSION   = 2
	CONSTANTE_CREACION_PILA = 10
	CONSTANTE_DECRECION     = 4
	MENSAJE_PANIC_PILA      = "La pila esta vacia"
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, CONSTANTE_CREACION_PILA),
		cantidad: 0,
	}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(MENSAJE_PANIC_PILA)
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(p.calcularNuevaCapacidad(true))
	}

	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic(MENSAJE_PANIC_PILA)
	}
	elemento := p.datos[p.cantidad-1]

	p.cantidad--
	if p.cantidad > CONSTANTE_CREACION_PILA && p.cantidad <= len(p.datos)/CONSTANTE_DECRECION {
		p.redimensionar(p.calcularNuevaCapacidad(false))
	}

	return elemento
}

/* Funciones auxiliares */

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {

	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

func (p *pilaDinamica[T]) calcularNuevaCapacidad(aumentar bool) int {
	capacidadActual := cap(p.datos)

	if aumentar {
		if capacidadActual == 0 {
			return CONSTANTE_CREACION_PILA
		}
		return capacidadActual * CONSTANTE_REDIMENSION
	} else {
		return capacidadActual / CONSTANTE_REDIMENSION
	}
}

///*ejericio 17guia*//

func (pila pilaDinamica[T]) Transformar(aplicar func(T) T) Pila[T] {
	// Crear una nueva pila para almacenar los elementos transformados
	nuevaPila := CrearPilaDinamica[T]()

	// Crear una pila auxiliar para preservar el orden original
	pilaAuxiliar := CrearPilaDinamica[T]()

	// Paso 1: Desapilar todos los elementos de la pila original
	// y apilarlos en la pila auxiliar (invierte el orden)
	for !pila.EstaVacia() {
		pilaAuxiliar.Apilar(pila.Desapilar())
	}

	// Paso 2: Procesar elementos desde la pila auxiliar
	// Los elementos salen en el orden original (desde el fondo hacia el tope)
	for !pilaAuxiliar.EstaVacia() {
		elemento := pilaAuxiliar.Desapilar()

		// Restaurar el elemento en la pila original
		pila.Apilar(elemento)

		// Aplicar la transformación y agregar a la nueva pila
		elementoTransformado := aplicar(elemento)
		nuevaPila.Apilar(elementoTransformado)
	}

	return nuevaPila
}
