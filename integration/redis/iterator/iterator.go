package iterator

type Iterator struct {
    data []interface{}
    index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

func (r *Iterator) HasNext() bool {
	if r.data == nil || len(r.data) == 0 {
		return false
	}
	return r.index < len(r.data)
}

func (r *Iterator) Next() (result interface{}) {
	result = r.data[r.index]
	r.index++
	return result
}


