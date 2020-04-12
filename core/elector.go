package core

import (
	"fmt"
	"log"
)

var (
	QueryLeader = "SELECT `name` FROM `executor` ORDER BY `id` LIMIT 1"
	SetLeader   = "UPDATE `executor` SET `type`='leader' WHERE `name`='%s'"
)

func Elector() {
	rows, err := ContextInstance.DBEngine.Query(QueryLeader)
	if err != nil {
		msg := fmt.Sprintf("elector executor failed, msg: %v", err)
		log.Println(msg)
		return
	}

	var executorName string
	for rows.Next() {
		err = rows.Scan(&executorName)
		if err != nil {
			msg := fmt.Sprintf("scan rows failed, msg: %v", err)
			log.Println(msg)
			return
		}
	}

	if !ContextInstance.IsLeader && executorName == ContextInstance.ExecutorName {
		sql, err := AssembleSQL(SetLeader, executorName)
		if err != nil {
			msg := fmt.Sprintf("sql assemble error, msg: %v", err)
			log.Println(msg)
			return
		}
		_, err = ContextInstance.DBEngine.Exec(sql)
		if err != nil {
			msg := fmt.Sprintf("set leader failed, msg: %v", err)
			println(msg)
		}
		ContextInstance.IsLeader = true
	}
}
