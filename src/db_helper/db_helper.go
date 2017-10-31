package dh

var debug = false

//* 用户接口---初始化连接池(内部是连接池实现)
func Connect(url string, port int, userName string, password string, dbName string) error {
	return conn.connect(url, port, userName, password, dbName)
}

//* 用户接口---使用框架本身的cid功能
func GetCid() int64 {
	return getCid()
}

//* 用户接口---注册 表-结构体 映射
func Register(table string, i interface{}) error {
	return registry.register(table, i)
}

//* 用户接口---执行事务
func Tx(f func() error) error {
	return transaction(f)
}

//* 用户接口---查询操作
func Select(is interface{}) error {
	s, args, err := generateQuerySql(is, nil)
	if err != nil {
		return err
	}
	return query(is, s, args)
}

//* 用户接口---查询操作
func SelectStrings(s string, ssp *[]string) error {
	return queryStrings(s, ssp)
}

//* 用户接口---查询操作
func SelectString(s string, sp *string) error {
	return queryString(s, sp)
}

//* 用户接口---查询操作
func SelectStructs(s string, ssp interface{}) error {
	return queryStructs(s, ssp)
}

//* 用户接口---添加查询条件
func WhereEqual(condition string, arg interface{}) *Conditions {
	c := &Conditions{}
	whereManager.conditions[getGID()] = c
	c.condition = "WHERE " + condition + " = ?"
	c.args = append(c.args, arg)
	return c
}

func WhereIsNull(condition string) *Conditions {
	c := &Conditions{}
	whereManager.conditions[getGID()] = c
	c.condition = "WHERE " + condition + "IS NULL"
	return c
}

func WhereIsNotNull(condition string) *Conditions {
	c := &Conditions{}
	whereManager.conditions[getGID()] = c
	c.condition = "WHERE " + condition + "IS NOT NULL"
	return c
}

//* 用户接口---执行插入操作
func Insert(i interface{}) (int64, error) {
	s, args, err := generateInsertSql(i)
	if err != nil {
		return 0, err
	}
	return insert(s, args)
}

//* 用户接口---执行可选的插入操作
func InsertSelective(i interface{}) (int64, error) {
	s, args, err := generateInsertSelectiveSql(i)
	if err != nil {
		return 0, err
	}
	return insert(s, args)
}

//* 用户接口---根据id执行更新操作
func UpdateById(i interface{}) (int64, error) {
	s, args, err := generateUpdateByIdSql(i)
	if err != nil {
		return 0, err
	}
	return update(s, args)
}

//* 用户接口---根据id执行删除操作
func DeleteById(i interface{}) (int64, error) {
	s, args, err := generateDeleteByIdSql(i)
	if err != nil {
		return 0, err
	}
	return del(s, args)
}

//* 用户接口---根据id执行可选的更新操作
func UpdateByIdSelective(i interface{}) (int64, error) {
	s, args, err := generateUpdateByIdSelectiveSql(i)
	if err != nil {
		return 0, err
	}
	return update(s, args)
}

//* 用户接口---开启调试模式
func DebugOn() {
	debug = true
}

//* 用户接口---关闭调试模式
func DebugOff() {
	debug = false
}
