package replace

import (
	"fmt"
	"testing"
	"time"

	"gonum.org/v1/gonum/stat"
)

func TestString(t *testing.T) {
	base := "ファルコン・パンチ"
	words := len([]rune(base))

	var (
		count int
		name  string

		meanCount = 100

		// 平均回数、平均時間を取得
		counts = make([]float64, meanCount)
		times  = make([]float64, meanCount)
	)

	for i := 0; i < meanCount; i++ {
		count = 0
		start := time.Now()

		for {
			name = NameOfToday(base)
			if name == base {
				break
			}
			count++
		}

		fmt.Printf("%d times success\n", i)
		counts[i] = float64(count)
		times[i] = time.Now().Sub(start).Seconds()
	}

	fmt.Println(base)
	fmt.Printf("culc count: %d, words: %d, words, Count: %.3ftimes/1word, Time: %.3fsec/1word\n", meanCount, words, stat.Mean(counts, nil), stat.Mean(times, nil))
}
