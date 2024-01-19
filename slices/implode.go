package slices

func Implode[Arr ~[]T, T any](arr ...Arr) Arr {
	resultSliceLength := 0
	for i := 0; i < len(arr); i++ {
		resultSliceLength += len(arr[i])
	}
	newSlice := make(Arr, resultSliceLength)
	index := 0
	for _, s := range arr {
		for i := 0; i < len(s); i++ {
			newSlice[index] = s[i]
			index += 1
		}
	}
	return newSlice
}

func ImplodeWithoutCalc[Arr ~[]T, T any](arr ...Arr) Arr {

	newSlice := make(Arr, 0)
	for _, s := range arr {
		newSlice = append(newSlice, s...)
	}
	return newSlice
}
