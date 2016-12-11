// Copyright Cristian Echeverría Rabí

// Common utilities for go programs
package values

import (
	"fmt"
)

//----------------------------------------------------------------------------------------

// ValError clase para identificar Errores en verificación de valores
type ValError struct {
	s string // Error description
	n int    // Error count
}

func (e *ValError) Error() string {
	return e.s
}

func (e *ValError) Count() int {
	return e.n
}

//----------------------------------------------------------------------------------------

// Checker retorna ValChecker, objeto para verificar conjunto de valores
// s: Inicializa Error message que pasará a ValChecker.
//    Puede identificar método y/o objeto a verificar
func Checker(s string) *ValChecker {
	return &ValChecker{msg: s}
}

//----------------------------------------------------------------------------------------

type ValChecker struct {
	msg     string  // Error msg
	i_name  string  // Item name
	i_value float64 // Item value
	n       int     // Error count
}

func (vc *ValChecker) add(sim string, limit float64) {
	vc.n += 1
	vc.msg += fmt.Sprintf("\n  %s requiered value %s %f [%f received]", vc.i_name, sim, limit, vc.i_value)
}

func (vc *ValChecker) Val(name string, val float64) *ValChecker {
	vc.i_name = name
	vc.i_value = val
	return vc
}

func (vc *ValChecker) Lt(limit float64) *ValChecker {
	if vc.i_value >= limit {
		vc.add("<", limit)
	}
	return vc
}

func (vc *ValChecker) Le(limit float64) *ValChecker {
	if vc.i_value > limit {
		vc.add("<=", limit)
	}
	return vc
}

func (vc *ValChecker) Gt(limit float64) *ValChecker {
	if vc.i_value <= limit {
		vc.add(">", limit)
	}
	return vc
}

func (vc *ValChecker) Ge(limit float64) *ValChecker {
	if vc.i_value < limit {
		vc.add(">=", limit)
	}
	return vc
}

func (vc *ValChecker) Error() error {
	if vc.n == 0 {
		return nil
	}
	vc.msg += fmt.Sprintf("\nTotal errors: %d", vc.n)
	return &ValError{s: vc.msg, n: vc.n}
}
