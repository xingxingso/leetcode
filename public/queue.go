package public

type Queue interface {
	IsEmpty() bool
	Pop() int
	Push(val int)
	Peek() int
}

type QueueInt struct {
	head *DoubleList
	tail *DoubleList
}

func NewQueueInt() *QueueInt {
	head, tail := &DoubleList{}, &DoubleList{}
	head.Next = tail
	tail.Pre = head
	return &QueueInt{
		head: head,
		tail: tail,
	}
}

func (q *QueueInt) IsEmpty() bool { return q.head.Next == q.tail }

func (q *QueueInt) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	node := q.head
	q.head.Next = node.Next
	return node.Val
}

func (q *QueueInt) Push(val int) {
	node := &DoubleList{
		Val: val,
	}
	node.Next = q.tail
	node.Pre = q.tail.Pre
	node.Pre.Next = node
	q.tail.Pre = node
}

func (q *QueueInt) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.head.Next.Val
}

type simpleQueue struct {
	data []int
}

func (q *simpleQueue) Push(val int) {
	q.data = append(q.data, val)
}

func (q *simpleQueue) Pop() int {
	n := q.data[0]
	q.data = q.data[1:]
	return n
}

func (q *simpleQueue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *simpleQueue) size() int {
	return len(q.data)
}

func (q *simpleQueue) Peek() int {
	return q.data[0]
}
