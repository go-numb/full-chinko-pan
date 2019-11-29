package replace

import (
	"fmt"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Println("exec time: ", end.Sub(start))
	}()

	base := "あああああああああああああああああああああああああああああああ！！！！！！！！！！！(ﾌﾞﾘﾌﾞﾘﾌﾞﾘﾌﾞﾘｭﾘｭﾘｭﾘｭﾘｭﾘｭ！！！！！！ﾌﾞﾂﾁﾁﾌﾞﾌﾞﾌﾞﾁﾁﾁﾁﾌﾞﾘﾘｲﾘﾌﾞﾌﾞﾌﾞﾌ"

	var (
		count int
		name  string
	)
	for {
		name = NameOfToday(base)
		if name == base {
			break
		}
		count++
		time.Sleep(time.Millisecond)
	}

	fmt.Printf("%d: %s\n", count, name)
}
