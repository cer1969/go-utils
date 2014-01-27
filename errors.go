// Copyright Cristian Echeverr�a Rab�

// Common utilities for go programs
package utils

type ValueError struct {
    s string
    n int
}

func (e *ValueError) Error() string {
    return e.s + " " + e.n
}