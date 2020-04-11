package core

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	TaskSelectSQL = "SELECT id FROM `task` WHERE `executor_name`=''"
	TaskUpdateSQL = "UPDATE `task` SET `executor_name`='%s',`status`='%s' WHERE id='%s'"
)

func Dispatcher() {
	if !ContextInstance.IsLeader {
		return
	}

	rows, err := ContextInstance.DBEngine.Query(TaskSelectSQL)
	if err != nil {
		msg := fmt.Sprintf("register executor failed, msg: %v", err)
		println(msg)
	}

	var ids []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			msg := fmt.Sprintf("register executor failed, msg: %v", err)
			println(msg)
		}
		ids = append(ids, id)
	}

	executorSlice := mapToSlice()
	if len(executorSlice) <= 0 {
		return
	}
	rand.Seed(time.Now().Unix())
	for _, id := range ids {
		executorName := executorSlice[rand.Intn(len(executorSlice))]
		sql := fmt.Sprintf(TaskUpdateSQL, executorName, "Distributed", id)
		_, err := ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("register executor failed, msg: %v", err)
			println(msg)
		}
	}
}

func mapToSlice() []string {
	ContextInstance.m.Lock()
	defer ContextInstance.m.Unlock()
	s := make([]string, 0, len(ContextInstance.ExecutorMap))
	for k, _ := range ContextInstance.ExecutorMap {
		s = append(s, k)
	}
	return s
}
