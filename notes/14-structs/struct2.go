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
	return Fraction{f.Num * f1.Num, f.Den * f1.Den}
}

func main() {
	f1 := Fraction{2, 3}
	f2 := Fraction{4, 5}
	f3 := f1.Mult(f2)

	fmt.Println(f1.String())
	fmt.Println(f3.String())
}
