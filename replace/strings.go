package replace

import (
	"math/rand"
	"strings"
	"time"
)

const (
	NAME = "ファルコン・パンチ"
)

func NameOfToday(base string) string {
	if base == "" {
		base = NAME
	}

	length := len([]rune(base))

	name := make(map[int]string)
	for i := 0; i < length; i++ {
		name[i] = string([]rune(base)[i])
	}

	output := make([]string, length)
	for i := 0; i < length; {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(length)

		s, ok := name[n]
		if !ok {
			continue
		}

		output[i] = s
		delete(name, n)
		i++

		time.Sleep(time.Nanosecond)
	}

	return strings.Join(output, "")
}
