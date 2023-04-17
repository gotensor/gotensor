package tensor

import (
	"fmt"
	"strings"
)

// Tensor datastructor base model
type Tensor struct {
	Data  []interface{}
	Shape []int
}

func NewTensor() Tensor {
	return Tensor{}
}

func (t Tensor) String() string {
	var s strings.Builder

	// Print the shape
	fmt.Fprintf(&s, "Shape: %v\n", t.Shape)

	// Print the data
	for i := 0; i < len(t.Data); i++ {
		fmt.Fprintf(&s, "%v ", t.Data[i])
		if ((i + 1) % t.Shape[len(t.Shape)-1].(int)) == 0 {
			fmt.Fprintln(&s)
		}
	}

	return s.String()
}
