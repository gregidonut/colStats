package main

import (
	"fmt"
	"testing"
)

func TestOperations(t *testing.T) {
	data := [][]float64{
		{10, 20, 15, 30, 45, 50, 100, 30},
		{5.5, 8, 2.2, 9.75, 8.45, 3, 2.5, 10.25, 4.75, 6.1, 7.67, 12.287, 5.47},
		{-10, -20},
		{102, 37, 44, 57, 67, 129},
	}
	tests := []struct {
		name string
		op   statsFunc
		want []float64
	}{
		{
			name: "Sum",
			op:   sum,
			want: []float64{300, 85.927, -30, 436},
		},
		{
			name: "Avg",
			op:   avg,
			want: []float64{37.5, 6.609769230769231, -15, 72.666666666666666},
		},
	}

	// for each test case and for each float in the
	// test's want map the data's index to the float
	// in the want's index as the expected output for
	// running the data slice
	for _, tt := range tests {
		for k, want := range tt.want {
			name := fmt.Sprintf("%sData%d", tt.name, k)
			t.Run(name, func(t *testing.T) {
				got := tt.op(data[k])
				if want != got {
					t.Errorf("want %g, got %g instead", want, got)
				}
			})
		}
	}
}
