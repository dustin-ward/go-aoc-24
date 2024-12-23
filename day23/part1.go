package day23

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Part1(inpt io.Reader) string {
    adj := make(map[string][]string)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]

        s := strings.Split(line, "-")
        c1 := s[0]
        c2 := s[1]

        if _, ok := adj[c1]; !ok {
            adj[c1] = make([]string,0)
        }
        if _, ok := adj[c2]; !ok {
            adj[c2] = make([]string,0)
        }

        adj[c1] = append(adj[c1], c2)
        adj[c2] = append(adj[c2], c1)
	}

    ans := make(map[string]struct{})
    for i,v := range adj {
        if i[0] != 't' {
            continue
        }

        for _, j := range v {
            if j == i {
                continue
            }
            for _, k := range adj[j] {
                if k == i || k == j {
                    continue
                }
                for _, x := range adj[k] {
                    if x == i {
                        c := []string{i,j,k}
                        slices.Sort(c)
                        ans[fmt.Sprintf("%s,%s,%s", c[0],c[1],c[2])] = struct{}{}
                        break
                    }
                }
            }
        }
    }

    return strconv.Itoa(len(ans))
}
