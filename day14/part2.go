package day14

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
    type xy struct {
        x int
        y int
    }
    type robot struct {
        p xy
        v xy
    }
    robots := make([]robot,0)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        var r robot
        fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.p.x, &r.p.y, &r.v.x, &r.v.y)
        robots = append(robots, r)
	}

    N := 103
    M := 101

    for iter := 0; iter < 10000; iter++ {
        B := make([][]rune,N)
        for i := range N {
            B[i] = make([]rune,M)
        }
        for i, _ := range robots {
            robots[i].p.x = ((robots[i].p.x + robots[i].v.x) + M) % M
            robots[i].p.y = ((robots[i].p.y + robots[i].v.y) + N) % N
            B[robots[i].p.y][robots[i].p.x] = 'X'
        }

        for _, r := range B {
            if strings.Contains(string(r), "XXXXXXXXXXXXXXXXXXXX") {
                return strconv.Itoa(iter+1)
            }
        }
    }

    return strconv.Itoa(0)
}
