package _case

import (
	"errors"
	"fmt"
	"reflect"
)

func ReflectCase() {
	type user struct {
		ID    int64
		Name  string
		Age   int
		Hobby []string
	}
	type outUser struct {
		ID    int64
		Name  string
		Age   int
		Hobby []string
	}

	u := user{ID: 1, Name: "张三", Age: 18, Hobby: []string{"篮球", "足球"}}
	out := outUser{}

	res := copy(&out, u)
	fmt.Println(res, out)

	listUser := []user{
		{ID: 1, Name: "张三", Age: 18, Hobby: []string{"篮球", "足球"}},
		{ID: 2, Name: "李四", Age: 19, Hobby: []string{"篮球", "排球"}},
		{ID: 3, Name: "王五", Age: 20, Hobby: []string{"篮球", "羽毛球"}},
	}
	list := sliceColum(listUser, "Hobby")
	fmt.Println(list)
}

func copy(dest interface{}, source interface{}) error {
	sType := reflect.TypeOf(source)
	sValue := reflect.ValueOf(source)
	//如果为指针类型则获取其值
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	dType := reflect.TypeOf(dest)
	dValue := reflect.ValueOf(dest)
	if dType.Kind() != reflect.Ptr {
		return errors.New("目标对象必须为struct指针类型")
	}
	dType = dType.Elem()
	dValue = dValue.Elem()
	if sValue.Kind() != reflect.Struct {
		return errors.New("源对象必须为struct类型或其指针")
	}
	if dValue.Kind() != reflect.Struct {
		return errors.New("目标对象必须为struct类型或其指针")
	}
	destObject := reflect.New(dType)
	for i := 0; i < dType.NumField(); i++ {
		destField := dType.Field(i)
		if sourceField, ok := sType.FieldByName(destField.Name); ok {
			if destField.Type != sourceField.Type {
				continue
			}
			value := sValue.FieldByName(destField.Name)
			destObject.Elem().FieldByName(destField.Name).Set(value)
		}
	}
	dValue.Set(destObject.Elem())
	return nil

}

func sliceColum(slice interface{}, column string) interface{} {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() == reflect.Struct {
		val := v.FieldByName(column)
		return val.Interface()
	}
	if t.Kind() != reflect.Slice {
		return nil
	}
	t = t.Elem()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f, _ := t.FieldByName(column)
	sliceType := reflect.SliceOf(f.Type)
	//确定切片类型
	s := reflect.MakeSlice(sliceType, 0, 0)
	for i := 0; i < v.Len(); i++ {
		o := v.Index(i)
		if o.Kind() == reflect.Struct {
			val := o.FieldByName(column)
			s = reflect.Append(s, val)
		}
		if o.Kind() == reflect.Ptr {
			v1 := o.Elem()
			val := v1.FieldByName(column)
			s = reflect.Append(s, val)
		}
	}
	return s.Interface()

}
