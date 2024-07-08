// Copyright © ivanlobanov. All rights reserved.
package slice

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Parameters of slice
const (
	_resizeFactor_    int    = 2
	_defaultCapacity_ uint64 = 2
	_defaultLength_   uint64 = 0
)

// Slices must be generics.
type SliceGenerics[T any] struct {
	ptr *[]T
	len uint64
	cap uint64
}

// NewSliceGenerics
// Constructor for sliceGeneric - you can pass any type.
func NewSliceGenerics[T any]() SliceGenerics[T] {
	heapSlice := make([]T, _defaultCapacity_)
	return SliceGenerics[T]{
		ptr: &heapSlice,
		len: _defaultLength_,
		cap: _defaultCapacity_,
	}
}

// ////////// public methods ////////////
// Append ()
// This method is used to append `el T` to []T of structure inner array.
func (th *SliceGenerics[T]) Append(el T) {
	if th.len == th.cap {
		th.resize()
	}
	*th.ptr = append(*th.ptr, el)
	th.len++
}

// PopBack ()
// This method is used to pop the last elements of structure inner array.
func (th *SliceGenerics[T]) PopBack() error {
	if th.len == _defaultLength_ {
		return errors.New("array is empty")
	}
	th.len--
	if th.len*2 < th.cap {
		th.cap /= 2
		newHeapSlice := make([]T, th.cap)
		*th.ptr = append(newHeapSlice, *th.ptr...)
	}
	return nil
}

// Copy()
// This method is used to copy elements from `sl SliceGenerics[T]` to []T of structure inner array.
func (th *SliceGenerics[T]) Copy(sl SliceGenerics[T]) {
	newHeapSlice := make([]T, sl.cap)
	*th.ptr = append(newHeapSlice, *sl.ptr...)
	th.len = sl.len
	th.cap = sl.cap
}

// GetLen ()
// This method is used to get length of []T of structure inner array.
func (th *SliceGenerics[T]) GetLen() uint64 {
	return th.len
}

// GetCapacity ()
// This method is used to get capacity of []T of structure inner array.
func (th *SliceGenerics[T]) GetCapacity() uint64 {
	return th.cap
}

// GetElements ()
// This method is used to get elements of []T of structure inner array.
func (th *SliceGenerics[T]) GetElements() []T {
	return *th.ptr
}

// Print ()
// This method is used to print elements of []T of structure inner array.
func (th *SliceGenerics[T]) Print() error {
	stdout := bufio.NewWriter(os.Stdout)
	_, err := stdout.WriteString(fmt.Sprintf("%v", *th.ptr))
	if err != nil {
		return err
	}
	err = stdout.Flush()
	if err != nil {
		return err
	}
	return nil
}

// private methods
// resize()
// This method is used to resize inner array.
func (th *SliceGenerics[T]) resize() {
	th.cap *= 2
	newHeapSlice := make([]T, th.cap)
	copy(newHeapSlice, *th.ptr)
	th.ptr = &newHeapSlice
}
