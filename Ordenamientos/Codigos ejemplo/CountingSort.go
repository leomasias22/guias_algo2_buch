package codigo

func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Encontrar min y max
	min, max := arr[0], arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	// Crear array de conteo
	range_ := max - min + 1
	count := make([]int, range_)

	// Contar ocurrencias
	for _, val := range arr {
		count[val-min]++
	}

	// Convertir a posiciones acumulativas
	for i := 1; i < range_; i++ {
		count[i] += count[i-1]
	}

	// Construir array resultado
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := arr[i] - min
		count[index]--
		result[count[index]] = arr[i]
	}

	return result
}
