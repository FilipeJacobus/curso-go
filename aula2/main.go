package main

import (
	"fmt"
	"math/rand"
)

var (
	variavel1 = 1
	variavel2 = "Texto"
)

const (
	constante1 = 2
)

func main() {
	fmt.Println(rand.Int())
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(constante1)

	resp, str := soma(1, 2)
	fmt.Println(resp, str)
}

func soma(val1, val2 int) (int, string) {

	s := val1 + val2

	return s, "GO!"

}

/*

tipos basicos

inteiros:
	ex: -1,0,1,2,3...

	int
		int8 int16 int32 int64
		rune = int32

	uint
		uint8 uint16 uint32 uint64
		byte = uint8

float
	ex: 1.2345, 4.2, 78.8
	float32 float64

strings
	ex: "Hello world!", "curso GO!"

booleanos
	ex: true false

complexos
	ex: 2+3i
	complex64 complex128


*/
