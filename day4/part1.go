package day4

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(inpt io.Reader) string {
    var board []string

    var dx = []int{-1,-1,0,1,1,1,0,-1}
    var dy = []int{0,-1,-1,-1,0,1,1,1}

    var n,m int

    bounds := func (i,j int) bool {
        return i>=0 && i<n && j>=0 && j<m
    }

    const XMAS = "XMAS"

    reader := bufio.NewReader(inpt)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            //EOF
            break
        }
        line = line[:len(line)-1]
        board = append(board, line)
    }

    n = len(board)
    m = len(board[0])

    ans := 0
    for i := range n {
        for j := range m {
            for k := range 8 {
                for cnt, ch := range "XMAS" {
                    i2 := i+(dy[k]*cnt)
                    j2 := j+(dx[k]*cnt)
                    if !bounds(i2,j2) || board[i2][j2] != byte(ch) {
                        goto invalid
                    }
                }
                ans++
invalid:
            }
        }
    }

    return strconv.Itoa(ans)
}
