/*
 * matrixhelp.go, part of gochem.
 *
 *
 * Copyright 2012 Raul Mera <rmera{at}chemDOThelsinkiDOTfi>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as
 * published by the Free Software Foundation; either version 2.1 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General
 * Public License along with this program.  If not, see
 * <http://www.gnu.org/licenses/>.
 *
 *
 * Gochem is developed at the laboratory for instruction in Swedish, Department of Chemistry,
 * University of Helsinki, Finland.
 *
 *
 */

package chem

//A munch of unexported mathematical functions, most of them just for convenience.

import (
	"fmt"
	"math"

	"github.com/rmera/gochem/v3"
	"gonum.org/v1/gonum/mat"
)

const appzero float64 = 0.000000000001 //used to correct floating point
//errors. Everything equal or less than this is considered zero. This probably sucks.

//Computes the inverse matrix of F and puts it in target. If target is nil, it alloactes
//a new one. Returns target.
func gnInverse(F, target *v3.Matrix) (*v3.Matrix, error) {
	if target == nil {
		r := F.NVecs()
		target = v3.Zeros(r)
	}
	err := target.Dense.Inverse(F.Dense)
	if err != nil {
		err = CError{err.Error(), []string{"mat.Inverse", "gnInverse"}}
	}
	return target, err
}

//cross Takes 2 3-len column or row vectors and returns a column or a row
//vector, respectively, with the Cross product of them.
//should panic
func cross(a, b *v3.Matrix) *v3.Matrix {
	c := v3.Zeros(1)
	c.Cross(a, b)
	return c
}

//invSqrt return the inverse of the square root of val, or zero if
//|val|<appzero. Returns -1 if val is negative
func invSqrt(val float64) float64 {
	if math.Abs(val) <= appzero { //Not sure if need the
		return 0
	} else if val < 0 { //negative
		panic("attempted to get the square root of a negative number")
	}
	return 1.0 / math.Sqrt(val)
}

//Returns and empty, but not nil, Dense. It barely allocates memory
func emptyDense() (*mat.Dense, error) {
	a := make([]float64, 0, 0)
	return mat.NewDense(0, 0, a), nil

}

//Returns an zero-filled Dense with the given dimensions
//It is to be substituted by the Gonum function.
func gnZeros(r, c int) *mat.Dense {
	f := make([]float64, r*c, r*c)
	return mat.NewDense(r, c, f)

}

//Returns an identity matrix spanning span cols and rows
func gnEye(span int) *mat.Dense {
	A := gnZeros(span, span)
	for i := 0; i < span; i++ {
		A.Set(i, i, 1.0)
	}
	return A
}

//returns a rows,cols matrix filled with gnOnes.
func gnOnes(rows, cols int) *mat.Dense {
	gnOnes := gnZeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			gnOnes.Set(i, j, 1)
		}
	}
	return gnOnes
}

//The 2 following functions may not even be used.
func gnMul(A, B mat.Matrix) *mat.Dense {
	ar, _ := A.Dims()
	_, bc := B.Dims()
	C := gnZeros(ar, bc)
	C.Mul(A, B)
	return C
}

func gnCopy(A mat.Matrix) *mat.Dense {
	r, c := A.Dims()
	B := gnZeros(r, c)
	B.Copy(A)
	return B
}

func gnT(A mat.Matrix) *mat.Dense {
	r, c := A.Dims()
	B := gnZeros(c, r)
	B.Copy(A.T())
	return B
}

//This is a temporal function. It returns the determinant of a 3x3 matrix. Panics if the matrix is not 3x3
func det(A mat.Matrix) float64 {
	r, c := A.Dims()
	if r != 3 || c != 3 {
		panic("Determinants are for now only available for 3x3 matrices")
	}
	return (A.At(0, 0)*(A.At(1, 1)*A.At(2, 2)-A.At(2, 1)*A.At(1, 2)) - A.At(1, 0)*(A.At(0, 1)*A.At(2, 2)-A.At(2, 1)*A.At(0, 2)) + A.At(2, 0)*(A.At(0, 1)*A.At(1, 2)-A.At(1, 1)*A.At(0, 2)))
}

/**These are from the current proposal for gonum, by Dan Kortschak. It will be taken out
 * from here when gonum is implemented. The gn prefix is appended to the names to make them
 * unimported and to allow easy use of search/replace to add the "num" prefix when I change to
 * gonum.**/

// A gnPanicker is a function that may panic.
type gnPanicker func()

// Maybe will recover a panic with a type matrix.Error or a VecError from fn, and return this error.
// Any other error is re-panicked. It is a small modification
//Maybe this funciton should be exported.
func gnMaybe(fn gnPanicker) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case v3.Error:
				err = e
			case mat.Error:
				err = CError{fmt.Sprintf("goChem: Error in gonum function: %s", e), []string{"gnMaybe"}}
			default:
				panic(r)
			}
		}
	}()
	fn()
	return err
}

//Puts A**exp on the receiver, in a pretty naive way.
func pow(A mat.Matrix, F *mat.Dense, exp float64) {
	ar, ac := A.Dims()
	fr, fc := F.Dims()
	if ar != fr || ac != fc {
		panic(mat.ErrShape)
	}
	for i := 0; i < ar; i++ {
		for j := 0; j < ac; j++ {
			F.Set(i, j, math.Pow(A.At(i, j), exp))
		}

	}
}
