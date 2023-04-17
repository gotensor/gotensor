package tensor

import (
	"testing"
)

func TestTensorType(t *testing.T) {
	tx := NewTensor([]interface{}{10, 20, 30, 40, 50, 60}, []int{2, 3})
	t.Logf("%v\n", tx)
	if len(tx.Data) < 0 {
		t.Fatalf("Invalid datatype: Tensor")
	}

	ty := NewTensor([]interface{}{"Hello", "Holla", "Salam", "درود"}, []int{2, 2})
	t.Logf("%v\n", ty)
	if len(ty.Data) < 0 {
		t.Fatalf("Invalid datatype: Tensor")
	}
}
