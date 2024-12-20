package day20

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	complibgo "github.com/dustin-ward/comp-lib-go"
)

func Part1(inpt io.Reader) string {
    board := make([][]byte,0)

    type xy struct {
        y int
        x int
    }
    var start,end xy
    _ = start

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]

        if v := strings.Index(line, "S"); v != -1 {
            start = xy{len(board),v}
            line = line[:v] + "." + line[v+1:]
        }
        if v := strings.Index(line, "E"); v != -1 {
            end = xy{len(board),v}
            line = line[:v] + "." + line[v+1:]
        }
        board = append(board, []byte(line))
	}

    N := len(board)
    M := len(board[0])

    var dy = []int{0,-1,0,1}
    var dx = []int{-1,0,1,0}

    bounds := func(i,j int)bool {
        return i>=0&&i<N&&j>=0&&j<M
    }

    getIdx := func(i,j int)int {
        return (i*M) + j
    }

    G := complibgo.NewGraphLst(N*M)
    for i, arr := range board {
        for j, ch := range arr {
            if ch == '#' {
                continue
            }
            for k := 0; k < 4; k++ {
                i2 := i+dy[k]
                j2 := j+dx[k]
                if bounds(i2,j2) && board[i2][j2] != '#' {
                    G.AddEdge(getIdx(i,j), getIdx(i2,j2), 1)
                }
            }
        }
    }
    D, _ := complibgo.Dijkstra_Sparse(G, getIdx(end.y, end.x))

    var dy2 = []int{-2,-1,-1,0,0,1,1,2}
    var dx2 = []int{0,-1,1,-2,2,-1,1,0}

    ans := make(map[int]int)
    for i, arr := range board {
        for j, ch := range arr {
            if ch == '#' {
                continue
            }

            for k := 0; k < 8; k++ {
                i2 := i+dy2[k]
                j2 := j+dx2[k]

                if bounds(i2,j2) && board[i2][j2] != '#' {
                    save := (D[getIdx(i,j)]-D[getIdx(i2,j2)])-2
                    if save > 0 {
                        ans[save]++
                    }
                }
            }
        }
    }

    var actualAns int
    for k,v := range ans {
        if k >= 100 {
            actualAns += v
        }
    }

    return strconv.Itoa(actualAns)
}
