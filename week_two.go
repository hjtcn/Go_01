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

/**
思路：
dao层遇到sql.ErrNoRows时，应该wrap这个error，抛给上层，因为dao层封装对于实体类的数据库的访问，不牵扯业务逻辑，
所以在dao层把原始错误存放，并抛给上层。上层再对该错误做处理
 */
