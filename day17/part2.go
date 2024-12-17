package day17

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
	var A, B, C int
	var programStr string

	reader := bufio.NewReader(inpt)

	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]
    fmt.Sscanf(line, "Register A: %d", &A)

	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	fmt.Sscanf(line, "Register B: %d", &B)

	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	fmt.Sscanf(line, "Register C: %d", &C)

	line, _ = reader.ReadString('\n')
	line, _ = reader.ReadString('\n')
	fmt.Sscanf(line, "Program: %s", &programStr)

    program := make([]int,0)

	strs := strings.Split(programStr, ",")
	for i := 0; i < len(strs); i++ {
		x, _ := strconv.Atoi(strs[i])
        program = append(program, x)
	}

    // 2,4 -- B = A % 8
    // 1,7 -- B ^= 7
    // 7,5 -- C = A / 2^B
    // 1,7 -- B ^= 7
    // 4,6 -- B ^= C
    // 0,3 -- A /= 8
    // 5,5 -- print(B%8)
    // 3,0 -- j->0

    // Is the same as:
    // while A
        // print (A % 8) ^ (A / (2 ^ ((A%8)^7))) % 8
        // A /= 8

    // (A) ^ (A / (2 ^ (A^7))) = 0 -> 000
    // (A) ^ (A / (2 ^ (A^7))) = 3 -> ???
    // (A) ^ (A / (2 ^ (A^7))) = 5 -> ???
    // (A) ^ (A / (2 ^ (A^7))) = 4    
    // (A) ^ (A / (2 ^ (A^7))) = 3    
    // (A) ^ (A / (2 ^ (A^7))) = 0    

    // Sample
    // A /= 8
    // print A%8
    
    slices.Reverse(program)
    var solutions = []int{0,1,2,3,4,5,6,7}
    for i, x := range program {
        nextSolutions := make([]int,0)
        for _, y := range solutions {
            if ((y%8) ^ (y / (1 << ((y%8)^7)))) % 8 == x {
                for j := 0; j < 8; j++ {
                    if i != len(program)-1 {
                        nextSolutions = append(nextSolutions, (y*8)+j)
                    } else {
                        nextSolutions = append(nextSolutions, y)
                    }
                }
            }
        }
        solutions = nextSolutions
    }

    var ans int = (1<<62)
    for _, x := range solutions {
        ans = min(ans, x)
    }

    return strconv.Itoa(ans)
}
