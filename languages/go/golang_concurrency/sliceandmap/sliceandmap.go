package sliceandmap

import "fmt"

func SliceAndMap() {
	// 配列
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	a3 := [...]int{10, 20} // [...]: 要素数を自動的に特定する。`:=`の場合は必ず`{}`で要素を指定する必要がある
	fmt.Println(a1)        // [0 0 0]
	fmt.Println(a2)        // [10 20 30]
	fmt.Println(a3)        // [10 20]

	// len: 配列、スライス、またはマップの現在の要素数を返す
	// cap: 配列やスライスが格納できる要素の最大数を返す
	fmt.Printf("%v %v\n", len(a3), cap(a3)) // 2 2
	fmt.Printf("%T %T\n", a2, a3)           // [3]int [2]int

	// [3]intと[2]intは別の型として認識されるため、a2にa3を代入することはできない
	// 配列は要素数を動的に増やすことができないため、より柔軟な操作をしたい場合はスライスを使用する

	// スライス
	var s1 []int
	s2 := []int{}
	fmt.Printf("s1: type=%[1]T val=%[1]v len=%v cap=%v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: type=%[1]T val=%[1]v len=%v cap=%v\n", s2, len(s2), cap(s2))

	// varで初期化した場合はnilとして認識される。{}で初期化した場合はnilと認識されない
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// スライスに対する要素の追加
	s1 = append(s1, 1, 2, 3) // append()で追加できる
	fmt.Printf("s1: type=%[1]T val=%[1]v len=%v cap=%v\n", s1, len(s1), cap(s1))
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...) // 別のスライスを追加する方法（...を変数の末尾につけることで可変長の引数として渡すことができる）
	fmt.Printf("s1: type=%[1]T val=%[1]v len=%v cap=%v\n", s1, len(s1), cap(s1))

	// makeを使ったスライスの定義方法
	// make([]int, 長さ、キャパシティ)
	s4 := make([]int, 0, 2)
	fmt.Printf("s4: type=%[1]T val=%[1]v len=%v cap=%v\n", s4, len(s4), cap(s4))
	s4 = append(s4, 1, 2, 3, 4) // append関数は予め決められていたスライスのキャパシティを超えて要素を追加することができる
	fmt.Printf("s4: type=%[1]T val=%[1]v len=%v cap=%v\n", s4, len(s4), cap(s4))

	// 指定したインデックスの要素を変更する方法
	s5 := make([]int, 4, 6)
	fmt.Printf("s5: val=%v len=%v cap=%v\n", s5, len(s5), cap(s5))

	// スライスを切り取って使用した場合、メモリを共有する仕様になっている
	s6 := s5[1:3] // インデックス番号2と3の要素を代入する
	s6[1] = 10
	fmt.Printf("s5: val=%v len=%v cap=%v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: val=%v len=%v cap=%v\n", s6, len(s6), cap(s6))
	s6 = append(s6, 2)
	fmt.Printf("s5: val=%v len=%v cap=%v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: val=%v len=%v cap=%v\n", s6, len(s6), cap(s6))

	// メモリを共有したくない場合は`copy()`を使う
	sc6 := make([]int, len(s5[1:3]))
	fmt.Printf("s5 source of copy: val=%v len=%v cap=%v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6 dst copy before: val=%v len=%v cap=%v\n", sc6, len(sc6), cap(sc6))
	copy(sc6, s5[1:3])
	fmt.Printf("sc6 dst of copy after: val=%v len=%v cap=%v\n", sc6, len(sc6), cap(sc6))
	sc6[1] = 12
	fmt.Printf("s5: val=%v len=%v cap=%v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6: val=%v len=%v cap=%v\n", sc6, len(sc6), cap(sc6))
}
