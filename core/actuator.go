package core

import (
	"fmt"
	"husky-task/model"
	"log"
)

var (
	QueryTask        = "SELECT `id`,`context` FROM `task` WHERE `need_scan`='Need' and `executor_name`='%s'"
	UpdateTaskStauts = "UPDATE `task` SET `status`='%s',`need_scan`='NotNeed',`update_time`=now() WHERE `id`=%d"
)

func Actuator() {
	sql := fmt.Sprintf(QueryTask, ContextInstance.ExecutorName)
	rows, err := ContextInstance.DBEngine.Query(sql)
	if err != nil {
		msg := fmt.Sprintf("register executor failed, msg: %v", err)
		log.Println(msg)
		return
	}

	task := map[int]string{}
	for rows.Next() {
		var (
			id      int
			context string
		)
		err = rows.Scan(&id, &context)
		if err != nil {
			msg := fmt.Sprintf("register executor failed, msg: %v", err)
			log.Println(msg)
			return
		}
		task[id] = context
	}

	for id, ctx := range task {
		sql := fmt.Sprintf(UpdateTaskStauts, TaskRunning, id)
		_, err := ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("renew executor failed, msg: %v", err)
			log.Println(msg)
			return
		}

		// process business
		entry := model.Task{
			Id:      id,
			Context: ctx,
		}
		ContextInstance.ChanTask <- entry
		errStr := <-ContextInstance.ChanTaskResult
		if errStr != TaskSuccess {
			sql = fmt.Sprintf(UpdateTaskStauts, TaskFailed, id)
			_, err = ContextInstance.DBEngine.Exec(sql)
			if err != nil {
				msg := fmt.Sprintf("task run failed, msg: %v", err)
				log.Println(msg)
				return
			}
			return
		}

		sql = fmt.Sprintf(UpdateTaskStauts, TaskSuccess, id)
		_, err = ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("task run failed, msg: %v", err)
			log.Println(msg)
			return
		}
	}
}
