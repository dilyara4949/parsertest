//go:build unit || all
// +build unit all

package assets

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name    string
	input   []byte
	want    []Product
	wantErr error
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    []Product
		wantErr bool
		err     error
	}{
		{
			name:    "empty input",
			input:   []byte{},
			wantErr: true,
		},
		{
			name:    "nil input",
			input:   nil,
			wantErr: true,
		},
		{
			name:  "ok",
			input: []byte(`[{}]`),
			want:  []Product{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
