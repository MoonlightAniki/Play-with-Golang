package queue

// 使用别名的方式拓展已有类型

// 使用Queue作为切片的别名
type Queue []interface{} //interface{}表示任何类型

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	front := (*q)[0]
	*q = (*q)[1:]
	return front
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
