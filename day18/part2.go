package day18

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	complibgo "github.com/dustin-ward/comp-lib-go"
)

func Part2(inpt io.Reader) string {
    var W int = 71
    var H int = 71

    var BYTES int = 1024

    board := make([][]byte,H)
    for i := range H {
        board[i] = []byte(strings.Repeat(".", W))
    }

    type xy struct {
        x int
        y int
    }
    bytes := make([]xy,0)

    linesRead := 0
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        var x,y int
        fmt.Sscanf(line, "%d,%d", &x, &y)

        bytes = append(bytes, xy{x,y})

        if linesRead == BYTES {
            continue
        }
        board[y][x] = '#'
        linesRead++
	}

    var dy = []int{0,-1,0,1}
    var dx = []int{-1,0,1,0}
    
    bounds := func(i,j int)bool {
        return i>=0&&i<H&&j>=0&&j<W
    }

    getIdx := func(i,j int)int {
        return (i*W)+j
    }

    var ans xy
    for idx := BYTES; idx < len(bytes); idx ++ {
        byte := bytes[idx]
        board[byte.y][byte.x] = '#'

        G := complibgo.NewGraphLst(W*H)
        for i, arr := range board {
            for j, ch := range arr {
                if ch == '#' {
                    continue
                }

                for k := 0; k<4; k++ {
                    i2 := i+dy[k]
                    j2 := j+dx[k]
                    if bounds(i2,j2) && board[i2][j2] != '#' {
                        G.AddEdge(getIdx(i,j), getIdx(i2,j2), 1)
                    }
                }
            }
        }

        D, _ := complibgo.Dijkstra_Sparse(G, 0)
        if D[(W*H)-1] == complibgo.INF {
            ans = byte
            break
        }
    }

    return fmt.Sprintf("%d,%d", ans.x, ans.y)
}
