package husky_task

import (
	"husky-task/core"
	"husky-task/executor"
	"log"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	dsn := "root:123456@tcp(192.168.80.1:3306)/test?charset=utf8"
	executor.StartExecutor("executor-1", dsn)
	for true {
		log.Printf("ExecutorMap: %v\n", core.ContextInstance.ExecutorMap)
		log.Printf("GRPool: %v\n", core.ContextInstance.GRPool.Size())
		go func() {
			for task := range core.ContextInstance.ChanTask {
				log.Printf("TaskContext: %v\n", task.Context)
				time.Sleep(1 * time.Second)
				core.ContextInstance.ChanTaskResult <- core.TaskSuccess
			}
		}()
		time.Sleep(2 * time.Second)
	}
}
