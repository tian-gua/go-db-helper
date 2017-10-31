package dh

import (
	"fmt"
	"log"
	"reflect"
)

func idFilter(s string) bool {
	if s == "Id" {
		return true
	}
	return false
}

func generateQuerySql(is interface{}, c *Conditions) (string, []interface{}, error) {
	r, err := registry.get(is)
	if err != nil {
		return "", nil, err
	}

	s := ""
	if c != nil {
		s = fmt.Sprintf("SELECT * FROM %s %s", r.tableName, c.condition)

		if debug {
			log.Print("SQL: " + s)
			log.Print(fmt.Sprintf("args: %v", c.args))
		}
		return s, c.args, nil
	} else {
		s = fmt.Sprintf("SELECT * FROM %s", r.tableName)

		if debug {
			log.Print("SQL: " + s)
		}
		return s, nil, nil
	}
}

func generateInsertSql(i interface{}) (string, []interface{}, error) {
	r, err := registry.get(i)
	if err != nil {
		return "", nil, err
	}

	cols, n := join(r.table.columnNames)
	s := fmt.Sprintf("INSERT INTO %s (id,%s) VALUES (default,%s)", r.tableName, cols, joinQuestionMark(n))
	args := getValues(i, r.reflectInfo, idFilter)

	if debug {
		log.Print("SQL: " + s)
		log.Print(fmt.Sprintf("args: %v", args))
	}
	return s, args, nil
}

func generateInsertSelectiveSql(i interface{}) (string, []interface{}, error) {
	r, err := registry.get(i)
	if err != nil {
		return "", nil, err
	}

	ks, vs := getValid(i, r.reflectInfo, idFilter)
	cols, n := joinWithParse(ks, r.table)
	s := fmt.Sprintf("INSERT INTO %s (id,%s) VALUES (default,%s)", r.tableName, cols, joinQuestionMark(n))
	args := vs

	if debug {
		log.Print("SQL: " + s)
		log.Print(fmt.Sprintf("args: %v", args))
	}
	return s, args, nil
}

func generateUpdateByIdSql(i interface{}) (string, []interface{}, error) {
	r, err := registry.get(i)
	if err != nil {
		return "", nil, err
	}

	id := getValue(i, "Id")
	if isZero(reflect.ValueOf(id)) {
		return "", nil, fmt.Errorf("invalid id: id = %d", id)
	}

	s := fmt.Sprintf("UPDATE %s SET %s WHERE id=?", r.tableName, joinKV(r.table.columnNames))
	args := append(getValues(i, r.reflectInfo, idFilter), id)

	if debug {
		log.Print("SQL: " + s)
		log.Print(fmt.Sprintf("args: %v", args))
	}
	return s, args, nil
}

func generateUpdateByIdSelectiveSql(i interface{}) (string, []interface{}, error) {
	r, err := registry.get(i)
	if err != nil {
		return "", nil, err
	}

	id := getValue(i, "Id")
	if isZero(reflect.ValueOf(id)) {
		return "", nil, fmt.Errorf("invalid id: id = %d", id)
	}

	ks, vs := getValid(i, r.reflectInfo, idFilter)
	s := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", r.tableName, joinKVWithParse(ks, r.table))
	args := append(vs, id)

	if debug {
		log.Print("SQL: " + s)
		log.Print(fmt.Sprintf("args: %v", args))
	}
	return s, args, nil
}

func generateDeleteByIdSql(i interface{}) (string, interface{}, error) {
	r, err := registry.get(i)
	if err != nil {
		return "", "", err
	}

	id := getValue(i, "Id")
	if isZero(reflect.ValueOf(id)) {
		return "", "", fmt.Errorf("invalid id: id = %d", id)
	}

	s := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.tableName)

	if debug {
		log.Print("SQL: " + s)
		log.Print(fmt.Sprintf("args: %v", id))
	}
	return s, id, nil
}

//* 拼接一个表的列名
func join(ss []string) (string, int) {
	s := ""
	i := 0
	for _, v := range ss {
		if v == "id" {
			continue
		}
		if i != 0 {
			s += ","
		}
		s += v
		i++
	}
	return s, i
}

//* 拼接问号
func joinQuestionMark(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		if i != 0 {
			s += ","
		}
		s += "?"
	}
	return s
}

//* 拼接一个表的列名
func joinWithParse(ss []string, t *tableInfo) (string, int) {
	s := ""
	i := 0
	for _, v := range ss {
		if v == "id" {
			continue
		}
		if i != 0 {
			s += ","
		}
		s += t.columns[v].name
		i++
	}
	return s, i
}

//* 拼接k=v
func joinKV(ks []string) string {
	kv := ""
	i := 0
	for _, v := range ks {
		if v == "id" {
			continue
		}
		if i != 0 {
			kv += ","
		}
		kv += v + " = ?"
		i++
	}
	return kv
}

//* 拼接k=v
func joinKVWithParse(ks []string, t *tableInfo) string {
	kv := ""
	i := 0
	for _, v := range ks {
		if v == "Id" {
			continue
		}
		if i != 0 {
			kv += ","
		}
		kv += t.columns[v].name + " = ?"
		i++
	}
	return kv
}
