package executor

import (
	"github.com/gogf/gf/os/grpool"
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

func StartExecutor(name string, dsn string) {
	err := core.InitEngine(dsn)
	if err != nil {
		log.Println("init db engine failed")
		return
	}
	InitGRPool()
	core.Register(name)
	StartClean()
	core.Elector()
	StartRenew()
	StartScanner()
	StartDispatcher()
	StartActuator()
}

func InitGRPool() {
	core.ContextInstance.GRPool = grpool.New(10)
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
