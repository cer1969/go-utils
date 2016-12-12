// Copyright Cristian Echeverría Rabí

package values

//----------------------------------------------------------------------------------------

// Checker retorna ValChecker, objeto para verificar conjunto de valores
// s: Inicializa Error message que pasará a ValChecker.
//    Puede identificar método y/o objeto a verificar
func SimpleChecker() *SimpleValChecker {
	return &SimpleValChecker{0.0, 0}
}

func Check(val float64) *SimpleValChecker {
	return &SimpleValChecker{val, 0}
}

//----------------------------------------------------------------------------------------

type SimpleValChecker struct {
	i_value float64 // Item value
	n       int     // Error count
}

func (vc *SimpleValChecker) Val(val float64) *SimpleValChecker {
	vc.i_value = val
	return vc
}

func (vc *SimpleValChecker) Lt(limit float64) *SimpleValChecker {
	if vc.i_value >= limit {
		vc.n += 1
	}
	return vc
}

func (vc *SimpleValChecker) Le(limit float64) *SimpleValChecker {
	if vc.i_value > limit {
		vc.n += 1
	}
	return vc
}

func (vc *SimpleValChecker) Gt(limit float64) *SimpleValChecker {
	if vc.i_value <= limit {
		vc.n += 1
	}
	return vc
}

func (vc *SimpleValChecker) Ge(limit float64) *SimpleValChecker {
	if vc.i_value < limit {
		vc.n += 1
	}
	return vc
}

func (vc *SimpleValChecker) Ok() bool {
	return vc.n == 0
}
