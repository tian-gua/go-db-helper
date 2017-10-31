package dh

//* 表
type tableInfo struct {
	tableName    string
	columns      map[string]*columnInfo
	col2filedMap map[string]string
	columnNames  []string //用于保证顺序一致
}

//* 字段
type columnInfo struct {
	name       string
	typeName   string
	primaryKey bool
}

//* 构造表格信息
func getTable(tableName string, i interface{}) *tableInfo {
	st := getReflectInfo(i)
	columnNames := []string{}
	columnsMap := make(map[string]*columnInfo)
	col2filedMap := make(map[string]string)
	for _, v := range (*st).fields {
		ci := &columnInfo{name: columnName2fieldName(v.Name), typeName: columnTypeParser(v)}
		columnsMap[v.Name] = ci
		columnNames = append(columnNames, ci.name)
		col2filedMap[columnName2fieldName(v.Name)] = v.Name
	}
	return &tableInfo{tableName: tableName, columns: columnsMap, columnNames: columnNames, col2filedMap: col2filedMap}
}

//* 列名-字段名 解析器
var columnName2fieldName func(string string) string

//* 字段名-列名 解析器
var fieldName2columnName func(string string) string

//* 列类型解析器
var columnTypeParser func(*StructField) string
