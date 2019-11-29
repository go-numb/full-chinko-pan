package replace

import (
	"math/rand"
	"strings"
	"time"
)

const (
	NAME = "ファルコン・パンチ"
)

func NameOfToday() string {
	length := len([]rune(NAME))

	name := make(map[int]string)
	for i := 0; i < length; i++ {
		name[i] = string([]rune(NAME)[i])
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
