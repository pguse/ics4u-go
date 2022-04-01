package main

import "fmt"

type Fraction struct {
	Num int
	Den int
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d / %d", f.Num, f.Den)
}
func (f Fraction) Mult(f1 Fraction) Fraction {
	return Fraction{}
}
func (f Fraction) Div(f1 Fraction) Fraction {
	return Fraction{}
}
func (f Fraction) Add(f1 Fraction) Fraction {
	return Fraction{}
}

func (f Fraction) Subt(f1 Fraction) Fraction {
	return Fraction{}
}

func main() {
	f1 := Fraction{2, 3}
	f2 := Fraction{4, 5}
	f3 := f1.Mult(f2)

	fmt.Println(f1.String())
	fmt.Println(f3.String())
	// Using the new methods
	fmt.Println(f1.Div(f2))
	fmt.Println(f1.Add(f2))
	fmt.Println(f1.Subt(f2))
}
