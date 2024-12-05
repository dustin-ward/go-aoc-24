package day1

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"

	"github.com/dustin-ward/go-aoc-24/util"
)

func Part1(inpt io.Reader) string {
	L := make([]int, 0)
	R := make([]int, 0)

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
		R = append(R, b)
	}

	slices.Sort(L)
	slices.Sort(R)

	ans := 0

	for i := range len(L) {
		ans += util.Abs(L[i] - R[i])
	}

	return strconv.Itoa(ans)
}
