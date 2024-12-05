package day4

import (
	"bufio"
	"io"
	"strconv"
)

func Part2(inpt io.Reader) string {
    var board []string

    var n,m int

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
            var chUL,chUR,chBR,chBL byte
            var d1,d2 bool
            if board[i][j] != 'A' || i == 0 || i == n-1 || j == 0 || j == m-1 {
                goto invalid
            }

            chUL = board[i-1][j-1]
            chBR = board[i+1][j+1]
            chUR = board[i-1][j+1]
            chBL = board[i+1][j-1]

            d1 = (chUL == 'M' && chBR == 'S') || (chUL == 'S' && chBR == 'M')
            d2 = (chUR == 'M' && chBL == 'S') || (chUR == 'S' && chBL == 'M')

            if !(d1 && d2) {
                goto invalid
            }

            ans++
invalid:
        }
    }

    return strconv.Itoa(ans)
}
