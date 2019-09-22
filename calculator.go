package calculator

import (
	"errors"
	"math"
	"strconv"
)

// MyNum is our internal representation of a number
type MyNum struct {
	n float64
}

// GetNumFromBin converts string Bin to MyNum
func GetNumFromBin(s string) *MyNum {
	retval := new(MyNum)
	anInt64, _ := strconv.ParseInt(s, 2, 64)
	retval.n = float64(anInt64)
	return retval
}

// GetNumFromOct converts string Oct to MyNum
func GetNumFromOct(s string) *MyNum {
	retval := new(MyNum)
	anInt64, _ := strconv.ParseInt(s, 8, 64)
	retval.n = float64(anInt64)
	return retval
}

// GetNumFromHex converts string Hex to MyNum
func GetNumFromHex(s string) *MyNum {
	retval := new(MyNum)
	anInt64, _ := strconv.ParseInt(s, 16, 64)
	retval.n = float64(anInt64)
	return retval
}

// GetNumFromInt converts int to MyNum
func GetNumFromInt(i int64) *MyNum {
	retval := new(MyNum)
	retval.n = float64(i)
	return retval
}

// Multiply two MyNums together
func (m *MyNum) Multiply(n *MyNum) *MyNum {
	m.n *= n.n
	return m
}

// Add two MyNums together
func (m *MyNum) Add(n *MyNum) *MyNum {
	m.n += n.n
	return m
}

// Subtract a MyNum
func (m *MyNum) Subtract(n *MyNum) *MyNum {
	m.n -= n.n
	return m
}

// Divide a MyNum
func (m *MyNum) Divide(n *MyNum) (*MyNum, error) {
	// What if 'n' is zero??
	if n.n == 0 {
		return m, errors.New("attempt to divide by zero")
	}
	m.n /= n.n
	return m, nil
}

// SqRoot will get square root of this number
func (m *MyNum) SqRoot() *MyNum {
	m.n = math.Sqrt(m.n)
	return m
}

// Exponent this number
func (m *MyNum) Exponent(n *MyNum) *MyNum {
	m.n = math.Pow(m.n, n.n)
	return m
}
