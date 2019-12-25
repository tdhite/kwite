/*
math.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package math

import (
	"math"

	"github.com/tdhite/kwite/internal/tplfunc/funcs"
)

func E() float64 {
	return math.E
}

func Pi() float64 {
	return math.Pi
}

func Phi() float64 {
	return math.Phi
}

func Sqrt2() float64 {
	return math.Sqrt2
}

func SqrtE() float64 {
	return math.SqrtE
}

func SqrtPi() float64 {
	return math.SqrtPi
}

func SqrtPhi() float64 {
	return math.SqrtPhi
}

func Ln2() float64 {
	return math.Ln2
}

func Log2E() float64 {
	return math.Log2E
}

func Ln10() float64 {
	return math.Ln10
}

func Log10E() float64 {
	return math.Log10E
}

// Adds all methods from the math package
func LoadFuncs() error {
	f := map[string]interface{}{
		// constants
		"E":       E,
		"Pi":      Pi,
		"Phi":     Phi,
		"Sqrt2":   Sqrt2,
		"SqrtE":   SqrtE,
		"SqrtPi":  SqrtPi,
		"SqrtPhi": SqrtPhi,
		"Ln2":     Ln2,
		"Log2E":   Log2E,
		"Ln10":    Ln10,
		"Log10E":  Log10E,

		// imported funcs
		"Abs":             math.Abs,
		"Acos":            math.Acos,
		"Acosh":           math.Acosh,
		"Asin":            math.Asin,
		"Asinh":           math.Asinh,
		"Atan":            math.Atan,
		"Atan2":           math.Atan2,
		"Atanh":           math.Atanh,
		"Cbrt":            math.Cbrt,
		"Ceil":            math.Ceil,
		"Copysign":        math.Copysign,
		"Cos":             math.Cos,
		"Cosh":            math.Cosh,
		"Dim":             math.Dim,
		"Erf":             math.Erf,
		"Erfc":            math.Erfc,
		"Erfcinv":         math.Erfcinv,
		"Erfinv":          math.Erfinv,
		"Exp":             math.Exp,
		"Exp2":            math.Exp2,
		"Expm1":           math.Expm1,
		"Float32bits":     math.Float32bits,
		"Float32frombits": math.Float32frombits,
		"Float64bits":     math.Float64bits,
		"Float64frombits": math.Float64frombits,
		"Floor":           math.Floor,
		//"Frexp":           math.Frexp,
		"Gamma": math.Gamma,
		"Hypot": math.Hypot,
		"Ilogb": math.Ilogb,
		"Inf":   math.Inf,
		"IsInf": math.IsInf,
		"IsNaN": math.IsNaN,
		"J0":    math.J0,
		"J1":    math.J1,
		"Jn":    math.Jn,
		"Ldexp": math.Ldexp,
		//		"Lgamma":          math.Lgamma,
		"Log":   math.Log,
		"Log10": math.Log10,
		"Log1p": math.Log1p,
		"Log2":  math.Log2,
		"Logb":  math.Logb,
		"Max":   math.Max,
		"Min":   math.Min,
		"Mod":   math.Mod,
		//"Modf":        math.Modf,
		"NaN":         math.NaN,
		"Nextafter":   math.Nextafter,
		"Nextafter32": math.Nextafter32,
		"Pow":         math.Pow,
		"Pow10":       math.Pow10,
		"Remainder":   math.Remainder,
		"Round":       math.Round,
		"RoundToEven": math.RoundToEven,
		"Signbit":     math.Signbit,
		"Sin":         math.Sin,
		//"Sincos":      math.Sincos,
		"Sinh":  math.Sinh,
		"Sqrt":  math.Sqrt,
		"Tan":   math.Tan,
		"Tanh":  math.Tanh,
		"Trunc": math.Trunc,
		"Y0":    math.Y0,
		"Y1":    math.Y1,
		"Yn":    math.Yn,
	}

	return funcs.AddMethods(f)
}
