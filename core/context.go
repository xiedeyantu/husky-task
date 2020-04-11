package core

import (
	"database/sql"
	"sync"
)

var ContextInstance = Context{
	ExecutorMap: map[string]string{},
	IsLeader:    false,
}

type Context struct {
	DBEngine     *sql.DB
	ExecutorName string
	ExecutorMap  map[string]string
	IsLeader     bool
	m            sync.Mutex
}
