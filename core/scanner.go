package core

import (
	"fmt"
	"log"
)

var (
	ScannerActiveExecutor = "SELECT `name`,`type` FROM `executor`"
)

func Scanner() {
	rows, err := ContextInstance.DBEngine.Query(ScannerActiveExecutor)
	if err != nil {
		msg := fmt.Sprintf("scanner executor failed, msg: %v", err)
		log.Println(msg)
		return
	}

	var (
		executorName   string
		executorType   string
		isLeaderExits  bool
		isMyselfActive bool
	)
	tmpExecutorMap := map[string]string{}
	for rows.Next() {
		err = rows.Scan(&executorName, &executorType)
		if err != nil {
			msg := fmt.Sprintf("scan rows failed, msg: %v", err)
			log.Println(msg)
			return
		}
		tmpExecutorMap[executorName] = executorType
		if executorName == ContextInstance.ExecutorName {
			isMyselfActive = true
			if executorType == ExecutorLeader {
				ContextInstance.IsLeader = true
			}
		}
		if executorType == ExecutorLeader {
			isLeaderExits = true
		}
	}

	// refresh executor map
	ContextInstance.mLock.Lock()
	ContextInstance.ExecutorMap = tmpExecutorMap
	ContextInstance.mLock.Unlock()

	if !isMyselfActive {
		Register()
	}

	if !isLeaderExits {
		Elector()
		Scanner()
	}
}
