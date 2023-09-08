package util

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(query interface{}) {
	*q = append(*q, query)
}

func (q *Queue) Dequeue() interface{} {
	data := (*q)[0]
	*q = (*q)[1:]

	return data
}
