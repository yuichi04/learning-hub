package variables

import "fmt"

// `var`での型指定
var (
	a int = 2 // 明示的に型を指定
	b     = 2 // 自動的に型を推論
)

// 代入値なしの場合は自動的に型に応じたゼロ値が代入される
var (
	l int    // 0
	m string // ""
	n bool   // false
)

// 定数の定義（上書きできない=readonly）
const secret = "secret"

// 定数の複数定義
type Os int // 型の別名定義
const (
	Mac Os = iota // iota: Go言語の組み込み定数。0から始まる連続した整数定数を生成するために使用される
	Windows
	Linux
)

func Variables() {
	// 型推論
	fmt.Println("型推論")
	fmt.Printf("a: %[1]T, b: %[2]T\n", a, b)

	// 短縮記法の型指定
	i := 1      // 自動的に型を推論
	j := int(1) // 明示的に型を指定
	fmt.Println("短縮記法の型指定")
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("j: %v %T\n", j, j)

	fmt.Println("短縮記法の型推論")
	d := 1.23456
	e := "hello"
	f := true
	fmt.Printf("d: %[1]v %[1]T e: %[2]v %[2]T f: %[3]v %[3]T\n", d, e, f)

	// ゼロ値
	fmt.Println("ゼロ値")
	fmt.Printf("l: %[1]v %[1]T m: %[2]v %[2]T n: %[3]v %[3]T\n", l, m, n)

	// 短縮記法の複数変数定義
	pi, title := 3.14, "Go"
	fmt.Println("短縮記法の複数変数定義")
	fmt.Printf("pi: %v title: %v\n", pi, title)

	// 型変換
	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println("型変換")
	fmt.Printf("x: %[1]v %[1]T y: %[2]v %[2]T z: %[3]v %[3]T\n", x, y, z)

	// 定数
	fmt.Println("定数")
	fmt.Printf("secret: %v %T\n", secret, secret)
	fmt.Printf("Mac: %v Windows: %v Linux: %v\n", Mac, Windows, Linux)
}
