package day22

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(inpt io.Reader) string {
    buyers := make([]int,0)
	reader := bufio.NewReader(inpt)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        x, _ := strconv.Atoi(line)
        buyers = append(buyers, x)
	}

    const MOD = 16777216

    for _ = range 2000 {
        for i := range buyers {
            x := buyers[i] * 64           
            buyers[i] ^= x
            buyers[i] = buyers[i] % MOD
            x = buyers[i] / 32
            buyers[i] ^= x
            buyers[i] = buyers[i] % MOD
            x = buyers[i] * 2048
            buyers[i] ^= x
            buyers[i] = buyers[i] % MOD
        }
    }

    var ans int
    for _, x := range buyers {
        ans += x
    }

    return strconv.Itoa(ans)
}
