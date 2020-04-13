package core

import (
	"database/sql"
	"github.com/gogf/gf/os/grpool"
	"husky-task/model"
	"sync"
)

var ContextInstance *Context

type Context struct {
	DBEngine       *sql.DB
	DSN            string
	ExecutorName   string
	ExecutorMap    map[string]string
	IsLeader       bool
	GRPool         *grpool.Pool
	mLock          sync.Mutex
	ChanTask       chan model.Task
	ChanTaskResult chan model.Task
	TaskGRPool     *grpool.Pool
}

func init() {
	ContextInstance = &Context{
		ExecutorMap:    map[string]string{},
		IsLeader:       false,
		GRPool:         grpool.New(10),
		TaskGRPool:     grpool.New(20),
		ChanTask:       make(chan model.Task, 0),
		ChanTaskResult: make(chan model.Task, 0),
	}
}
