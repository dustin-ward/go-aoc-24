package day5

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part1(inpt io.Reader) string {
    var ans int
    rules := true

    prereqs := make(map[int]map[int]bool)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]

        if line == "" {
            rules = false
            continue
        }

        if rules {
            var a,b int
            fmt.Sscanf(line, "%d|%d", &a, &b)
            if prereqs[a] == nil {
                prereqs[a] = make(map[int]bool)
            }
            prereqs[a][b] = true
        } else {
            var mid int
            curPages := make(map[int]bool)
            pages := strings.Split(line, ",")
            for _, page := range pages {
                pageInt, _ := strconv.Atoi(page)

                for c, _ := range curPages {
                    if _, ok := prereqs[pageInt][c]; ok {
                        goto invalid
                    }
                }
                
                curPages[pageInt] = true
            }
            mid, _ = strconv.Atoi(pages[len(pages)/2])
            ans += mid
invalid:
        }
	}

    return strconv.Itoa(ans)
}
