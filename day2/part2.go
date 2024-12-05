package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/dustin-ward/go-aoc-24/util"
)

func Part2(inpt io.Reader) string {
    var ans = 0
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
			return strconv.Itoa(ans)
		}
		line = line[:len(line)-1]

		masterReport := strings.Split(line, " ")
		N := len(masterReport)
		for skip := range N + 1 {
			report := make([]string, N)
			copy(report, masterReport)
			if skip < N {
				report = append(report[:skip], report[skip+1:]...)
			}
			last, _ := strconv.Atoi(report[0])
			lastDir := 0
			for _, v := range report[1:] {
				cur, _ := strconv.Atoi(v)
				if cur == last {
					goto invalid
				}

				del := cur - last
				dir := del / util.Abs(del)
				dis := util.Abs(del)

				if lastDir != 0 && dir != lastDir {
					goto invalid
				}
				if dis > 3 {
					goto invalid
				}

				last = cur
				lastDir = dir
			}
			ans++
			break
		invalid:
		}
	}
}
