package models

import "errors"

type Array struct {
	data     []any
	size     int
	capacity int
}

var MAX_CAPACITY = 100

func NewArray(capacity int) *Array {

	if capacity > MAX_CAPACITY {
		return 0, errors.New("max capacity is 100")
	}

	return &Array{
		data:     make([]any, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (array *Array) Append(value any) {
	if array.size < array.capacity {
		array.data = append(array.data, value)
		array.size++
	}
}

func (array *Array) AppendMultiple(values []any) {
	if array.size+len(values) < array.capacity {
		array.data = append(array.data, values)
		array.size += len(values)
	}
}

func (array *Array) Slice(begin int, end int) ([]any, error) {
	var slice []any

	// negative inputs
	if begin < 0 || end < 0 {
		return slice, errors.New("invalid slice input, arguments must be 0 or graeter")
	}

	// out of range inputs
	if begin > array.size || end > array.size {
		return slice, errors.New("invalid slice input, arguments must be within array size")
	}

	// nominal case
	if begin < end {
		for i := begin; i < end; i++ {
			slice = append(slice, i)
		}
	}
	// reverse
	if begin > end {
		for i := end; i < begin; i++ {
			slice = append(slice, i)
		}
	}
	return slice, nil
}

func (array *Array) Swap(src int, dst int) error {
	if src > array.size || dst > array.size {
		return errors.New("out of bounds swap")
	}
	var temp = array.Get(src)
	array.Set(src, array.Get(dst))
	array.Set(dst, temp)
	return nil
}

func (array *Array) Get(index int) (any, error) {
	if index < 0 || index >= array.size {
		return 0, errors.New("index out of range")
	}
	return array.data[index], nil
}

func (array *Array) Set(index int, value any) error {
	if index < 0 || index >= array.size {
		return errors.New("index out of range")
	}
	array.data[index] = value
	return nil
}

func (array *Array) Size() int {
	return array.size
}

func (array *Array) Capacity() int {
	return array.capacity
}
