package day12

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
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

    areas := make(map[xy]map[xy]struct{})
    visited := make([][]bool,N)
    for i, _ := range visited {
        visited[i] = make([]bool,M)
    }
    for i := 1; i<N-1; i++ {
        for j := 1; j<M-1; j++ {
            start := xy{i,j}
            if visited[i][j] {
                continue
            }
            areas[start] = make(map[xy]struct{})
            var queue = []xy{xy{i,j}}
            visited[i][j] = true

            for len(queue) > 0 {
                front := queue[0]
                queue = queue[1:]

                ch := garden[front.y][front.x]
                areas[start][front] = struct{}{}

                for k := 0; k < 4; k++ {
                    i2 := front.y+dy[k]
                    j2 := front.x+dx[k]
                    ch2 := garden[i2][j2]
                    
                    if !visited[i2][j2] && ch2 == ch {
                        // fmt.Println("adding")
                        queue = append(queue, xy{i2,j2})
                        visited[i2][j2] = true
                    }
                }
            }
        }
    }

    shapewalls := make(map[xy]int)
    for id, shape := range areas {

        // TOP WALLS
        prevLine := make(map[xy]struct{})
        for i := 0; i < N; i++ {
            nextSet := make(map[xy]struct{})
            inWall := false
            for j := 0; j < M; j++ {
                cur := xy{i,j}
                // Top wall
                if _, ok := shape[cur]; ok {
                    nextSet[cur] = struct{}{}
                    if _, ok2 := prevLine[xy{i-1,j}]; !ok2 && !inWall {
                        inWall = true
                        shapewalls[id]++
                    } else if ok2 {
                        inWall = false
                    }
                } else {
                    inWall = false
                }
            }
            prevLine = nextSet
        }

        // BOTTOM WALLS
        prevLine = make(map[xy]struct{})
        for i := N-1; i >= 0; i-- {
            nextSet := make(map[xy]struct{})
            inWall := false
            for j := 0; j < M; j++ {
                cur := xy{i,j}
                // Top wall
                if _, ok := shape[cur]; ok {
                    nextSet[cur] = struct{}{}
                    if _, ok2 := prevLine[xy{i+1,j}]; !ok2 && !inWall {
                        inWall = true
                        shapewalls[id]++
                    } else if ok2 {
                        inWall = false
                    }
                } else {
                    inWall = false
                }
            }
            prevLine = nextSet
        }

        // LEFT WALLS
        prevLine = make(map[xy]struct{})
        for j := 0; j < M; j++ {
            nextSet := make(map[xy]struct{})
            inWall := false
            for i := 0; i < N; i++ {
                cur := xy{i,j}
                // Top wall
                if _, ok := shape[cur]; ok {
                    nextSet[cur] = struct{}{}
                    if _, ok2 := prevLine[xy{i,j-1}]; !ok2 && !inWall {
                        inWall = true
                        shapewalls[id]++
                    } else if ok2 {
                        inWall = false
                    }
                } else {
                    inWall = false
                }
            }
            prevLine = nextSet
        }

        // RIGHT WALLS
        prevLine = make(map[xy]struct{})
        for j := M-1; j >= 0; j-- {
            nextSet := make(map[xy]struct{})
            inWall := false
            for i := 0; i < N; i++ {
                cur := xy{i,j}
                // Top wall
                if _, ok := shape[cur]; ok {
                    nextSet[cur] = struct{}{}
                    if _, ok2 := prevLine[xy{i,j+1}]; !ok2 && !inWall {
                        inWall = true
                        shapewalls[id]++
                    } else if ok2 {
                        inWall = false
                    }
                } else {
                    inWall = false
                }
            }
            prevLine = nextSet
        }
    }

    var ans int
    for id, shape := range areas {
        ans += len(shape) * shapewalls[id]
    }
    return strconv.Itoa(ans)
}
