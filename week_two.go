package main

import (
	"fmt"
	pkg_errors "github.com/pkg/errors"
)

func queryData(exec_sql string) error {
	err := db.QueryRow(exec_sql)

	if err != nil {
		return pkg_errors.Wrapf(err, "EXEC SQL:%s", exec_sql)
	}

	return err
}

/**
上层调用
 */
func upperLayer() {
	exec_sql := "select * from table_name where id = 1"
	err := queryData(exec_sql)

	if err != nil {
		fmt.Printf("origin error: %T %v\n", pkg_errors.Cause(err), pkg_errors.Cause(err))
	}
}
