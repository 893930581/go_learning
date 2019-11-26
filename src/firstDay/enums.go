package main

import "fmt"

func enums() {
	const (
		c = iota
		_
		java
		pthyon
		golang
		javascript
	)
	fmt.Println(c, java, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	enums()
}
