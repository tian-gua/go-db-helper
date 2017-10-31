package dh

import (
	"sync"
	"fmt"
	"log"
)

//* 表-结构体 映射信息集合
type Registry struct {
	mappings map[string]*registryInfo
	sync.Mutex
}

//* 表-结构体 映射信息
type registryInfo struct {
	typeName    string
	tableName   string
	reflectInfo *StructType
	table       *tableInfo
	i           interface{}
}

//* 添加映射信息
//* table---struct
func (r *Registry) register(table string, i interface{}) error {
	r.Lock()
	defer r.Unlock()
	st := getReflectInfo(i)

	if _, ok := r.mappings[st.typeName]; ok {
		return fmt.Errorf("duplicate mapping")
	}

	r.mappings[st.typeName] = &registryInfo{typeName: st.typeName, tableName: table, reflectInfo: st, table: getTable(table, i), i: i}
	if debug {
		log.Print(fmt.Sprintf("注册信息:%s---%s", table, st.typeName))
	}

	return nil
}

//* 获取映射信息
func (r *Registry) get(i interface{}) (*registryInfo, error) {
	r.Lock()
	defer r.Unlock()
	typeName := getTypeName(i)
	if v, ok := r.mappings[typeName]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("no registry info found: type = %s", typeName)
}

var registry *Registry
