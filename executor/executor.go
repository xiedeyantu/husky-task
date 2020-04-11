package executor

import (
	"husky-task/core"
	"husky-task/crontab"
)

var (
	RenewInterval      = 30 * 1000
	CleanInterval      = 60 * 1000
	ScannerInterval    = 5 * 1000
	DispatcherInterval = 5 * 1000
	ActuatorInterval   = 5 * 1000
)

func StartAll(name string) {
	StartRegister(name)
	StartClean()
	StartElector()
	StartRenew()
	StartScanner()
	StartDispatcher()
	StartActuator()
}

func StartRegister(name string) {
	core.Register(name)
}

func StartClean() {
	corn, err := crontab.NewCrontab("executor-clean", CleanInterval, core.Clean)
	if err != nil || corn == nil {
		println("new crontab failed")
	}
	go corn.Start()
}

func StartRenew() {
	corn, err := crontab.NewCrontab("executor-clean", RenewInterval, core.Renew)
	if err != nil || corn == nil {
		println("new crontab failed")
	}
	go corn.Start()
}

func StartElector() {
	core.Elector()
}

func StartScanner() {
	corn, err := crontab.NewCrontab("executor-clean", ScannerInterval, core.Scanner)
	if err != nil || corn == nil {
		println("new crontab failed")
	}
	go corn.Start()
}

func StartDispatcher() {
	corn, err := crontab.NewCrontab("executor-clean", DispatcherInterval, core.Dispatcher)
	if err != nil || corn == nil {
		println("new crontab failed")
	}
	go corn.Start()
}

func StartActuator() {
	corn, err := crontab.NewCrontab("executor-clean", ActuatorInterval, core.Actuator)
	if err != nil || corn == nil {
		println("new crontab failed")
	}
	go corn.Start()
}
