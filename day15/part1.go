package day15

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part1(inpt io.Reader) string {
    B := make([][]byte,0)

    type xy struct {
        x int
        y int
    }
    var robot xy

    instrs := false
    inst := ""
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]

        if !instrs {
            if line == "" {
                instrs = true
                continue
            }
            idx := strings.IndexByte(line, byte('@'))
            if idx != -1 {
                robot = xy{idx,len(B)}
            }
            B = append(B, []byte(line))
        } else {
            inst += line
        }
	}

    var dx = []int{-1,0,1,0}
    var dy = []int{0,-1,0,1}

    var move func(i,j,d int)bool
    move = func(i,j,d int)bool {
        if B[i][j] == '#' {
            return false
        }
        if B[i][j] == '.' {
            return true
        }

        i2 := i+dy[d]
        j2 := j+dx[d]
        if move(i2, j2, d) {
            B[i2][j2] = B[i][j]
            B[i][j] = '.'
            return true
        }
        return false
    }

    for _, ch := range inst {
        switch (ch) {
        case '<':
            if move(robot.y, robot.x, 0) {
                robot.x--
            }
        case '^':
            if move(robot.y, robot.x, 1) {
                robot.y--
            }
        case '>':
            if move(robot.y, robot.x, 2) {
                robot.x++
            }
        case 'v':
            if move(robot.y, robot.x, 3) {
                robot.y++
            }
        }
    }

    var ans int
    for i, v := range B {
        for j, ch := range v {
            if ch == byte('O') {
                ans += (100*i)+j
            }
        }
    }

    return strconv.Itoa(ans)
}
