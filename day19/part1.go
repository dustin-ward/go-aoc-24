package day19

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part1(inpt io.Reader) string {
	reader := bufio.NewReader(inpt)
    line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]

    towels := make(map[string]struct{})
    for _, str := range strings.Split(line, ",") {
        towels[strings.Trim(str, " ")] = struct{}{}
    }
    line, _ = reader.ReadString('\n')

    var f func(str string)bool
    f = func(str string)bool {
        if str == "" {
            return true
        }

        canDo := false
        for k, _ := range towels {
            if strings.HasPrefix(str, k) {
                canDo = canDo || f(str[len(k):])
            }
        }

        return canDo
    }

    var ans int
	for {
        line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            return strconv.Itoa(ans)
		}
        line = line[:len(line)-1]

        if f(line) {
            ans++
        }
	}
}
