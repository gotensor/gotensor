package tensor

// Tensor datastructor base model
type Tensor struct {
	Data  []interface{}
	Shape []interface{}
}

func NewTensor() Tensor {
	return Tensor{}
}
