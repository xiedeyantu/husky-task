package core

import (
	"fmt"
)

var (
	Expire         = 60 // second
	CleanUpdateSQL = "DELETE FROM executors WHERE `renew_time` < DATE_SUB(NOW(), INTERVAL %d SECOND)"
)

func Clean() {
	sql := fmt.Sprintf(CleanUpdateSQL, Expire)
	result, err := ContextInstance.DBEngine.Exec(sql)
	if err != nil {
		msg := fmt.Sprintf("clean executor failed, msg: %v", err)
		println(msg)
	}
	count, _ := result.RowsAffected()
	if count > 0 {
		Scanner()
	}
}
