package main

import (
	"fmt"
	"husky-task/core"
	"husky-task/executor"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	dsn := "root:123456@tcp(192.168.80.1:3306)/test?charset=utf8"
	err := core.InitEngine(dsn)
	if err != nil {
		println("init db engine failed")
	}

	executor.StartAll("executor-1")
	for true {
		fmt.Printf("list: %v\n", core.ContextInstance.ExecutorMap)
		time.Sleep(2 * time.Second)
	}
}
