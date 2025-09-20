package ordenamientos

type Alumno struct {
	Nombre   string
	Parcial1 int
	Parcial2 int
	Parcial3 int
	Promedio int
}

func RadixSortAlumnos(alumnos []Alumno) {
	// Calcular promedios (parte entera)
	for i := range alumnos {
		suma := alumnos[i].Parcial1 + alumnos[i].Parcial2 + alumnos[i].Parcial3
		alumnos[i].Promedio = suma / 3 // División entera automática en Go
	}

	// Aplicar Counting Sort por cada criterio (menos a más significativo)
	countingSort(alumnos, "parcial3") // Paso 1: Parcial 3
	countingSort(alumnos, "parcial2") // Paso 2: Parcial 2
	countingSort(alumnos, "parcial1") // Paso 3: Parcial 1
	countingSort(alumnos, "promedio") // Paso 4: Promedio
}

// Esto realmente ya estaria implementado, solamente necesitamos el codigo previo a este.
func countingSort(alumnos []Alumno, criterio string) {
	n := len(alumnos)
	if n <= 1 {
		return
	}

	// Crear arreglo de conteo (0-10, k=11)
	count := make([]int, 11)
	output := make([]Alumno, n)

	// Función para obtener el valor según el criterio
	getValor := func(a Alumno) int {
		switch criterio {
		case "promedio":
			return a.Promedio
		case "parcial1":
			return a.Parcial1
		case "parcial2":
			return a.Parcial2
		case "parcial3":
			return a.Parcial3
		default:
			return 0
		}
	}

	// Contar ocurrencias
	for i := 0; i < n; i++ {
		valor := getValor(alumnos[i])
		count[valor]++
	}

	// Transformar a posiciones acumulativas
	for i := 1; i <= 10; i++ {
		count[i] += count[i-1]
	}

	// Construir arreglo ordenado (desde el final para estabilidad)
	for i := n - 1; i >= 0; i-- {
		valor := getValor(alumnos[i])
		output[count[valor]-1] = alumnos[i]
		count[valor]--
	}

	// Copiar de vuelta al arreglo original
	copy(alumnos, output)
}
