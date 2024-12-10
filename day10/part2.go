package day10

import (
	"bufio"
	"io"
	"strconv"
)

func Part2(inpt io.Reader) string {
    board := make([][]byte,0)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        board = append(board, []byte(line))
	}
    N := len(board)
    M := len(board[0])
    
    getId := func(i, j int)int {
        return (i*M)+j
    }

    bounds := func(i,j int)bool {
        return i>=0&&i<N&&j>=0&&j<M
    }

    adj := make([][]int,N*M)

    var dx = []int{-1,0,1,0}
    var dy = []int{0,-1,0,1}

    type coord struct {
        y int
        x int
    }
    heads := make([]coord,0)

    getPos := func(x int)coord {
        return coord{x/N, x%N}
    }


    for i, v := range board {
        for j, _ := range v {
            if board[i][j] == byte('0') {
                heads = append(heads, coord{i,j})
            }

            for k:=0; k<4; k++ {
                i2 := i+dy[k]
                j2 := j+dx[k]
                if bounds(i2,j2) && board[i2][j2] == board[i][j]+1 {
                    adj[getId(i,j)] = append(adj[getId(i, j)], getId(i2,j2))
                }
            }
        }
    }

    var ans int
    for _, head := range heads {
        visited := make([]bool, N*M)
        var stack = []int{getId(head.y,head.x)}

        for len(stack) > 0 {
            node := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            visited[node] = true

            crd := getPos(node)
            if board[crd.y][crd.x] == byte('9') {
                ans++
            }

            for _, nbr := range adj[node] {
                stack = append(stack, nbr)
            }
        }
    }

    return strconv.Itoa(ans)
}
