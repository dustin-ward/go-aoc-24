package day11

import (
	"bufio"
	"io"
)

func Part1(inpt io.Reader) string {
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
			return "Not implemented"
		}
		line = line[:len(line)-1]
    }
}
