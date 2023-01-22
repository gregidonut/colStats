package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		fileNames []string
		cfg       config
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr error
	}{
		{
			name: "runAvg1File",
			args: args{
				fileNames: []string{"./testdata/example.csv"},
				cfg: config{
					op:     "avg",
					column: 3,
				},
			},
			wantOut: "227.6\n",
			wantErr: nil,
		},
		{
			name: "runAvgMultipleFiles",
			args: args{
				fileNames: []string{"./testdata/example.csv", "./testdata/example2.csv"},
				cfg: config{
					op:     "avg",
					column: 3,
				},
			},
			wantOut: "233.84\n",
			wantErr: nil,
		},
		{
			name: "runFailRead",
			args: args{
				fileNames: []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
				cfg: config{
					op:     "avg",
					column: 2,
				},
			},
			wantOut: "",
			wantErr: os.ErrNotExist,
		},
		{
			name: "runFailColumn",
			args: args{
				fileNames: []string{"./testdata/example.csv"},
				cfg: config{
					op:     "avg",
					column: 0,
				},
			},
			wantOut: "",
			wantErr: ErrInvalidColumn,
		},
		{
			name: "runFailNoFiles",
			args: args{
				fileNames: []string{},
				cfg: config{
					op:     "avg",
					column: 2,
				},
			},
			wantOut: "",
			wantErr: ErrNoFiles,
		},
		{
			name: "runFailOperation",
			args: args{
				fileNames: []string{"./testdata/example.csv"},
				cfg: config{
					op:     "invalid",
					column: 3,
				},
			},
			wantOut: "",
			wantErr: ErrInvalidOperation,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := run(tt.args.fileNames, tt.args.cfg, out)
			if tt.wantErr != nil {
				if err == nil {
					t.Error("wanted error but didn't get one")
				}
				if err != tt.wantErr {
					t.Errorf("run() error = %q, wantErr %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}
			if out.String() != tt.wantOut {
				t.Errorf("run() gotOut = %v, want %v", out.String(), tt.wantOut)
			}
		})
	}
}
