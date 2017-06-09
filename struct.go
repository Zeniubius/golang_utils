package utils

import (
	"reflect"
)



func IsEmptyValue(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}


// 判断item是否在数组里
// 如果数组为空则返回false
func ItemInArray(item string, max []string) (has bool) {
	return ArrayStringIndex(item, max) != -1
}


func ArrayStringIndex(item string, max []string) (index int) {
	index = -1
	if max == nil || len(max) == 0 {
		return
	}
	for i, l := 0, len(max); i < l; i++ {
		if max[i] == item {
			index = i
			return
		}
	}
	return
}


// 获取不为空的在inFields中的 结构体中的字段
func GetNotEmptyFields(obj interface{}, inFields ...string) (fields []string) {
	fields = []string{}
	pointer := reflect.Indirect(reflect.ValueOf(obj))
	types := pointer.Type()
	fieldNum := pointer.NumField()
	for i := 0; i < fieldNum; i++ {
		v := pointer.Field(i)
		name := types.Field(i).Name
		if inFields != nil && len(inFields) != 0 {
			if !ItemInArray(name, inFields) {
				continue
			}
		}
		if IsEmptyValue(v.Interface()) {
			continue
		}
		fields = append(fields, name)
	}
	return
}

//删除指定字段
func RemoveFields(fields []string, delete ...string) (result []string){
	if(len(delete) == 0){
		result = fields
		return
	}

	for _, field := range fields{
		if !ItemInArray(field, delete){
			result = append(result, field)
		}
	}
	return
}

func MergeSlice(s1 []interface{}, s2 []interface{}) []interface{} {
	if s1 == nil {
		if s2 == nil {
			return []interface{}{}
		}
		return s2
	}
	if s2 == nil {
		if s1 == nil {
			return []interface{}{}
		}
		return s1
	}
	temp := []interface{}{}

	for _, v1 := range s1 {
		inS2 := false
		for _, v2 := range s2 {
			if v1 == v2 {
				inS2 = true
			}
		}
		if !inS2 {
			temp = append(temp, v1)
		}
	}
	temp = append(temp, s2...)
	return temp
}

// 求交集
func IntersectionSlice(s1 []interface{}, s2 []interface{}) []interface{} {
	temp := []interface{}{}
	for _, v1 := range s1 {
		inS2 := false
		for _, v2 := range s2 {
			if v1 == v2 {
				inS2 = true
			}
		}
		if inS2 {
			temp = append(temp, v1)
		}
	}

	UnDuplicatesSlice(&temp)
	return temp
}

func UnDuplicatesSlice(is *[]interface{}) {
	t := map[interface{}]bool{}
	temp := []interface{}{}
	for _, i := range *is {
		if t[i] == true {
			continue
		}
		t[i] = true
		temp = append(temp, i)
	}
	*is = temp
}
