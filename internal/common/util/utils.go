package util

import "sort"

// Median return the median of number list
func Median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	length := len(dataCopy)
	mNum := length / 2
	if length%2 == 0 {
		return (dataCopy[mNum-1] + dataCopy[mNum]) / 2
	}
	return dataCopy[mNum]
}
