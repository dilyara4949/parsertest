//go:build unit || all
// +build unit all

package assets

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

type TestCase struct {
	name    string
	input   []byte
	want    []Product
	wantErr error
}

func InitData() []TestCase {
	tests := make([]TestCase, 0)

	tests = append(tests, TestCase{"nil array", nil, nil, errEmptyInput})

	products := []Product{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}
	data, _ := json.Marshal(products)

	tests = append(tests, TestCase{"empty array", data, nil, errManyElements})

	products = []Product{}
	data, _ = json.Marshal(products)

	tests = append(tests, TestCase{"valid", []byte(`[]`), products, nil})
	return tests
}

func TestParse(t *testing.T) {
	tests := InitData()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
