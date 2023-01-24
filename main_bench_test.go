package main

import (
	"io"
	"path/filepath"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	fileNames, err := filepath.Glob("./testdata/benchmark/*csv")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := run(fileNames, config{op: "avg", column: 2}, io.Discard); err != nil {
			b.Error(err)
		}
	}
}
