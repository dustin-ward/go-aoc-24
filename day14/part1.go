package day14

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Part1(inpt io.Reader) string {
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

    Q := make([]int, 4)
    for _, r := range robots {
        for _ = range 100 {
            r.p.x = ((r.p.x + r.v.x) + M) % M
            r.p.y = ((r.p.y + r.v.y) + N) % N
        }

        if r.p.x < M/2 {
            if r.p.y < N/2 {
                Q[0]++
            } else if r.p.y > N/2 {
                Q[1]++
            }
        } else if r.p.x > M/2{
            if r.p.y < N/2 {
                Q[2]++
            } else if r.p.y > N/2 {
                Q[3]++
            }
        }
    }

    ans := Q[0] * Q[1] * Q[2] * Q[3]
    return strconv.Itoa(ans)
}
