package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"nana"`
	Score int    `json:"score"`
}

func mapStruct(x interface{}){
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		panic("we need a ptr of struct !!")
	}
	for i := 0; i < t.Elem().NumField(); i++ {
		v.Elem().Field(i)
		t.Elem().Field(i)
		fmt.Printf("value? = %v \n",v.Elem().Field(i).Interface())
		fmt.Printf("struct name:%v,struct tag:%v,struct type:%v,struct index:%v \n",t.Elem().Field(i).Name,t.Elem().Field(i).Tag,t.Elem().Field(i).Type,t.Elem().Field(i).Index)
	}
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	d, _ := json.Marshal(stu1)
	fmt.Println(string(d))
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v value:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}



	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	stu2 := new(student)
	mapStruct(stu2)
}