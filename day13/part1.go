package day13

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Part1(inpt io.Reader) string {
    type xy struct {
        x int
        y int
    }
    type Machine struct {
        A xy
        B xy
        Prize xy
    }

    machines := make([]Machine,0)

	reader := bufio.NewReader(inpt)
	for {
        var curMachine Machine
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]
        fmt.Sscanf(line, "Button A: X+%d, Y+%d", &curMachine.A.x, &curMachine.A.y)
		line, _ = reader.ReadString('\n')
		line = line[:len(line)-1]
        fmt.Sscanf(line, "Button B: X+%d, Y+%d", &curMachine.B.x, &curMachine.B.y)
		line, _ = reader.ReadString('\n')
		line = line[:len(line)-1]
        fmt.Sscanf(line, "Prize: X=%d, Y=%d", &curMachine.Prize.x, &curMachine.Prize.y)
        machines = append(machines, curMachine)
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
	}

    var ans int
    for _, machine := range machines {
        tempAns := int(1E9)
        for a := 0; a<=100; a++ {
            for b := 0; b<=100; b++ {
                if (machine.Prize.x - (machine.A.x*a + machine.B.x*b) == 0) &&
                    (machine.Prize.y - (machine.A.y*a + machine.B.y*b) == 0) {
                    tempAns = min(tempAns, 3*a + b)
                }
            }
        }
        if tempAns != int(1E9) {
            ans += tempAns
        }
    }

    return strconv.Itoa(ans)
}
