package reader

import (
	"bufio"
	"os"
)

func Reader(f string) *bufio.Scanner {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}
