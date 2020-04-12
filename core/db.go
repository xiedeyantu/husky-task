package core

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

const AvailableStr = `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`

func InitEngine(dsn string) error {
	engine, err := sql.Open("mysql", dsn)
	if err == nil {
		engine.SetMaxIdleConns(5)
		engine.SetMaxOpenConns(5)
		ContextInstance.DBEngine = engine
	} else {
		fmt.Println("InitEngine error: ", err.Error())
	}
	return err
}

func AssembleSQL(sql string, args ...interface{}) (string, error) {
	re, _ := regexp.Compile(AvailableStr)
	for _, v := range args {
		if re.MatchString(v.(string)) {
			return "", errors.New("SQL statements are risky")
		}
	}
	sqlFilled := fmt.Sprintf(sql, args...)
	return sqlFilled, nil
}
