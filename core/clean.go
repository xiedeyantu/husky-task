package core

import (
	"fmt"
	"log"
)

var (
	Expire              = "60" // second
	CleanExpireExecutor = "DELETE FROM `executor` WHERE `renew_time` < DATE_SUB(NOW(), INTERVAL %s SECOND)"
)

func Clean() {
	sql, err := AssembleSQL(CleanExpireExecutor, Expire)
	if err != nil {
		msg := fmt.Sprintf("sql assemble error, msg: %v", err)
		log.Println(msg)
		return
	}
	result, err := ContextInstance.DBEngine.Exec(sql)
	if err != nil {
		msg := fmt.Sprintf("clean executor failed, msg: %v", err)
		log.Println(msg)
		return
	}

	// If there are deleted actuators, refresh the executor map cache
	count, _ := result.RowsAffected()
	if count > 0 {
		Scanner()
	}
}

func CleanOldTask() {

}
