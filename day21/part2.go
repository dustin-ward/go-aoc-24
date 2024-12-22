package day21

import (
	"bufio"
	"io"
	"strconv"
)

func Part2(inpt io.Reader) string {
	codes := make([]string, 0)

	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
			break
		}
		line = line[:len(line)-1]
		codes = append(codes, line)
	}

	memo := [4][3][4][3][27]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for i2 := 0; i2 < 4; i2++ {
				for j2 := 0; j2 < 3; j2++ {
					for r := 0; r < 27; r++ {
						memo[i][j][i2][j2][r] = -1
					}
				}
			}
		}
	}

	type xy struct {
		y int
		x int
	}
	keypadBase := map[rune]xy{
		'0': {0, 1},
		'A': {0, 2},
		'1': {1, 0},
		'2': {1, 1},
		'3': {1, 2},
		'4': {2, 0},
		'5': {2, 1},
		'6': {2, 2},
		'7': {3, 0},
		'8': {3, 1},
		'9': {3, 2},
	}
	keypadDir := map[rune]xy{
		'<': {0, 0},
		'v': {0, 1},
		'>': {0, 2},
		'^': {1, 1},
		'A': {1, 2},
	}

	var f func(i1, j1, i2, j2, r int) int
	f = func(i1, j1, i2, j2, r int) int {
		if r == 0 {
			return 1
		}

        if memo[i1][j1][i2][j2][r] != -1 {
            return memo[i1][j1][i2][j2][r]
        }

		start := xy{i1, j1}
		target := xy{i2, j2}

		right := j2 > j1
		up := i2 > i1

		var ans int = (1 << 62)

		pathCost := func(path []int) int {
			cost := 0
			cur := keypadDir['A']
			for _, k := range path {
				var nxt xy
				switch k {
				case 0:
					nxt = keypadDir['<']
				case 1:
					nxt = keypadDir['v']
				case 2:
					nxt = keypadDir['>']
				case 3:
					nxt = keypadDir['^']
				}
				cost += f(cur.y, cur.x, nxt.y, nxt.x, r-1)
				cur = nxt
			}
			nxt := keypadDir['A']
			cost += f(cur.y, cur.x, nxt.y, nxt.x, r-1)

			return cost
		}

		var f2 func(pos xy, path []int)
		f2 = func(pos xy, path []int) {
			if pos == target {
				ans = min(ans, pathCost(path))
				return
			}

			var invalid xy
			if r == 26 {
				invalid = xy{0, 0}
			} else {
				invalid = xy{1, 0}
			}

			if pos.y != target.y {
				nxt := pos
				if up {
					nxt.y++
					if nxt != invalid {
						f2(nxt, append(path, 3))
					}
				} else {
					nxt.y--
					if nxt != invalid {
						f2(nxt, append(path, 1))
					}
				}
			}

			if pos.x != target.x {
				nxt := pos
				if right {
					nxt.x++
					if nxt != invalid {
						f2(nxt, append(path, 2))
					}
				} else {
					nxt.x--
					if nxt != invalid {
						f2(nxt, append(path, 0))
					}
				}
			}
		}
		f2(start, []int{})

        memo[i1][j1][i2][j2][r] = ans
		return ans
	}

	var ans int
	for _, code := range codes {
		pressLength := 0
		cur := keypadBase['A']
		for _, ch := range code {
			target := keypadBase[ch]
			pressLength += f(cur.y, cur.x, target.y, target.x, 26)
			cur = target
		}

		num, _ := strconv.Atoi(code[:3])
		ans += num * pressLength
	}

	return strconv.Itoa(ans)
}
