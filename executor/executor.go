package executor

import (
	"husky-task/core"
	"husky-task/crontab"
	"log"
)

var (
	RenewInterval      = 30 * 1000
	CleanInterval      = 60 * 1000
	ScannerInterval    = 5 * 1000
	DispatcherInterval = 5 * 1000
	ActuatorInterval   = 5 * 1000
)

type TaskFunc func(context string) error

func StartExecutor() {
	err := core.InitEngine(core.ContextInstance.DSN)
	if err != nil {
		log.Println("init db engine failed")
		return
	}
	StartClean()
	StartRenew()
	StartScanner()
	StartReDispatcher()
	StartDispatcher()
	StartActuator()
	StartSetTaskResult()
}

func StartClean() {
	corn, err := crontab.NewCrontab("executor-clean", CleanInterval, core.Clean)
	if err != nil || corn == nil {
		log.Println("new crontab failed")
		return
	}
	_ = core.ContextInstance.GRPool.Add(func() {
		corn.Start()
	})
}

func StartRenew() {
	corn, err := crontab.NewCrontab("executor-renew", RenewInterval, core.Renew)
	if err != nil || corn == nil {
		log.Println("new crontab failed")
		return
	}
	_ = core.ContextInstance.GRPool.Add(func() {
		corn.Start()
	})
}

func StartScanner() {
	corn, err := crontab.NewCrontab("executor-scanner", ScannerInterval, core.Scanner)
	if err != nil || corn == nil {
		log.Println("new crontab failed")
		return
	}
	_ = core.ContextInstance.GRPool.Add(func() {
		corn.Start()
	})
}

func StartDispatcher() {
	corn, err := crontab.NewCrontab("executor-dispatcher", DispatcherInterval, core.Dispatcher)
	if err != nil || corn == nil {
		log.Println("new crontab failed")
		return
	}
	_ = core.ContextInstance.GRPool.Add(func() {
		corn.Start()
	})
}

func StartReDispatcher() {
	corn, err := crontab.NewCrontab("executor-redispatcher", DispatcherInterval, core.ReDispatcher)
	if err != nil || corn == nil {
		log.Println("new crontab failed")
		return
	}
	_ = core.ContextInstance.GRPool.Add(func() {
		corn.Start()
	})
}

func StartActuator() {
	corn, err := crontab.NewCrontab("executor-actuator", ActuatorInterval, core.Actuator)
	if err != nil || corn == nil {
		log.Println("new crontab failed")
		return
	}
	_ = core.ContextInstance.GRPool.Add(func() {
		corn.Start()
	})
}

func StartSetTaskResult() {
	_ = core.ContextInstance.GRPool.Add(func() {
		core.SetTaskResult()
	})
}
