package main

import "fmt"

type Man struct {
	name string
	age int
}

func (m *Man) eat(){
	(*m).age++
}

func (m Man) sayHello(){
	fmt.Println("hello")
}

func main() {
	capMan := new(Man)
	fmt.Println(capMan)
	capMan2 := Man{name: "wangbo", age: 6}
	capMan2.eat()
	fmt.Println(capMan2)
	fmt.Println(capMan2.name)
	fmt.Printf("%T",capMan)
	fmt.Printf("%T",capMan2)
	capMan.sayHello()
}