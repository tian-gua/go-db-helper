package dh

import (
	"reflect"
)

//结构体类型信息
type StructType struct {
	typeName  string                  //结构体名称
	fieldsMap map[string]*StructField //结构体字段
	fields    []*StructField          //用于保证顺序一致
}

//结构体字段信息
type StructField struct {
	Name     string //字段名称
	TypeName string //字段类型
}

//* 获取结构体反射信息
func getReflectInfo(i interface{}) *StructType {
	v := reflect.Indirect(reflect.ValueOf(i))
	t := v.Type()
	st := &StructType{typeName: t.Name(), fieldsMap: make(map[string]*StructField, t.NumField()), fields: []*StructField{}}

	for j := 0; j < t.NumField(); j++ {
		sf := &StructField{Name: t.Field(j).Name, TypeName: t.Field(j).Type.Name()}
		st.fieldsMap[t.Field(j).Name] = sf
		st.fields = append(st.fields, sf)
	}

	return st
}

//* 获取类型的名称
func getTypeName(i interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(i))
	if v.Type().Kind() == reflect.Slice {
		return v.Type().Elem().Name()
	}
	return v.Type().Name()
}

//* 获取结构体指定字段
func getValue(i interface{}, name string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(i))
	return v.FieldByName(name).Interface()
}

//* 获取结构体指定字段的反射
func getReflectValue(i interface{}, name string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(i))
	return v.FieldByName(name).Interface()
}

//* 获取结构体的所有字段的字符串值
func getValues(i interface{}, st *StructType, filter func(s string) bool) []interface{} {
	var is []interface{}
	value := reflect.Indirect(reflect.ValueOf(i))

	for _, v := range st.fields {
		if filter(v.Name) {
			continue
		}

		is = append(is, value.FieldByName(v.Name).Interface())
	}
	return is
}

//* 获取有效(字段不为零值)的字段的名字和值
//* 用2个切片是为了保证顺序一致
func getValid(i interface{}, st *StructType, filter func(s string) bool) ([]string, []interface{}) {
	var vs []interface{}
	ks := []string{}
	value := reflect.Indirect(reflect.ValueOf(i))

	for _, v := range st.fields {
		if filter(v.Name) {
			continue
		}

		fv := value.FieldByName(v.Name)
		if !isZero(fv) {
			ks = append(ks, v.Name)
			vs = append(vs, value.FieldByName(v.Name).Interface())
		}
	}
	return ks, vs
}

//* 判断值是否为零值
func isZero(v reflect.Value) bool {
	if v.IsValid() && !(reflect.Zero(v.Type()).Interface() == v.Interface()) {
		return false
	}
	return true
}

func getStructFieldInterfaces(cols []string, r *registryInfo) (reflect.Value, []interface{}) {
	var filedInterfaces []interface{}
	value := reflect.Indirect(reflect.ValueOf(r.i))
	newValue := reflect.New(value.Type())

	c := r.table.col2filedMap
	for _, v := range cols {
		filedInterfaces = append(filedInterfaces, newValue.Elem().FieldByName(c[v]).Addr().Interface())
	}

	return newValue, filedInterfaces
}

func getStructFieldInterfaces_(cols []string, sprv reflect.Value) []interface{} {
	var fis []interface{}
	for _, v := range cols {
		fis = append(fis, sprv.Elem().FieldByName(columnName2fieldName(v)).Addr().Interface())
	}
	return fis
}
