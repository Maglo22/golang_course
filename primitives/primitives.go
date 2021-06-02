package main

import (
	"fmt"
)

func main() {
	booleans()
	integers()
	floatPoints()
	complexNumbers()
	texts()
}

func booleans() {
	// default value is false
	var f bool
	var bt bool = true
	bf := 1 == 2

	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", bt, bt)
	fmt.Printf("%v, %T\n", bf, bf)
}

func integers() {
	// default value is 0
	// sizes:
	// int -> variable size, but min 32 bits
	// signed -> int8, int16, int32, int64
	// unsigned -> uint8, uint16, uint32
	var a int8 = 10
	var b int8 = 3

	// arithmetic operations
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// bit operations
	c := 10 // 1010
	d := 3  // 0011

	// and
	fmt.Println(c & d) // 0010 -> 2
	// or
	fmt.Println(c | d) // 1011 -> 11
	// exclusive
	fmt.Println(c ^ d) // 1001 -> 9
	// neither
	fmt.Println(c &^ d) // 0100 -> 8

	// shifting
	s := 8 // 2^3
	// shift left 3 places
	fmt.Println(s << 3) // 2^3 * 2^3 = 2^6 -> 64
	// shift right 3 places
	fmt.Println(s >> 3) // 2^3 / 2^3 = 2^0 -> 1
}

func floatPoints() {
	// default value is 0
	// sizes:
	// float32, float64
	var f float64 = 3.14
	f = 13.7e72
	f = 2.1e14

	fmt.Printf("%v, %T\n", f, f)

	// arithmetic operations
	a := 10.2
	b := 3.7

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}

func complexNumbers() {
	// default value 0 + 0i
	// sizes:
	// complex64 (float32 values), complex128 (float64 values)
	var c complex64 = 1 + 2i

	fmt.Printf("%v, %T\n", c, c)

	// get only real part
	fmt.Printf("%v, %T\n", real(c), real(c))
	// get only imaginary part
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	// create a complex number
	var real float32 = 2.0
	var imaginary float32 = 10.0

	var com complex64 = complex(real, imaginary)
	fmt.Printf("%v, %T\n", com, com)

	// arithmetic operations
	a := 5 + 2i
	b := 2 + 5i

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}

func texts() {
	// strings
	// UTF8
	// aliases for bytes
	// immutable
	s := "a string"
	fmt.Printf("%v, %T\n", s, s)

	// concatenate
	s2 := "another string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// convert to collection of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	// runes
	// UTF32
	// type aliases for int32
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
