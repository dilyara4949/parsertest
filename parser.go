package assets

import (
	"encoding/json"
	"errors"
)

type Product struct {
	Name   string   `json:"name"`
	Market string   `json:"market"`
	Tags   []string `json:"tags"`
}

var (
	errEmptyInput   = errors.New("empty input")
	errManyElements = errors.New("input should contain 10 or less elements")
)

/*
	Parse consumes a JSON with an array of products.
	Example:
	[
	   {
		  "name":"test name",
		  "market":"USA",
		  "tags":[
			 "t1",
			 "t2",
		  ]
	   }
	]
*/

func Parse(input []byte) ([]Product, error) {
	if input == nil || len(input) == 0 {
		return nil, errEmptyInput
	}

	if len(input) > 10 {
		return nil, errManyElements
	}

	var p []Product

	if err := json.Unmarshal(input, &p); err != nil {
		return nil, err
	}

	return p, nil
}
