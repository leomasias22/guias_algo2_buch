package codigo

// BucketSort ordena cualquier tipo usando bucket sort con función de mapeo
func BucketSort[T any](arr []T, keyFunc func(T) int) []T {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// Encontrar min y max
	min, max := keyFunc(arr[0]), keyFunc(arr[0])
	for _, item := range arr {
		key := keyFunc(item)
		if key < min {
			min = key
		}
		if key > max {
			max = key
		}
	}

	// Si todos son iguales
	if min == max {
		return arr
	}

	// Crear buckets
	buckets := make([][]T, n)
	range_ := max - min

	// Distribuir elementos en buckets
	for _, item := range arr {
		key := keyFunc(item)
		index := (key - min) * (n - 1) / range_
		if index >= n {
			index = n - 1
		}
		buckets[index] = append(buckets[index], item)
	}

	// Ordenar cada bucket usando insertion sort
	for i := range buckets {
		insertionSort(buckets[i], keyFunc)
	}

	// Concatenar buckets
	result := make([]T, 0, n)
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}

// insertionSort ordena un slice usando insertion sort
func insertionSort[T any](arr []T, keyFunc func(T) int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		keyValue := keyFunc(key)
		j := i - 1

		for j >= 0 && keyFunc(arr[j]) > keyValue {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
