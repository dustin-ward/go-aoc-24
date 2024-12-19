package day19

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
	reader := bufio.NewReader(inpt)
    line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]

    towels := make(map[string]struct{})
    for _, str := range strings.Split(line, ",") {
        towels[strings.Trim(str, " ")] = struct{}{}
    }
    line, _ = reader.ReadString('\n')

    memo := make(map[string]int)
    var f func(str string)int
    f = func(str string)int {
        if str == "" {
            return 1
        }

        if v, ok := memo[str]; ok {
            return v
        }

        ans := 0
        for k, _ := range towels {
            if strings.HasPrefix(str, k) {
                ans +=f(str[len(k):])
            }
        }

        memo[str] = ans
        return ans
    }

    var ans int
	for {
        line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            return strconv.Itoa(ans)
		}
        line = line[:len(line)-1]

        ans += f(line)
	}
}
