package main

import (
	"errors"
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	t := reflect.ValueOf(x)
	fmt.Printf("%v\n",t)
}

func reflectSetValue(x interface{}) (err error) {
	t := reflect.TypeOf(x)
	fmt.Println("->",t.Kind(),t.Elem().Kind())
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Float32 {
		err = errors.New("Arg must be a ptr of Int")
		return
	}
	v := reflect.ValueOf(x)
	fmt.Print()
	switch t.Elem().Kind() {
	case reflect.Int64:
		fmt.Printf("type is Int64,value is %d\n",int64(v.Int()))
		v.Elem().SetInt(5000)
	case reflect.Float32:
		if !v.Elem().IsValid() {
			fmt.Println("没有被分配储存空间")
		}else{
			fmt.Printf("type is Float32,value is %f\n",float32(v.Elem().Float()))
		}

		v.Elem().SetFloat(56.1)
	}
	return nil
}

func reflectType(x interface{}){
	t := reflect.TypeOf(x)
	fmt.Printf("%v\n",t)
	fmt.Printf("type %v, prototype %v \n",t.Name(),t.Kind())
}

type myInt int64

type duck interface {
	quack()
}

type person struct {
	name string
}

func main() {
	var a rune			//类型别名
	var b myInt			//自定义类型
	var c *float32 		//指针
	var d person		//其实也算是类型别名，但是结构体
	var e chan int		//管道
	var f []int
	var g [5]int

	fmt.Println(reflect.TypeOf(a))

	reflectType(a)		//type int32	, prototype int32
	reflectType(b)		//type myInt	, prototype int64
	reflectType(c)		//type 			, prototype ptr
	reflectType(d)		//type person	, prototype struct
	reflectType(e)		//type 			, prototype chan
	reflectType(f)		//type 			, prototype slice
	reflectType(g)		//type			, prototype array

	reflectValue(a)		//type int32	, prototype int32
	reflectValue(b)		//type myInt	, prototype int64
	reflectValue(c)		//type 			, prototype ptr
	reflectValue(d)		//type person	, prototype struct
	reflectValue(e)		//type 			, prototype chan
	reflectValue(f)		//type 			, prototype slice
	reflectValue(g)		//type			, prototype array

	var p = new(float32)
	reflectSetValue(p)
	fmt.Print(*p)
}