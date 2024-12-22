package day22

import (
	"bufio"
	"io"
	"strconv"
)

func Part2(inpt io.Reader) string {
    buyers := [][2001]int{}
    prices := [][2001]int{}
    delta := [][2001]int{}
	reader := bufio.NewReader(inpt)
    bIdx := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]
        x, _ := strconv.Atoi(line)
        buyers = append(buyers, [2001]int{})
        prices = append(prices, [2001]int{})
        delta = append(delta, [2001]int{})
        buyers[bIdx][0] = x
        prices[bIdx][0] = x%10
        delta[bIdx][0] = 0
        bIdx++
	}

    const MOD = 16777216

    for i := range buyers {
        for j := 1; j <= 2000; j++ {
            temp := buyers[i][j-1]
            x := temp * 64           
            temp ^= x
            temp = temp % MOD
            x = temp / 32
            temp ^= x
            temp = temp % MOD
            x = temp * 2048
            temp ^= x
            temp = temp % MOD

            buyers[i][j] = temp
            prices[i][j] = temp % 10
            delta[i][j] = prices[i][j] - prices[i][j-1]
        }
    }

    ansMap := make(map[string]int)
    for i := range buyers {
        buyersMap := make(map[string]bool)
        for j := 1; j<=2000-3; j++ {
            var key string
            key += strconv.Itoa(delta[i][j]) + ","
            key += strconv.Itoa(delta[i][j+1]) + ","
            key += strconv.Itoa(delta[i][j+2]) + ","
            key += strconv.Itoa(delta[i][j+3])

            if _, ok := buyersMap[key]; !ok {
                buyersMap[key] = true
                ansMap[key] += prices[i][j+3]
            }
        }
    }

    var ans int
    for _, v := range ansMap {
        ans = max(ans, v)
    }
    return strconv.Itoa(ans)
}
