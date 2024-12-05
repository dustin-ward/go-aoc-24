package day3

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strconv"
)

func Part2(inpt io.Reader) string {
	var ans = 0
	rMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	rDo := regexp.MustCompile(`do\(\)`)
	rDont := regexp.MustCompile(`don't\(\)`)

	fullLine := ""
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
			break
		}
		line = line[:len(line)-1]

		fullLine += line
	}

	type instrs struct {
		pos      int
		instType int
		a        int
		b        int
	}
	switches := make([]instrs, 0)

	dos := rDo.FindAllStringIndex(fullLine, -1)
	for _, do := range dos {
		switches = append(switches, instrs{
			do[0],
			1,
			0,
			0,
		})
	}
	donts := rDont.FindAllStringIndex(fullLine, -1)
	for _, dont := range donts {
		switches = append(switches, instrs{
			dont[0],
			-1,
			0,
			0,
		})
	}
	for _, match := range rMul.FindAllStringSubmatchIndex(fullLine, -1) {
		a, _ := strconv.Atoi(fullLine[match[2]:match[3]])
		b, _ := strconv.Atoi(fullLine[match[4]:match[5]])
		switches = append(switches, instrs{
			match[0],
			0,
			a,
			b,
		})
	}

	sort.Slice(switches, func(i, j int) bool {
		return switches[i].pos < switches[j].pos
	})

	state := true
	for _, instr := range switches {
		switch instr.instType {
		case -1:
			state = false
		case 0:
			if state {
				ans += instr.a * instr.b
			}
		case 1:
			state = true

		}
	}

    return strconv.Itoa(ans)
}
