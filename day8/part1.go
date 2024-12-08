package day8

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(inpt io.Reader) string {
    var board [][]byte

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

    type coord struct {
        x int
        y int
    }

    antennas := make(map[byte][]coord)
    for i, v := range board {
        for j, v2 := range v {
            if v2 != byte('.') {
                antennas[v2] = append(antennas[v2], coord{j,i})
            }
        }
    }

    antinodes := make([][]bool, N)
    for i := range N {
        antinodes[i] = make([]bool, M)
    }

    bounds := func(i,j int)bool {
        return i>=0 && i<N && j>=0 && j<M
    }

    var ans int
    for _, coords := range antennas {
        for i, c1 := range coords {
            for j, c2 := range coords {
                if i == j {
                    continue
                }

                dx := c2.x-c1.x
                dy := c2.y-c1.y
                i2 := c2.y+dy
                j2 := c2.x+dx
                if bounds(i2,j2) {
                    if !antinodes[i2][j2] {
                        ans++
                    }
                    antinodes[i2][j2] = true
                }
            }
        }
    }

    return strconv.Itoa(ans)
}
