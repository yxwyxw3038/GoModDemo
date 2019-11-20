package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ToString(t interface{}) string {
	if t == nil {
		return ""
	}
	s := t.(string)
	return s
}

func ToInt(t interface{}) int {
	if t == nil {
		return 0
	}
	s := 0
	switch v := t.(type) {
	case int:
		s = t.(int)
	// case int32:
	// 	strInt32 := strconv.FormatInt(v, 10)
	// 	s ,_ := strconv.Atoi(strInt32)
	case int64:
		strInt64 := strconv.FormatInt(v, 10)
		s, _ = strconv.Atoi(strInt64)
	case string:
		s, _ = strconv.Atoi(v)
	default:
		s = 0
	}
	return s
}
func GetNullToStr(s string) string {
	s = strings.Trim(s, " ")
	if s == "" {
		s = " "
	}
	return s
}

// GetTagName 获取结构体中Tag的值，如果没有tag则返回字段值
func GetTagName(structName interface{}, tagstr string) []string {
	// 获取type
	tag := reflect.TypeOf(structName)
	// 如果是反射Ptr类型, 就获取他的 element type
	if tag.Kind() == reflect.Ptr {
		tag = tag.Elem()
	}

	// 判断是否是struct
	if tag.Kind() != reflect.Struct {
		fmt.Println("Check type error not Struct")
		return nil
	}
	fmt.Println(tag.Kind())
	// 获取字段数量
	fieldNum := tag.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		// tag 名字
		tagName := tag.Field(i).Tag.Get(tagstr)
		// if tagName != IGNORE {
		// tag为-时, 不解析
		if tagName == "-" || tagName == "" {
			// 字段名字
			tagName = tag.Field(i).Name
		}
		result = append(result, tagName)
		// }
	}
	return result
}

func GetMapByStruct(st interface{}) (map[string]interface{}, error) {
	// 获取type
	rt := reflect.TypeOf(st)
	rv := reflect.ValueOf(st)
	// 如果是反射Ptr类型, 就获取他的 element type
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	result := make(map[string]interface{})
	// 判断是否是struct
	if rt.Kind() != reflect.Struct {
		return result, errors.New("不是结构无法反射成相应Map")
	}
	// 获取字段数量
	fieldNum := rt.NumField()
	for i := 0; i < fieldNum; i++ {
		Name := rt.Field(i).Name
		v := rv.Field(i).Interface()
		result[Name] = v

	}
	return result, nil
}
func SetStructByMap(st interface{}, hasMap map[string]interface{}) {
	rt := reflect.TypeOf(st)
	rv := reflect.ValueOf(st)
	// 如果是反射Ptr类型, 就获取他的 element type
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	// 获取字段数量
	fieldNum := rt.NumField()
	for i := 0; i < fieldNum; i++ {
		Name := rt.Field(i).Name
		if rv.Field(i).CanSet() {
			if v, ok := hasMap[Name]; ok {
				dataVal := reflect.ValueOf(v)
				rv.Field(i).Set(dataVal)
			}
		}

	}
}
func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}
func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) {
	srcv := reflect.ValueOf(SrcStructPtr)
	dstv := reflect.ValueOf(DstStructPtr)
	srct := reflect.TypeOf(SrcStructPtr)
	dstt := reflect.TypeOf(DstStructPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := DeepFields(reflect.ValueOf(SrcStructPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}
