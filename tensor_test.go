package tensor

import (
	"testing"
)

func TestTensorType(t *testing.T) {
	tx := Tensor{[]int{1, 2, 3}}

	if len(tx) < 1 {
		t.Fatalf("Invalid datatype: Tensor")
	}
}
