package core

import (
	"database/sql"
	"github.com/gogf/gf/os/grpool"
	"husky-task/model"
	"sync"
)

var ContextInstance = Context{
	ExecutorMap:    map[string]string{},
	IsLeader:       false,
	ChanTask:       make(chan model.Task, 0),
	ChanTaskResult: make(chan model.Task, 0),
}

type Context struct {
	DBEngine       *sql.DB
	ExecutorName   string
	ExecutorMap    map[string]string
	IsLeader       bool
	GRPool         *grpool.Pool
	mLock          sync.Mutex
	ChanTask       chan model.Task
	ChanTaskResult chan model.Task
	TaskGRPool     *grpool.Pool
}
