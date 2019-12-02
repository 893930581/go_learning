package dk

import "fmt"

type Ducker interface {
	Quack()
}

func DoQuck (d Ducker) {
	d.Quack()
}

type Duck struct {
	name string
}

func (d *Duck) Quack() {
	fmt.Println("嘎嘎嘎")
}

type Man struct {
	name string
}

func (m Man) Quack() {
	fmt.Println("女王大人")
}

//func main(){
//	var ducker Ducker
//	ducker = &Duck{name:"吉利"}
//	duck := Duck{name:"唐老鸭"}
//	man := Man{name:"比利"}
//	DoQuck(&duck)
//	DoQuck(man)
//	if name,ok := ducker.(*Duck) ; ok {
//		fmt.Println(name)
//	}
//}