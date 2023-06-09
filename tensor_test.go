package gotensor

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

func TestTensorGet(t *testing.T) {
	tx := NewTensor([]interface{}{2, 4, 8, 16, 32, 64, 32, 16, 8, 4, 2}, []int{3, 3, 3})
	v, err := tx.Get(0, 0, 2)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if v != 8 {
		t.Fatalf("Invalid value")
	}
}

func TestTensorSet(t *testing.T) {
	tx := NewTensor([]interface{}{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, []int{3, 3, 3})
	err := tx.Set(4, 0, 0, 0)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	v, err := tx.Get(0, 0, 0)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if v != 4 {
		t.Fatalf("Invalid value")
	}

}

func TestTensorLen(t *testing.T) {
	tx := NewTensor([]interface{}{2, 2, 2, 2, 2, 2, 2, 2, 2}, []int{3, 5, 3, 4})
	l := tx.Len()
	if l != 180 {
		t.Fatalf("invalid length")
	}
}

func TestTensorZeros(t *testing.T) {
	tx := Zeros([]int{3, 4, 5, 6})
	v, err := tx.Get(2, 2, 2, 2)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if v != 0 {
		t.Fatalf("Value must be zero")
	}
}

func TestTensorOnes(t *testing.T) {
	tx := Ones([]int{3, 4, 5, 6})
	v, err := tx.Get(2, 2, 2, 2)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if v != 1 {
		t.Fatalf("Value must be one")
	}
}
