package core

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitEngine(dsn string) error {
	engine, err := sql.Open("mysql", dsn)
	if err == nil {
		engine.SetMaxIdleConns(5)
		engine.SetMaxOpenConns(10)
		ContextInstance.DBEngine = engine
	} else {
		fmt.Println("InitEngine error: ", err.Error())
	}
	return err
}
