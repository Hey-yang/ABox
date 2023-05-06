package Test01

import (
	"errors"
	"fmt"
	"reflect"
)

type user struct {
	ID    int64
	Name  string
	Hobby []string
}
type outUser struct {
	ID    int64
	Name  string
	Hobby []string
}

func ReflectCase() {

	u := user{1, "name", []string{"1", "2"}}
	out := outUser{}
	err := copy(&out, u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	listUser := []user{
		{1, "代号1", []string{"1", "2"}},
		{2, "代号2", []string{"3", "4"}},
		{3, "代号2", []string{"5", "6"}},
	}

	hobbySli, _ := sliceColumn(listUser, "Hobby")
	fmt.Println(hobbySli)

}

// 这里any和 interface{}区别
// 没有区别
// any 有一个真身，本质上是 interface{} 的别名：
func copy(dest interface{}, source interface{}) error {
	sType := reflect.TypeOf(source)
	sValue := reflect.ValueOf(source)
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	if sType.Kind() != reflect.Struct {
		return errors.New("源对象必须要是struct类型或者struct指针类型")
	}
	dType := reflect.TypeOf(dest)
	dValue := reflect.ValueOf(dest)
	if dType.Kind() != reflect.Ptr {
		return errors.New("目标对象必须要是struct指针类型")
	}
	dType = dType.Elem()
	dValue = dValue.Elem()
	if dType.Kind() != reflect.Struct {
		return errors.New("目标对象必须要是struct指针类型")
	}
	destObj := reflect.New(dType)
	for i := 0; i < dType.NumField(); i++ {
		destField := dType.Field(i)
		if sourceField, ok := sType.FieldByName(destField.Name); ok {
			if sourceField.Type != destField.Type {
				continue
			}
			value := sValue.FieldByName(destField.Name)
			destObj.Elem().FieldByName(destField.Name).Set(value)
		}
	}
	dValue.Set(destObj.Elem())
	return nil
}

func sliceColumn(slice any, column string) (any, error) {
	sType := reflect.TypeOf(slice)
	sValue := reflect.ValueOf(slice)
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	if sType.Kind() == reflect.Struct {
		val := sValue.FieldByName(column)
		return val.Interface(), nil
	}
	if sType.Kind() != reflect.Slice {
		return nil, errors.New("对象必须要是Slice类型")
	}
	sType = sType.Elem()
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
	}
	fType, _ := sType.FieldByName(column)
	sliceType := reflect.SliceOf(fType.Type)
	s := reflect.MakeSlice(sliceType, 0, 0)
	for i := 0; i < sValue.Len(); i++ {
		o := sValue.Index(i)
		if o.Kind() == reflect.Struct {
			val := o.FieldByName(column)
			s = reflect.Append(s, val)
		}
	}
	return s, errors.New("对象必须要是Slice类型")
}

func ReflectTest() {
	//Test01.ReflectCase()
	val := "12313"
	s := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	fmt.Println(s, v)
	val2 := &val
	s1 := reflect.TypeOf(val2)
	v1 := reflect.ValueOf(val2)

	fmt.Println(s1, v1)
	fmt.Println(s1.Elem(), v1.Elem())
	fmt.Println(s1.Kind(), v1.Kind())
	fmt.Println(s1.Elem().Kind(), v1.Elem().Kind())

	zVal := &val
	s = reflect.TypeOf(zVal)
	v = reflect.ValueOf(zVal)
	fmt.Println(s, v)
	v.Elem().SetString("11231111")
	fmt.Println(val)

	s = reflect.TypeOf(val)
	v = reflect.ValueOf(val)
	NewVal := reflect.New(s)
	fmt.Println(reflect.TypeOf(NewVal))

	NewVal.Elem().SetString("111111")
	fmt.Println(NewVal.Elem())

	listUser := []user{
		{1, "代号1", []string{"1", "2"}},
		{2, "代号2", []string{"3", "4"}},
		{3, "代号2", []string{"5", "6"}},
	}
	s = reflect.TypeOf(listUser)
	fmt.Println(s.Elem())
}
