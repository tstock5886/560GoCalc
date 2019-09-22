package calculator

import (
	"testing"
)

func TestConv(t *testing.T) {
	a := GetNumFromBin("101011")
	if a.n != 43 {
		t.Errorf("Binary 101011 should be 43, but we have %f\n", a.n)
	}
	a = GetNumFromHex("0A")
	if a.n != 10 {
		t.Errorf("Hex 0A should be 10, but we have %f\n", a.n)
	}
	b := GetNumFromOct("12")
	if b.n != 10 {
		t.Errorf("Oct 11 should be 10, but we have %f\n", b.n)
	}
	c := GetNumFromInt(100)
	if c.n != 100 {
		t.Errorf("Int 100 should be 100, but we have %f\n", c.n)
	}
	a.Add(b)
	if a.n != 20 {
		t.Errorf("Add 10 to 10 should be 20, but we have %f\n", a.n)
	}
	a.Subtract(b)
	if a.n != 10 {
		t.Errorf("Sub 10 from 20 should be 10, but we have %f\n", a.n)
	}
	a.Multiply(b)
	if a.n != 100 {
		t.Errorf("Multiply 10 and 10 should be 100, but we have %f\n", a.n)
	}
	a.Divide(b)
	if a.n != 10 {
		t.Errorf("Divide 100 by 10 should be 10, but we have %f\n", a.n)
	}
	zero := GetNumFromInt(0)
	_, err := a.Divide(zero)
	if err == nil {
		t.Errorf("Divide by zero should throw an error\n")
	}
	five := GetNumFromInt(5)
	a.Exponent(five)
	if a.n != 100000 {
		t.Errorf("The 10^5 should be 100000. We got %f\n", a.n)
	}

	oneHundred := GetNumFromInt(100)
	oneHundred.SqRoot()
	if oneHundred.n != 10 {
		t.Errorf("The sqroot of 100 should be 10. We got %f\n", oneHundred.n)
	}

}
