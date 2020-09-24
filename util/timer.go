package util

import (
	"fmt"
	"time"
)

func ExecuteMeasured(fn func(), description string) {
	s := time.Now()
	fn()
	t := time.Since(s)
	println(fmt.Sprintf("%s in %d ms", description, t.Milliseconds()))
}
