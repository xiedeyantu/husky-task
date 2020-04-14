package core

import (
	"fmt"
	"log"
)

var RegisterExecutor = "INSERT INTO `executor` (`name`,`renew_time`) VALUES ('%s',now())"

func Register() {
	if ContextInstance.ExecutorName == "" {
		msg := "executor name should not be empty"
		log.Println(msg)
		return
	}
	sql, err := AssembleSQL(RegisterExecutor, ContextInstance.ExecutorName)
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
