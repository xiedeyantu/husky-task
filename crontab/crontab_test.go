package crontab

import (
	"testing"
	"time"
)

func printOK() {
	time.Sleep(3 * time.Second)
	println(time.Now().String() + ":ok")
}

func TestCrontab(t *testing.T) {
	corn, err := NewCrontab("executor", 2000, printOK)
	if err != nil {
		println("new crontab failed")
	}
	corn.Start()
}
