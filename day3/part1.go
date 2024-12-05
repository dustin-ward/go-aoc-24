package day3

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func Part1(inpt io.Reader) string {
	var ans = 0
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
			return strconv.Itoa(ans)
		}
		line = line[:len(line)-1]

		for _, match := range r.FindAllStringSubmatch(line, -1) {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			ans += a * b
		}
	}

}
