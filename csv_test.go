package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
	"testing/iotest"
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

func TestCSVFloat(t *testing.T) {
	csvData := `IP Address,Requests,Response Time
192.168.0.199,2056,236
192.168.0.88,899,220
192.168.0.199,3054,226
192.168.0.100,4133,218
192.168.0.199,950,238
`
	tests := []struct {
		name    string
		col     int
		want    []float64
		wantErr error
		r       io.Reader
	}{
		{
			name:    "Column2",
			col:     2,
			want:    []float64{2056, 899, 3054, 4133, 950},
			wantErr: nil,
			r:       bytes.NewBufferString(csvData),
		},
		{
			name:    "Column3",
			col:     3,
			want:    []float64{236, 220, 226, 218, 238},
			wantErr: nil,
			r:       bytes.NewBufferString(csvData),
		},
		{
			name:    "FailRead",
			col:     1,
			want:    nil,
			wantErr: iotest.ErrTimeout,
			r:       iotest.TimeoutReader(bytes.NewReader([]byte{0})),
		},
		{
			name:    "FailedNotNumber",
			col:     1,
			want:    nil,
			wantErr: ErrNotNumber,
			r:       bytes.NewBufferString(csvData),
		},
		{
			name:    "FailedInvalidColumn",
			col:     4,
			want:    nil,
			wantErr: ErrInvalidColumn,
			r:       bytes.NewBufferString(csvData),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := csv2Float(tt.r, tt.col)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error, but didn't get oone")
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %q, got %q", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}
			// this loop has the same idea from TestOperations() where it maps the
			// index of the float in tt.want to the index of the []float that was parsed
			// from the mock csv string(i.e. input:[]float to output:float
			for i, want := range tt.want {
				if got[i] != want {
					t.Errorf("want %g, got %g", want, got[i])
				}
			}
		})
	}
}
