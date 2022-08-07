package main

import "encoding/json"

func decodeJson(jsonBytes []byte) (Book, error) {

	var book Book
	e := json.Unmarshal(jsonBytes, &book)

	if e != nil {
		return book, e
	}
	return book, nil
}
