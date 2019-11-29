package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}){
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)//如果我们传进去的不是一个int就会报运行时错误
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}