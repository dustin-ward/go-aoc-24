package day11

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
    var ans int
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            return strconv.Itoa(ans)
		}
		line = line[:len(line)-1]

        memo := make(map[int]map[int]int)
        var f func(int,int)int
        f = func(x, d int)int {
            if d == 0 {
                return 1
            }

            if r, ok := memo[x][d]; ok {
                return r
            }

            s := strconv.Itoa(x)
            tempAns := 0
            if x == 0 {
                tempAns += f(1, d-1)
            } else if len(s)%2==0 {
                l, _ := strconv.Atoi(s[:len(s)/2])   
                r, _ := strconv.Atoi(s[len(s)/2:])   
                tempAns += f(l, d-1)
                tempAns += f(r, d-1)
            } else {
                tempAns += f(x*2024, d-1)
            }

            if memo[x] == nil {
                memo[x] = make(map[int]int)
            }
            memo[x][d] = tempAns
            return tempAns
        }

        rocks := strings.Split(line, " ")
        for _, rock := range rocks {
            rInt, _ := strconv.Atoi(rock)
            ans += f(rInt, 75)
        }
	}
}
