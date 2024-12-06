package day6

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part1(inpt io.Reader) string {
    board := make([]string,0)

    type pos struct {
        x int
        y int
    }
    var guard pos

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
            guard.x = i
            guard.y = len(board)
        }

        board = append(board, line)
        temp := make([]bool,len(line))
        visited = append(visited,temp)
	}
    N := len(board)
    M := len(board[0])

    curDir := 0
    for (guard.x >= 0 && guard.x < M && guard.y >= 0 && guard.y < N) {
        visited[guard.y][guard.x] = true
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
    }

    ans := 0
    for _, v := range visited {
        for _, v2 := range v {
            if v2 {
                ans++
            }
        }
    }
    return strconv.Itoa(ans)
}
