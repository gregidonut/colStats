package main

func sum(data []float64) float64 {
	var sum float64
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// statsFunc is an auxiliary type using the same signature
// as the above sum() and avg() this will prove useful when
// this type is used as an input parameter on a calling
// to better manage testing and make the calling function more
// concise
type statsFunc func(data []float64) float64
