package core

import (
	"fmt"
)

var RegisterSQL = "INSERT INTO executors (`name`,`status`,`renew_time`) VALUES ('%s','%s',now()) " +
	"ON DUPLICATE KEY UPDATE `status`='UP',`renew_time`=now()"

func Register(name string) {
	ContextInstance.ExecutorName = name
	sql := fmt.Sprintf(RegisterSQL, ContextInstance.ExecutorName, "UP")
	_, err := ContextInstance.DBEngine.Exec(sql)
	if err != nil {
		msg := fmt.Sprintf("register executor failed, msg: %v", err)
		println(msg)
	}

}
