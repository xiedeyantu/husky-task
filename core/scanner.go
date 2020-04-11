package core

import (
	"fmt"
)

var (
	ScannerSelectSQL = "SELECT `name`,`type` FROM executors WHERE `status`='UP'"
)

func Scanner() {
	rows, err := ContextInstance.DBEngine.Query(ScannerSelectSQL)
	if err != nil {
		msg := fmt.Sprintf("register executor failed, msg: %v", err)
		println(msg)
	}

	var (
		executorName  string
		executorType  string
		isExitsLeader bool
		isMyselfAlive bool
	)
	tmpMap := map[string]string{}
	for rows.Next() {
		err = rows.Scan(&executorName, &executorType)
		if err != nil {
			msg := fmt.Sprintf("register executor failed, msg: %v", err)
			println(msg)
		}
		tmpMap[executorName] = executorType
		if executorName == ContextInstance.ExecutorName {
			isMyselfAlive = true
		}
		if executorType == "leader" {
			isExitsLeader = true
		}
	}

	ContextInstance.m.Lock()
	ContextInstance.ExecutorMap = tmpMap
	ContextInstance.m.Unlock()

	if !isMyselfAlive {
		Register(ContextInstance.ExecutorName)
	}

	if !isExitsLeader {
		Elector()
		Scanner()
	}
}
