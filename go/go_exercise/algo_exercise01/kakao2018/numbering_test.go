package kakao2018

import (
	"fmt"
	"testing"
)

func TestHappy100(t *testing.T) {
	start := 1
	size := 15
	prefix	:= "cix"
	postfix := "-5.nm"

	for i := start; i <= size; i++ {
		fmt.Printf("%s%02d%s\n", prefix, i, postfix)
	}
}
