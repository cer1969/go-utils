// Copyright Cristian Echeverría Rabí

package values

import (
	"fmt"
	"testing"
)

func TestLt(t *testing.T) {
	vc := Checker("Prueba Lt 1")
	vc.Val("Número", 30.0).Lt(31.0)
	err := vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Lt 2")
	vc.Val("Número", 30.0).Lt(30.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc = Checker("Prueba Lt 3")
	vc.Val("Número", 30.0).Lt(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestLe(t *testing.T) {
	vc := Checker("Prueba Le 1")
	vc.Val("Número", 30.0).Le(31.0)
	err := vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Le 2")
	vc.Val("Número", 30.0).Le(30.0)
	err = vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Le 3")
	vc.Val("Número", 30.0).Le(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestGt(t *testing.T) {
	vc := Checker("Prueba Gt 1")
	vc.Val("Número", 30.0).Gt(29.0)
	err := vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Gt 2")
	vc.Val("Número", 30.0).Gt(30.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc = Checker("Prueba Gt 3")
	vc.Val("Número", 30.0).Gt(31.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestGe(t *testing.T) {
	vc := Checker("Prueba Ge 1")
	vc.Val("Número", 30.0).Ge(29.0)
	err := vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Ge 2")
	vc.Val("Número", 30.0).Ge(30.0)
	err = vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Ge 3")
	vc.Val("Número", 30.0).Ge(31.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func TestMixed(t *testing.T) {
	vc := Checker("Prueba Ge Le 1")
	vc.Val("Número", 35.0).Ge(30.0).Le(40.0)
	err := vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Ge Le 2")
	vc.Val("Número", 30.0).Ge(30.0).Le(30.0)
	err = vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Ge Le 3")
	vc.Val("Número", 30.0).Ge(31.0).Le(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc = Checker("Prueba Gt Lt 1")
	vc.Val("Número", 35.0).Gt(30.0).Lt(40.0)
	err = vc.Error()
	if err != nil {
		t.Error("Error not expected", err)
	}

	vc = Checker("Prueba Gt Lt 2")
	vc.Val("Número", 30.0).Gt(30.0).Lt(30.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}

	vc = Checker("Prueba Gt Lt 3")
	vc.Val("Número", 30.0).Gt(31.0).Lt(29.0)
	err = vc.Error()
	if err == nil {
		t.Error("Error expected")
	}
}

func ExampleChecker() {
	vc := Checker("Prueba")
	vc.Val("Número", 30).Gt(40.0).Gt(100.0)
	vc.Val("Sol", 0.0).Gt(0.0).Lt(1.0)
	vc.Val("Sol", 0.5).Gt(0.0).Lt(1.0)
	err, _ := vc.Error().(*ValError) // Verifica y convierte vc.Error() en *ValError
	fmt.Printf(err.Error())
	fmt.Printf("\n%d", err.Count())
	// Output:
	// Prueba
	//   Número requiered value > 40.000000 [30.000000 received]
	//   Número requiered value > 100.000000 [30.000000 received]
	//   Sol requiered value > 0.000000 [0.000000 received]
	// Total errors: 3
	// 3
}
