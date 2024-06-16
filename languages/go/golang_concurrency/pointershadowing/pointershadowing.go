package pointershadowing

import (
	"fmt"
	"unsafe"
)

func PointerShadowing() {
	// Pointer
	// ----------------------------------------
	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)

	var ui2 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui2)

	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1 // p1: ui1のメモリアドレスを保持
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1)) // ポインタのバイト数の確認
	fmt.Printf("memory address of p1: %p\n", &p1)

	// デリファレンス（ポインタが指す値を取得できる）
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	// ダブルポインタ
	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1)) // ポインタのバイト数の確認
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)

	// Shadowing
	// ----------------------------------------
	ok, result := true, "A"
	if ok {
		result := "B" // ここで定義しなおしているため、if文内が変数のスコープになる
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)
}
