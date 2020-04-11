package core

import (
	"fmt"
	"time"
)

var (
	ActuatorSelectSQL    = "SELECT `id`,`context` FROM `task` WHERE `status`='Distributed' and `executor_name`='%s'"
	SetActuatorStatusSQL = "UPDATE `task` SET `status`='%s',`update_time`=now() WHERE `id`=%d"
)

func Actuator() {
	sql := fmt.Sprintf(ActuatorSelectSQL, ContextInstance.ExecutorName)
	rows, err := ContextInstance.DBEngine.Query(sql)
	if err != nil {
		msg := fmt.Sprintf("register executor failed, msg: %v", err)
		println(msg)
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
			println(msg)
		}
		task[id] = context
	}

	for id, ctx := range task {
		sql := fmt.Sprintf(SetActuatorStatusSQL, "Running", id)
		_, err := ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("renew executor failed, msg: %v", err)
			println(msg)
		}

		// process business
		println("ctx:" + string(ctx))
		time.Sleep(2 * time.Second)

		sql = fmt.Sprintf(SetActuatorStatusSQL, "Success", id)
		_, err = ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("renew executor failed, msg: %v", err)
			println(msg)
		}
	}
}
