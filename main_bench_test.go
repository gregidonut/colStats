package main

import (
	"io"
	"path/filepath"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	tests := []struct {
		operation string
	}{
		{
			operation: "avg",
		},
		{
			operation: "max",
		},
		{
			operation: "min",
		},
	}

	for _, tt := range tests {
		b.Run(tt.operation, func(b *testing.B) {
			fileNames, err := filepath.Glob("./testdata/benchmark/*csv")
			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := run(fileNames, config{op: tt.operation, column: 2}, io.Discard); err != nil {
					b.Error(err)
				}
			}
		})
	}
}
