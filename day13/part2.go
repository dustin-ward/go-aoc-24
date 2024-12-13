package day13

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Part2(inpt io.Reader) string {
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
        curMachine.Prize.x += 10000000000000
        curMachine.Prize.y += 10000000000000
        machines = append(machines, curMachine)
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
	}

    abs := func(x int)int {
        return max(x, -x)
    }

    var ans int
    for _, machine := range machines {
        // (m.A.x * a) + (m.B.x * b) = m.P.x
        // (m.A.y * a) + (m.B.y * b) = m.P.y
        D := abs((machine.A.x*machine.B.y) - (machine.B.x*machine.A.y))
        if D == 0 {
            continue
        }

        Dx := abs((machine.B.x*machine.Prize.y) - (machine.B.y*machine.Prize.x))
        Dy := abs((machine.A.x*machine.Prize.y) - (machine.A.y*machine.Prize.x))

        if Dx % D != 0 || Dy % D != 0 {
            continue
        }
        a := Dx / D
        b := Dy / D

        ans += 3*a + b
    }

    return strconv.Itoa(ans)
}
