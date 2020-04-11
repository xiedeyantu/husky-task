package core

import (
	"fmt"
)

var (
	ElectorSelectSQL = "SELECT `name` FROM executors WHERE `status`='UP' ORDER BY id LIMIT 1"
	ElectorUpdateSQL = "UPDATE executors SET `type`='leader' WHERE `name`='%s'"
)

func Elector() {
	rows, err := ContextInstance.DBEngine.Query(ElectorSelectSQL)
	if err != nil {
		msg := fmt.Sprintf("register executor failed, msg: %v", err)
		println(msg)
	}

	var executorName string
	for rows.Next() {
		err = rows.Scan(&executorName)
		if err != nil {
			msg := fmt.Sprintf("register executor failed, msg: %v", err)
			println(msg)
		}
	}

	if !ContextInstance.IsLeader && executorName == ContextInstance.ExecutorName {
		sql := fmt.Sprintf(ElectorUpdateSQL, executorName)
		_, err := ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("clean executor failed, msg: %v", err)
			println(msg)
		}
		ContextInstance.IsLeader = true
	}
}
