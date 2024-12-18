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

    board := make([][]byte,H)
    for i := range H {
        board[i] = []byte(strings.Repeat(".", W))
    }

    type xy struct {
        x int
        y int
    }
    bytes := make([]xy,0)

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
        board[y][x] = '#'
	}

    UF := complibgo.NewUnionFind(W*H)

    var dy = []int{0,-1,0,1}
    var dx = []int{-1,0,1,0}

    bounds := func(i,j int)bool {
        return i>=0&&i<H&&j>=0&&j<W
    }

    getIdx := func(i,j int)int {
        return (i*W)+j
    }

    for i, arr := range board {
        for j := range arr {
            if board[i][j] == '#' {
                continue
            }
            for k := 0; k<4; k++ {
                i2 := i+dy[k]
                j2 := j+dx[k]
                if bounds(i2,j2) && board[i2][j2] != '#' {
                    UF.Merge(getIdx(i,j), getIdx(i2,j2))
                }
            }
        }
    }

    byteIdx := len(bytes)
    for byteIdx >= 0 && UF.Find(getIdx(0,0)) != UF.Find(getIdx(H-1,W-1)) {
        byteIdx--
        
        x,y := bytes[byteIdx].x, bytes[byteIdx].y
        board[y][x] = '.'

        for k := 0; k < 4; k++ {
            x2 := x+dx[k]
            y2 := y+dy[k]
            if bounds(x2,y2) && board[y2][x2] != '#' {
                UF.Merge(getIdx(y,x), getIdx(y2,x2))
            }
        }
    }

    ans := bytes[byteIdx]
    return fmt.Sprintf("%d,%d", ans.x, ans.y)
}
