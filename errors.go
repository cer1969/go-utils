// Copyright Cristian Echeverría Rabí

// Common utilities for go programs
package utils

import (
	"fmt"
)

//----------------------------------------------------------------------------------------

type ValueError struct {
	s string
	n int
}

func (e *ValueError) Error() string {
	return e.s
}

func (e *ValueError) Count() int {
	return e.n
}

//----------------------------------------------------------------------------------------

func NewValueChecker(fname string) *ValueChecker {
	return &ValueChecker{f: fname}
}

//----------------------------------------------------------------------------------------

type ValueChecker struct {
	f string
	s string
	n int
}

func (ts *ValueChecker) add(ref string, sim string, val float64, limit float64) {
	ts.n += 1
	ts.s += fmt.Sprintf(" %s requiered value %s %f [%f received];", ref, sim, limit, val)
}

func (ts *ValueChecker) Lt(ref string, val float64, limit float64) {
	if val >= limit {
		ts.add(ref, "<", val, limit)
	}
}

func (ts *ValueChecker) Le(ref string, val float64, limit float64) {
	if val > limit {
		ts.add(ref, "<=", val, limit)
	}
}

func (ts *ValueChecker) Gt(ref string, val float64, limit float64) {
	if val <= limit {
		ts.add(ref, ">", val, limit)
	}
}

func (ts *ValueChecker) Ge(ref string, val float64, limit float64) {
	if val < limit {
		ts.add(ref, ">=", val, limit)
	}
}

func (ts *ValueChecker) Error() error {
	if ts.n == 0 {
		return nil
	}
	msg := fmt.Sprintf("%s value error [%d]: %s", ts.f, ts.n, ts.s)
	return &ValueError{s: msg, n: ts.n}
}
