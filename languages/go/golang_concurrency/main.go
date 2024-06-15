package main

import (
	"fmt"
	"golang_concurrency/calculator"
	"os"

	"github.com/joho/godotenv"
)

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

// 代入値なしの場合は自動的にそれぞれの型で設定された値が代入される
var (
	l int    // 0
	m string // ""
	n bool   // false
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))

	fmt.Println("----------------------------------------")

	i := 1
	ui := uint16(2)
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.23456
	s := "hello"
	b := true
	fmt.Printf("f: %v %T\n", f, f)
	fmt.Printf("s: %v %T\n", s, s)
	fmt.Printf("d: %v %T\n", b, b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

	i = 2
	fmt.Printf("i: %v\n", i)

	fmt.Printf("l:%v m:%v n:%v\n", l, m, n)

	fmt.Println("----------------------------------------")

}
