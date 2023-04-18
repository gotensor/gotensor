package gotensor

import (
	"errors"
	"fmt"
	"strings"
)

// Tensor represents a multidimensional array of values.
type Tensor struct {
	// Data holds the tensor values.
	Data []interface{}

	// Shape holds the tensor dimensions.
	Shape []int
}

// NewTensor creates a new tensor with the given data and shape.
func NewTensor(data []interface{}, shape []int) Tensor {
	tx := Tensor{Shape: shape}
	tx.Data = make([]interface{}, tx.Len())
	tx.Data = data
	return tx
}

// Zeros creates a new tensor initialized with zeros of the given shape.
func Zeros(shape []int) Tensor {
	tx := Tensor{Shape: shape}
	tx.Data = make([]interface{}, tx.Len())
	for i := 0; i < tx.Len(); i++ {
		tx.Data[i] = 0
	}
	return tx
}

// Ones creates a new tensor initialized with ones of the given shape.
func Ones(shape []int) Tensor {
	tx := Tensor{Shape: shape}
	tx.Data = make([]interface{}, tx.Len())
	for i := 0; i < tx.Len(); i++ {
		tx.Data[i] = 1
	}
	return tx
}

// String returns a string representation of the tensor.
func (t Tensor) String() string {
	var s strings.Builder

	// Print the shape
	fmt.Fprintf(&s, "Shape: %v\n", t.Shape)

	// Print the data
	for i := 0; i < len(t.Data); i++ {
		fmt.Fprintf(&s, "%v ", t.Data[i])
		if ((i + 1) % t.Shape[len(t.Shape)-1]) == 0 {
			fmt.Fprintln(&s)
		}
	}

	return s.String()
}

// Len returns the total number of elements in the tensor.
func (t *Tensor) Len() int {
	if len(t.Shape) > 0 {
		l := 1
		for _, v := range t.Shape {
			l *= v
		}
		return l
	} else {
		return 0
	}
}

// Get retrieves a single value from the tensor given the indices specified.
// It returns an error if the indices provided are out of bounds.
func (t *Tensor) Get(indices ...int) (interface{}, error) {
	if len(indices) != len(t.Shape) {
		return 0, errors.New("Index dimensions do not match tensor shape")
	}
	index := 0
	stride := 1
	for i := len(t.Shape) - 1; i >= 0; i-- {
		if indices[i] >= t.Shape[i] || indices[i] < 0 {
			return 0, errors.New("Index out of range")
		}
		index += indices[i] * stride
		stride *= t.Shape[i]
	}
	return t.Data[index], nil
}

// Set sets a single value in the tensor given the value to set and the indices specified.
// It returns an error if the indices provided are out of bounds.
func (t *Tensor) Set(value interface{}, indices ...int) error {
	if len(indices) != len(t.Shape) {
		return errors.New("Index dimensions do not match tensor shape")
	}
	index := 0
	stride := 1
	for i := len(t.Shape) - 1; i >= 0; i-- {
		if indices[i] >= t.Shape[i] || indices[i] < 0 {
			return errors.New("Index out of range")
		}
		index += indices[i] * stride
		stride *= t.Shape[i]
	}
	t.Data[index] = value
	return nil
}
