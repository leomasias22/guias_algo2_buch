package codigo

func RadixSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Encontrar el valor máximo
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	// Ordenar por cada dígito
	result := make([]int, len(arr))
	copy(result, arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		// Counting sort por dígito actual
		n := len(result)
		output := make([]int, n)
		count := make([]int, 10)

		// Contar ocurrencias de cada dígito
		for i := 0; i < n; i++ {
			digit := (result[i] / exp) % 10
			count[digit]++
		}

		// Convertir a posiciones acumulativas
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		// Construir array resultado
		for i := n - 1; i >= 0; i-- {
			digit := (result[i] / exp) % 10
			output[count[digit]-1] = result[i]
			count[digit]--
		}

		result = output
	}

	return result
}
