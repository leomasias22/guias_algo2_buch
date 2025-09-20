package ordenamientos

type Persona struct {
	Edad         int
	Nombre       string
	Nacionalidad int // 0-31
}

// OrdenarFilaCajita - Algoritmo principal de ordenamiento lineal
func OrdenarFilaCajita(personas []Persona) {
	if len(personas) <= 1 {
		return
	}

	// Paso 1: Separar niños y adultos
	var niños, adultos []Persona

	for _, persona := range personas {
		if persona.Edad <= 12 {
			niños = append(niños, persona)
		} else {
			adultos = append(adultos, persona)
		}
	}

	// Paso 2: Ordenar niños SOLO por edad
	if len(niños) > 0 {
		CountingSort(niños, "edad")
	}

	// Paso 3: Ordenar adultos usando Radix Sort
	// Primero por edad (menos significativo), luego por nacionalidad (más significativo)
	if len(adultos) > 0 {
		CountingSort(adultos, "edad")         // Menos significativo
		CountingSort(adultos, "nacionalidad") // Más significativo
	}

	// Paso 4: Concatenar resultado (niños primero, luego adultos)
	copy(personas[:len(niños)], niños)
	copy(personas[len(niños):], adultos)
}

func CountingSort(personas []Persona, criterio string) {}
