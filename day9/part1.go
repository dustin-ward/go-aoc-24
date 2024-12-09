package day9

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(inpt io.Reader) string {
	reader := bufio.NewReader(inpt)
    var ans int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            return strconv.Itoa(ans)
		}
		line = line[:len(line)-1]

        fs := make([]int,0)
        id := 0
        for i := 0; i<len(line); i+=2 {
            file := int(line[i]-byte('0'))
            var space int
            if i == len(line)-1 {
                space = 0
            } else {
                space = int(line[i+1]-byte('0'))
            }
            for i := 0; i<file; i++ {
                fs = append(fs, id)
            }
            for i := 0; i<space; i++ {
                fs = append(fs, -1)
            }
            id++
        }

        N := len(fs)-1
        for i := 0; i<len(fs); i++ {
            if fs[i] != -1 {
                continue
            }
            
            for fs[N] == -1 {
                N--
            }
            if N <= i {
                break
            }

            fs[i] = fs[N]
            fs[N] = -1
            N--
        }

        for i, x := range fs {
            if x == -1 {
                continue
            }
            ans += i*x
        }
	}
}
