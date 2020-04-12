package core

import (
	"fmt"
	"log"
)

var RenewExecutor = "UPDATE `executor` SET `renew_time`=now() WHERE `name`='%s'"

func Renew() {
	sql, err := AssembleSQL(RenewExecutor, ContextInstance.ExecutorName)
	if err != nil {
		msg := fmt.Sprintf("sql assemble error, msg: %v", err)
		log.Println(msg)
		return
	}
	_, err = ContextInstance.DBEngine.Exec(sql)
	if err != nil {
		msg := fmt.Sprintf("renew executor failed, msg: %v", err)
		log.Println(msg)
		return
	}
}
