package utils

import (
	"time"
)

func Delay() {
	duration := time.Duration(time.Millisecond * 1)
	time.Sleep(duration)
}

func DelayLong() {
	time.Sleep(time.Duration(time.Second * 5))
}
