package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	u := User{ID: 1, Name: "Alice"}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	fmt.Println("结构体类型:", t.Name())

	// 2. 遍历字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("字段名: %s 值: %v Tag: %s\n", field.Name, value, tag)
	}

	vp := reflect.ValueOf(&u).Elem()
	nameField := vp.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Bob")
	}
	fmt.Println("修改后的结构体:", u)
}
