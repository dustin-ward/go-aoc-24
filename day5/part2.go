package day5

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
	var ans int
	rules := true

	prereqs := make(map[int]map[int]struct{})
	rulesList := make([]string, 0)

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
			var a, b int
			fmt.Sscanf(line, "%d|%d", &a, &b)
			if prereqs[a] == nil {
				prereqs[a] = make(map[int]struct{})
			}
			if prereqs[b] == nil {
				prereqs[b] = make(map[int]struct{})
			}
			prereqs[a][b] = struct{}{}
		} else {
			rulesList = append(rulesList, line)
		}
	}

    for _, rule := range rulesList {
        s := strings.Split(rule, ",")
        orig := make([]int,len(s))
        sorted := make([]int,len(s))

        for i, str := range s {
            x, _ := strconv.Atoi(str)
            orig[i] = x
            sorted[i] = x
        }

        slices.SortStableFunc(sorted, func(i,j int) int {
            if _, ok := prereqs[i][j]; ok {
                return -1
            } else if _, ok := prereqs[j][i]; ok {
                return 1
            }
            return 0
        })

        for i, _ := range orig {
            if orig[i] != sorted[i] {
                goto good
            }
        }
        continue
good:
        ans += sorted[len(sorted)/2]
    }

	return strconv.Itoa(ans)
}
