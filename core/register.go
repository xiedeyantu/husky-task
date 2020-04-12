package core

import (
	"fmt"
	"log"
)

var RegisterExecutor = "INSERT INTO `executor` (`name`,`renew_time`) VALUES ('%s',now()) " +
	"ON DUPLICATE KEY UPDATE `renew_time`=now()"

func Register(name string) {
	sql, err := AssembleSQL(RegisterExecutor, name)
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
	ContextInstance.ExecutorName = name
}
