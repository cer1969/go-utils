// Copyright Cristian Echeverría Rabí

package values

import (
	"fmt"
	"testing"
)

func TestSCLt(t *testing.T) {
	vc := SimpleChecker()
	vc.Val(30.0).Lt(31.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Lt(30.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Lt(29.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
}

func TestSCLe(t *testing.T) {
	vc := SimpleChecker()
	vc.Val(30.0).Le(31.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Le(30.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Le(29.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
}

func TestSCGt(t *testing.T) {
	vc := SimpleChecker()
	vc.Val(30.0).Gt(29.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Gt(30.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Gt(31.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
}

func TestSCGe(t *testing.T) {
	vc := SimpleChecker()
	vc.Val(30.0).Ge(29.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Ge(30.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Ge(31.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
}

func TestSCMixed(t *testing.T) {
	vc := SimpleChecker()
	vc.Val(35.0).Ge(30.0).Le(40.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Ge(30.0).Le(30.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Ge(31.0).Le(29.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}

	vc = SimpleChecker()
	vc.Val(35.0).Gt(30.0).Lt(40.0)
	if !vc.Ok() {
		t.Error("Ok expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Gt(30.0).Lt(30.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
	vc = SimpleChecker()
	vc.Val(30.0).Gt(31.0).Lt(29.0)
	if vc.Ok() {
		t.Error("Ok not expected")
	}
}

func ExampleSimpleChecker() {
	vc1 := SimpleChecker()
	vc1.Val(30).Gt(40.0).Gt(100.0)
	vc1.Val(0.0).Gt(0.0).Lt(1.0)
	vc2 := SimpleChecker()
	vc2.Val(0.5).Gt(0.0).Lt(1.0)
	fmt.Printf("%v - %v", vc1.Ok(), vc2.Ok())
	// Output:
	// false - true
}
