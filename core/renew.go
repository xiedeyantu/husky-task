package core

import (
	"fmt"
)

var RenewSQL = "UPDATE executors SET `status`='UP',`renew_time`=now() WHERE `name`='%s'"

func Renew() {
	sql := fmt.Sprintf(RenewSQL, ContextInstance.ExecutorName)
	_, err := ContextInstance.DBEngine.Exec(sql)
	if err != nil {
		msg := fmt.Sprintf("renew executor failed, msg: %v", err)
		println(msg)
	}
}
