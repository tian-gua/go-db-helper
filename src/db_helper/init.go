package dh

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
)

func init() {
	//* 初始化默认Cid序列
	defaultCidSeq = &cidSeq{cid: 0}
	log.Print("初始化Cid序列完成")
	//* 初始化默认连接器
	conn = &connection{}
	log.Print("初始化默认连接器完成")
	//* 初始化映射对象
	registry = &Registry{mappings: make(map[string]*registryInfo)}
	log.Print("初始化映射对象完成")
	fieldName2columnName = func(s string) string {
		columnName := ""
		bs := []byte(s)
		for i, v := range bs {
			if i == 0 {
				v += 32
				columnName += string(v)
				continue
			}

			if v == '_' {
				columnName += "_"
				continue
			}

			if v <= 90 {
				columnName += "_" + string(v+32)
				continue
			}

			columnName += string(v)
		}
		return columnName
	}
	columnName2fieldName = func(s string) string {
		fieldName := ""
		bs := []byte(s)
		previous := ""
		for i, v := range bs {
			if i == 0 {
				fieldName += string(v - 32)
				continue
			}

			if v == '_' {
				previous = "_"
				continue
			}

			if previous == "_" {
				fieldName += string(v - 32)
				previous += string(v)
				continue
			}

			fieldName += string(v)
			previous += string(v)
		}
		return fieldName
	}
	log.Print("初始化列名解析器完成")
	columnTypeParser = func(sf *StructField) string {
		switch sf.TypeName {
		case "int":
			return "INT"
		case "int64":
			return "BIGINT"
		}
		return ""
	}
	log.Print("初始化列类型解析器完成")
	tm = &txManager{txMap: make(map[uint64]*sql.Tx)}
	log.Print("初始化事务管理器完成")
	whereManager = &where{conditions: make(map[uint64]*Conditions)}
	log.Print("初始条件管理器完成")
}
