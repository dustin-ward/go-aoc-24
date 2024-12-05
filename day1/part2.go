package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Part2(inpt io.Reader) string {
	L := make([]int, 0)
	R := make(map[int]int)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
			break
		}
		line = line[:len(line)-1]

		var a, b int
		fmt.Sscanf(line, "%d %d", &a, &b)

		L = append(L, a)
		R[b]++
	}

	ans := 0

	for _, x := range L {
		ans += x * R[x]
	}

	return strconv.Itoa(ans)
}
