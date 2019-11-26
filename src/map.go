package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name":   "wangbo",
		"course": "golang",
		"site":   "ccmouse",
	}
	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m1, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("->Getting values")
	courseName, ok := m1["course"]
	fmt.Println(courseName, ok)

	fmt.Println("->is value there")
	name, ok := m1["casas"]
	fmt.Println(name, ok)

	fmt.Println("->is value there and Getting values")
	if courseName, ok := m1["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("键不存在")
	}

	fmt.Println("->delete value")
	n, ok := m1["course"]
	fmt.Println(n)
	delete(m1, "course")
}
