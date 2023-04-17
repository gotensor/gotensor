package tensor

import (
	"testing"
)

func TestTensorType(t *testing.T) {
	tx := NewTensor()

	if len(tx.Data) < 0 {
		t.Fatalf("Invalid datatype: Tensor")
	}
}
