package dh

import (
	"database/sql"
	"fmt"
	"reflect"
)

func getStmt(s string) (*sql.Stmt, error) {
	tx, db, err := tm.get()
	if err != nil {
		return nil, err
	}

	if tx != nil {
		return tx.Prepare(s)
	}

	if db != nil {
		return db.Prepare(s)
	}

	return nil, fmt.Errorf("unknown error")
}

//* 执行select语句
func query(is interface{}, s string, args []interface{}) error {
	r, err := registry.get(is)
	if err != nil {
		return err
	}

	stmt, err := getStmt(s)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	ssrv := reflect.ValueOf(is).Elem()
	for rows.Next() {
		newsprv, values := getStructFieldInterfaces(cols, r)
		err = rows.Scan(values...)
		if err != nil {
			return err
		}
		ssrv.Set(reflect.Append(ssrv, newsprv.Elem()))
	}
	return nil
}

//* 执行select语句
func queryStrings(s string, ssp *[]string) error {
	stmt, err := getStmt(s)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	if len(cols) > 1 {
		return fmt.Errorf("found %d columns but one", len(cols))
	}

	news := ""
	ssrv := reflect.ValueOf(ssp).Elem()
	for rows.Next() {
		err = rows.Scan(&news)
		if err != nil {
			return err
		}
		ssrv.Set(reflect.Append(ssrv, reflect.ValueOf(news)))
	}
	return nil
}

//* 执行select语句
func queryString(s string, sp *string) error {
	stmt, err := getStmt(s)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	if len(cols) > 1 {
		return fmt.Errorf("found %d columns but one", len(cols))
	}

	rowNumber := 1
	for rows.Next() {
		if rowNumber > 1 {
			return fmt.Errorf("found more than one row")
		}
		err = rows.Scan(sp)
		if err != nil {
			return err
		}
		rowNumber++
	}
	return nil
}

//* 执行select语句
func queryStructs(s string, ssp interface{}) error {
	stmt, err := getStmt(s)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	ssrv := reflect.ValueOf(ssp).Elem()
	newsp := reflect.New(ssrv.Type().Elem())
	for rows.Next() {
		err = rows.Scan(getStructFieldInterfaces_(cols,newsp)...)
		if err != nil {
			return err
		}
		ssrv.Set(reflect.Append(ssrv, newsp.Elem()))
	}
	return nil
}

//* 执行insert语句
//* 返回:
//* 1.新增的id
//* 2.错误信息
func insert(s string, args []interface{}) (int64, error) {
	stmt, err := getStmt(s)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	return r.LastInsertId()
}

//* 执行update语句
//* 返回:
//* 1.受影响的行数
//* 2.错误信息
func update(s string, args []interface{}) (int64, error) {
	stmt, err := getStmt(s)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

//* 执行delete语句
//* 返回:
//* 1.受影响的行数
//* 2.错误信息
func del(s string, arg interface{}) (int64, error) {
	var args []interface{}
	args = append(args, arg)
	return update(s, args)
}
