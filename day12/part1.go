package day12

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part1(inpt io.Reader) string {
    garden := make([][]byte,0)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        garden = append(garden, []byte("."+line+"."))
	}
    buf := []byte(strings.Repeat(".",len(garden[0])))
    garden = append(garden, buf)
    garden = append([][]byte{buf}, garden...)
    N := len(garden)
    M := len(garden[0])
    
    var dx = []int{-1,0,1,0}
    var dy = []int{0,-1,0,1}
    type xy struct {
        y int
        x int
    }
    type stats struct {
        area int
        perim int
    }

    areas := make(map[xy]stats)
    visited := make([][]bool,N)
    for i := range visited {
        visited[i] = make([]bool,M)
    }
    for i := 1; i<N-1; i++ {
        for j := 1; j<M-1; j++ {
            start := xy{i,j}
            if visited[i][j] {
                continue
            }
            var queue = []xy{{i,j}}
            visited[i][j] = true

            for len(queue) > 0 {
                front := queue[0]
                queue = queue[1:]

                ch := garden[front.y][front.x]

                newPerim := 0
                for k := 0; k < 4; k++ {
                    i2 := front.y+dy[k]
                    j2 := front.x+dx[k]
                    ch2 := garden[i2][j2]
                    
                    if !visited[i2][j2] && ch2 == ch {
                        queue = append(queue, xy{i2,j2})
                        visited[i2][j2] = true
                    }
                    if ch2 != ch {
                        newPerim++
                    }
                }

                areas[start] = stats{
                    areas[start].area +1,
                    areas[start].perim + newPerim,
                }
            }
        }
    }

    var ans int
    for _, stat := range areas {
        ans += stat.area * stat.perim
    }

    return strconv.Itoa(ans)
}
