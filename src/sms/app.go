package main

import (
	"fmt"
	"go_learning/src/sms/smr"
	"os"
)

func showMenu(){
	fmt.Println("welcome to SMS!!")
	fmt.Println(`
					1.Add Student
					2.Delete Student
					3.Show All Students
					4.Update Student
					5.quit`)
}

func main() {
	sm := smr.StudentManager{make(map[int64]smr.Student,100)}
	for {
		showMenu()
		fmt.Print("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("添加学生")
			sm.Add()
		case 2:
			fmt.Println("删除学生")
			sm.Delete()
		case 3:
			fmt.Println("显示所有学生")
			sm.ShowAll()
		case 4:
			fmt.Println("更改学生学生")
			sm.Update()
		case 5:
			fmt.Println("88")
			os.Exit(1)
		default:
			fmt.Println("滚！")

		}
	}



}