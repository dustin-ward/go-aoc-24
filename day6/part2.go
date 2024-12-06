package day6

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
    board := make([][]byte,0)

    type pos struct {
        x int
        y int
    }
    var startGuard pos

    var dy = []int{-1,0,1,0}
    var dx = []int{0,1,0,-1}

    visited := make([][]bool,0)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        if i := strings.Index(line, "^"); i != -1 {
            startGuard.x = i
            startGuard.y = len(board)
        }

        board = append(board, []byte(line))
        temp := make([]bool,len(line))
        visited = append(visited,temp)
	}
    N := len(board)
    M := len(board[0])

    simulate := func(maxIter int, init bool)bool {
        guard := startGuard
        iters := 0
        curDir := 0
        for (guard.x >= 0 && guard.x < M && guard.y >= 0 && guard.y < N) {
            if !init && iters >= maxIter {
                return true
            }
            if init {
                visited[guard.y][guard.x] = true
            }
            nextX := guard.x + dx[curDir]
            nextY := guard.y + dy[curDir]

            if nextX<0 || nextX>=M || nextY<0 || nextY>=N {
                break
            }

            if board[nextY][nextX] == '#' {
                curDir = (curDir + 1) % 4
            } else {
                guard.x = nextX
                guard.y = nextY
            }
            iters++
        }
        return false
    }
    simulate(0, true)

    ans := 0
    MAX_ITERS := 10000
    for i, v := range visited {
        for j, v2 := range v {
            if v2 {
                board[i][j] = '#'
                inf := simulate(MAX_ITERS,false)
                if inf {
                    ans++
                }
                board[i][j] = '.'
            }
        }
    }

    return strconv.Itoa(ans)
}
