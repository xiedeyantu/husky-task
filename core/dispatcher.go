package core

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	QueryNoDispatcherTask = "SELECT id FROM `task` WHERE `executor_name`=''"
	DispatcherTask        = "UPDATE `task` SET `executor_name`='%s',`status`='%s' WHERE id='%s'"
	ReDispatcherExecutor  = "UPDATE `task` SET `executor_name`='%s',`status`='Distributed',`need_scan`='Need' " +
		"WHERE `status`!='Success' and `update_time` < DATE_SUB(NOW(), INTERVAL %s MINUTE)"
	ReDispatcherInterval = "5" // minutes
)

func Dispatcher() {
	if !ContextInstance.IsLeader {
		return
	}

	rows, err := ContextInstance.DBEngine.Query(QueryNoDispatcherTask)
	if err != nil {
		msg := fmt.Sprintf("dispatche task failed, msg: %v", err)
		log.Println(msg)
		return
	}

	var taskIds []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			msg := fmt.Sprintf("scan rows failed, msg: %v", err)
			println(msg)
		}
		taskIds = append(taskIds, id)
	}

	executorSlice := mapToSlice()
	if len(executorSlice) <= 0 {
		return
	}
	rand.Seed(time.Now().Unix())
	for _, taskId := range taskIds {
		executorName := executorSlice[rand.Intn(len(executorSlice))]
		sql := fmt.Sprintf(DispatcherTask, executorName, TaskDistributed, taskId)
		_, err := ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("register executor failed, msg: %v", err)
			log.Println(msg)
			return
		}
	}
}

func ReDispatcher() {
	if !ContextInstance.IsLeader {
		return
	}

	sql, err := AssembleSQL(ReDispatcherExecutor, ContextInstance.ExecutorName, ReDispatcherInterval)
	if err != nil {
		msg := fmt.Sprintf("sql assemble error, msg: %v", err)
		log.Println(msg)
		return
	}
	_, err = ContextInstance.DBEngine.Exec(sql)
	if err != nil {
		msg := fmt.Sprintf("register executor failed, msg: %v", err)
		log.Println(msg)
		return
	}
}

func mapToSlice() []string {
	ContextInstance.mLock.Lock()
	defer ContextInstance.mLock.Unlock()
	s := make([]string, 0, len(ContextInstance.ExecutorMap))
	for k, _ := range ContextInstance.ExecutorMap {
		s = append(s, k)
	}
	return s
}
