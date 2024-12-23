package day23

import (
	"bufio"
	"io"
	"slices"
	"sort"
	"strings"
)

func Part2(inpt io.Reader) string {
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

    cliques := make([]map[string]struct{},0)

    var bronKerbosch func(R,P,X map[string]struct{})
    bronKerbosch = func(R,P,X map[string]struct{}) {
        if len(P) == 0 && len(X) == 0 {
            cliques = append(cliques, R)
            return
        }

        for len(P) > 0 {
            var v string
            for k,_ := range P {
                v = k
                break
            }

            newR := make(map[string]struct{})
            for k,v := range R {
                newR[k] = v
            }
            newR[v] = struct{}{}
            newP := make(map[string]struct{})
            for k, _ := range P {
                if slices.Index(adj[v], k) != -1 {
                    newP[k] = struct{}{}
                }
            }
            newX := make(map[string]struct{})
            for k, _ := range X {
                if slices.Index(adj[v], k) != -1 {
                    newX[k] = struct{}{}
                }
            }

            bronKerbosch(newR,newP,newX)
            delete(P, v)
            X[v] = struct{}{}
        }
    }
    R := make(map[string]struct{})
    P := make(map[string]struct{})
    X := make(map[string]struct{})
    for k := range adj {
        P[k] = struct{}{}
    }
    bronKerbosch(R,P,X)

    sort.Slice(cliques, func(i,j int)bool{
        return len(cliques[i])>len(cliques[j])
    })
    c := make([]string,0)
    for k := range cliques[0] {
        c = append(c,k)
    }
    slices.Sort(c)
    ans := ""
    for _,k := range c {
        ans += k + ","
    }

    return ans[:len(ans)-1]
}
