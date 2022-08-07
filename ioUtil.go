package main

import "os"

func readFile(file string) ([]byte, error) {
	b, e := os.ReadFile(file)
	if e != nil {
		return nil, e
	}
	return b, nil
}
