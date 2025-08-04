package utils

import (
	"fmt"
	"time"
)

type Timer struct {
	start time.Time
}

func (t *Timer) DiffMilliAnd() {

}

func (t *Timer) Start() {
	t.start = time.Now()
}

func (t *Timer) DiffMilli() int64 {
	return time.Since(t.start).Milliseconds()
}

func (t *Timer) PrintDiffMilli(action string) {
	fmt.Printf("Time it took for <%s> was <%d> milliseconds.\n", action, t.DiffMilli())
}
