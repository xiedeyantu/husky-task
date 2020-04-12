package model

import "time"

type Task struct {
	Id           int
	Name         string
	Type         string
	Context      string
	ExecutorName string
	NeedScan     string
	Status       string
	UpdateTime   time.Time
}
