package husky_task

import (
	"fmt"
	"husky-task/core"
	"husky-task/crontab"
	"husky-task/executor"
	"log"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	core.ContextInstance.DSN = "root:123456@tcp(192.168.80.1:3306)/test?charset=utf8"
	core.ContextInstance.ExecutorName = "executor-1"
	executor.StartExecutor()

	corn, err := crontab.NewCrontab("test-insert", 2000, insertTasks)
	if err != nil || corn == nil {
		log.Println("new crontab failed")
		return
	}
	_ = core.ContextInstance.GRPool.Add(func() {
		corn.Start()
	})

	for true {
		log.Printf("ExecutorMap: %v\n", core.ContextInstance.ExecutorMap)
		log.Printf("GRPool: %v\n", core.ContextInstance.GRPool.Size())
		log.Printf("TaskGRPool: %v\n", core.ContextInstance.TaskGRPool.Size())
		go func() {
			for task := range core.ContextInstance.ChanTask {
				_ = core.ContextInstance.TaskGRPool.Add(func() {
					log.Printf("TaskContext: %v\n", task.Context)
					time.Sleep(500 * time.Millisecond)
					task.Status = core.TaskSuccess
					core.ContextInstance.ChanTaskResult <- task
				})
			}
		}()
		time.Sleep(2 * time.Second)
	}

}

func insertTasks() {
	sql := "INSERT INTO task (`name`,`type`,`context`,`executor_name`,`status`,`need_scan`,`update_time`) " +
		"VALUES ('task1','type1','balabala1','','','Need',NOW()),('task2','type1','balabala2','','','Need',NOW())," +
		"('task3','type2','balabala3','','','Need',NOW()),('task4','type1','balabala4','','','Need',NOW())," +
		"('task5','type2','balabala5','','','Need',NOW())"
	_, err := core.ContextInstance.DBEngine.Exec(sql)
	if err != nil {
		msg := fmt.Sprintf("insert tasks failed, msg: %v", err)
		log.Println(msg)
		return
	}
}
