package tensor

import (
	"errors"
	"fmt"
	"strings"
)

// Tensor datastructor base model
type Tensor struct {
	Data  []interface{}
	Shape []int
}

func NewTensor(data []interface{}, shape []int) Tensor {
	return Tensor{Data: data, Shape: shape}
}

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

// Get single value
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
