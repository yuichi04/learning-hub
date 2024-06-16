package main

import (
	"fmt"
	"golang_concurrency/greeting"
	"golang_concurrency/moduleandpackage"
	"golang_concurrency/pointershadowing"
	"golang_concurrency/sliceandmap"
	"golang_concurrency/variables"
)

func main() {
	greeting.Greeting()

	fmt.Println("----------------------------------------")

	moduleandpackage.GoDotEnv()
	fmt.Println(moduleandpackage.Offset)
	fmt.Println(moduleandpackage.Sum(1, 2))
	fmt.Println(moduleandpackage.Multiply(1, 2))

	fmt.Println("----------------------------------------")

	variables.Variables()

	fmt.Println("----------------------------------------")

	pointershadowing.PointerShadowing()

	fmt.Println("----------------------------------------")

	sliceandmap.SliceAndMap()
}
