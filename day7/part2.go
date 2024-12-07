package day7

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
    var operatorFuncs = []func(a,b int)int{
        func(a,b int) int {
            return a+b
        },
        func(a,b int) int {
            return a*b
        },
        func(a,b int) int {
            aStr := strconv.Itoa(a)
            bStr := strconv.Itoa(b)
            aStr += bStr
            temp, _ := strconv.Atoi(aStr)
            return temp
        },
    }

    var f func(int, int, []string) bool
    f = func(target, sum int, args []string) bool {
        if len(args) == 0 {
            return target == sum
        }
        nextNum, _ := strconv.Atoi(args[0])

        possible := false
        for _, op := range operatorFuncs {
            possible = possible || f(target, op(sum, nextNum), args[1:])
        }
        return possible
    }

    var ans int

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
			return strconv.Itoa(ans)
		}
		line = line[:len(line)-1]

        s := strings.Split(line, ":")
        testVal, _ := strconv.Atoi(s[0])
        args := strings.Split(s[1][1:], " ")

        if f(testVal,0,args) {
            ans += testVal
        }
	}
}
