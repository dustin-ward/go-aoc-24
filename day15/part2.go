package day15

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part2(inpt io.Reader) string {
    B := make([][]byte,0)

    type xy struct {
        x int
        y int
    }
    var robot xy

    instrs := false
    inst := ""
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]

        if !instrs {
            if line == "" {
                instrs = true
                continue
            }

            newLine := ""
            for _, ch := range line {
                switch(ch) {
                case '#':
                    newLine += "##"
                case 'O':
                    newLine += "[]"
                case '.':
                    newLine += ".."
                case '@':
                    newLine += "@."
                }
            }

            idx := strings.IndexByte(newLine, byte('@'))
            if idx != -1 {
                robot = xy{idx,len(B)}
            }

            B = append(B, []byte(newLine))
        } else {
            inst += line
        }
	}

    var dx = []int{-1,0,1,0}
    var dy = []int{0,-1,0,1}

    type move struct {
        pos xy
        d int
    }
    var todoMoves []move
    var movedMemo map[xy]bool
    
    var checkMove func(i,j,d int)bool
    checkMove = func(i,j,d int)bool {
        if B[i][j] == '#' {
            movedMemo[xy{j,i}] = false
            return false
        }
        if B[i][j] == '.' {
            movedMemo[xy{j,i}] = true
            return true
        }

        i2 := i+dy[d]
        j2 := j+dx[d]
        var toCheck = []xy{{j2,i2}}

        if d%2 != 0 {
            if B[i2][j2] == '[' {
                toCheck = append(toCheck, xy{j2+1,i2})
            } else if B[i2][j2] == ']' {
                toCheck = append(toCheck, xy{j2-1,i2})
            }
        }

        valid := true
        for _, check := range toCheck {
            var r bool
            if v, ok := movedMemo[check]; ok {
                r = v
            } else {
                r = checkMove(check.y, check.x, d)
            }
            valid = valid && r
        }
        if valid {
            todoMoves = append(todoMoves, move{xy{j,i},d})
        }

        movedMemo[xy{j,i}] = valid
        return valid
    }

    applyMoves := func(){
        for _, m := range todoMoves {
            i2 := m.pos.y+dy[m.d]
            j2 := m.pos.x+dx[m.d]
            B[i2][j2] = B[m.pos.y][m.pos.x]
            B[m.pos.y][m.pos.x] = '.'
        }
    }

    for _, ch := range inst {
        todoMoves = make([]move,0)
        movedMemo = make(map[xy]bool)
        switch (ch) {
        case '<':
            if checkMove(robot.y, robot.x, 0) {
                applyMoves()
                robot.x--
            }
        case '^':
            if checkMove(robot.y, robot.x, 1) {
                applyMoves()
                robot.y--
            }
        case '>':
            if checkMove(robot.y, robot.x, 2) {
                applyMoves()
                robot.x++
            }
        case 'v':
            if checkMove(robot.y, robot.x, 3) {
                applyMoves()
                robot.y++
            }
        }
    }

    var ans int
    for i, v := range B {
        for j, ch := range v {
            if ch == byte('[') {
                ans += (100*i)+j
            }
        }
    }

    return strconv.Itoa(ans)
}
