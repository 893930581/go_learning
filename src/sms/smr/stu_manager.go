package smr

import "fmt"

type Student struct {
	id	 int64
	name string
}

type StudentManager struct {
	Students map[int64]Student
}

func (smr StudentManager) ShowAll() {
	for _, v := range smr.Students {
		fmt.Printf("序号：%v，姓名：%v\n",v.id,v.name)
	}
}

func (smr StudentManager) Add() {
	var (
		id 	 int64
		name string
	)
	fmt.Println("请输入学生序号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)

	stu := Student{id,name}
	smr.Students[stu.id] = stu
	fmt.Printf("添加成功,学生序号为%v,姓名为%v\n",id,name)
}

func (smr StudentManager) Delete() {
	var id int64
	fmt.Println("请输入要删除的学生的id：")
	fmt.Scanln(&id)
	if _,ok := smr.Students[id];ok{
		delete(smr.Students,id)
		fmt.Println("已经删除掉序号为",id,"学生。")
	}else{
		fmt.Println("查无此人。")
	}
}

func (smr StudentManager) Update() {
	var (
		id   int64
		name string)

	fmt.Println("请输入要修改的学生的id：")
	fmt.Scanln(&id)
	stObj,ok := smr.Students[id]
	if !ok {
		fmt.Println("查无此人。")
		return
	}
	fmt.Println("请输入要修改的学生的姓名：")
	fmt.Scanln(&name)
	stObj.name = name
	smr.Students[id] = stObj
	fmt.Println("修改成功。")
}