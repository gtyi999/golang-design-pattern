package factorymethod

import "testing"


func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)

	factory = PlusOperatorFactory{}
	if Compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}

	factory = MinusOperatorFactory{}
	if Compute(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
