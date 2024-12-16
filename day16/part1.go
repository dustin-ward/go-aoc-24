package day16

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/dustin-ward/comp-lib-go"
)

func Part1(inpt io.Reader) string {
    var startI, startJ int
    var endI, endJ int
    B := make([][]byte,0)
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        if x := strings.Index(line, "S"); x != -1 {
            startJ = x
            startI = len(B)
        }
        if x := strings.Index(line, "E"); x != -1 {
            endJ = x
            endI = len(B)
        }
        B = append(B, []byte(line))
	}
    N := len(B)
    M := len(B[0])

    var dx = []int{-1,0,1,0}
    var dy = []int{0,-1,0,1}

    getIdx := func(i,j,d int)int {
        return (i*M*4)+(j*4)+d
    }

    G := complibgo.NewGraphLst(N*M*4)

    for i, arr := range B {
        for j, ch := range arr {
            if ch == '#' {
                continue
            }
            for k := 0; k<4; k++ {
                i2 := i+dy[k]
                j2 := j+dx[k]
                u := getIdx(i,j,k)
                v := getIdx(i2,j2,k)
                if B[i2][j2] != '#' {
                    G.AddEdge(u,v,1)
                }

                l := getIdx(i,j,(k+1)%4)
                r := getIdx(i,j,(k+3)%4)
                G.AddEdge(u,l,1000)
                G.AddEdge(u,r,1000)
            }
        }
    }

    src := getIdx(startI, startJ, 2)
    D, _ := complibgo.Dijkstra_Sparse(G, src)

    ans := (1<<62)
    for k := 0; k<4; k++ {
        i := getIdx(endI, endJ, k)
        ans = min(ans, D[i])
    }

    return strconv.Itoa(ans)
}
