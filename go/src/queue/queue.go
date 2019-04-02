package queue

//interface{}  any任何类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
	//如果想限定int类型,用v.(int)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
