package day17

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part1(inpt io.Reader) string {
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

    output := ""

    combo := func(arg int)int {
        switch arg {
        case 0:
            return 0
        case 1:
            return 1
        case 2:
            return 2
        case 3:
            return 3
        case 4:
            return A
        case 5:
            return B
        case 6:
            return C
        case 7:
            panic("RESERVED")
        default:
            panic("AAH")
        }
    }

    pc := 0
    for pc < len(program) {
        arg := program[pc+1]
        switch program[pc] {
        case 0:
            A = A / (1<<combo(arg))
        case 1:
            B = B ^ arg
        case 2:
            B = combo(arg) % 8
        case 3:
            if A != 0{
                pc = arg-2
            }
        case 4:
            B = B ^ C
        case 5:
            output += strconv.Itoa(combo(arg)%8) + ","
        case 6:
            B = A / (1<<combo(arg))
        case 7:
            C = A / (1<<combo(arg))
        }
        pc += 2
    }

	return output
}
