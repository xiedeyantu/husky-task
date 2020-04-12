package core

import "testing"

func TestAssembleSQL(t *testing.T) {
	sql := "select * from a where b=%s"

	sqlFilled, err := AssembleSQL(sql, "1")
	if err != nil {
		println("err1:" + err.Error())
	}
	println("sqlFilled1:" + sqlFilled)

	sqlFilled, err = AssembleSQL(sql, "c' and 1=1 b='c")
	if err != nil {
		println("err2:" + err.Error())
	}
	println("sqlFilled2:" + sqlFilled)
}
