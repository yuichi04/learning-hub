## Go 言語のジェネリクス（Generics）について

### 概要

ジェネリクスは、Go 1.18 で導入された機能で、型に依存しない関数やデータ構造を定義することができます。これにより、コードの再利用性が向上し、異なる型に対して同じロジックを適用できる柔軟性が提供されます。

### ジェネリック関数の定義

ジェネリック関数は、型パラメータを使用して定義します。型パラメータは関数名の後に角括弧`[]`で囲んで指定します。

#### サンプルコード

```go
package main

import "fmt"

// ジェネリック関数の定義
func Print[T any](value T) {
    fmt.Println(value)
}

func main() {
    Print(42)      // 42
    Print("Hello") // Hello
    Print(3.14)    // 3.14
}
```

この例では、`Print`関数が型パラメータ`T`を受け取り、任意の型の値を受け入れます。

### ジェネリック型の定義

ジェネリック型を定義することもできます。型パラメータを使用して、汎用的なデータ構造を作成できます。

#### サンプルコード

```go
package main

import "fmt"

// ジェネリック型の定義
type Box[T any] struct {
    Value T
}

func main() {
    intBox := Box[int]{Value: 123}
    strBox := Box[string]{Value: "Hello"}

    fmt.Println(intBox.Value) // 123
    fmt.Println(strBox.Value) // Hello
}
```

この例では、`Box`というジェネリック型を定義し、任意の型の値を保持することができます。

### 複数の型パラメータ

複数の型パラメータを使用することもできます。型パラメータはコンマで区切ります。

#### サンプルコード

```go
package main

import "fmt"

// 複数の型パラメータを持つ関数
func Pair[A, B any](a A, b B) (A, B) {
    return a, b
}

func main() {
    a, b := Pair(1, "one")
    fmt.Println(a, b) // 1 one
}
```

この例では、`Pair`関数が 2 つの型パラメータ`A`と`B`を受け取り、異なる型の値を返します。

### 型制約

ジェネリクスでは、型制約を使用して型パラメータに特定のメソッドや特性を持つ型を指定できます。型制約はインターフェースを使用して定義されます。

#### サンプルコード

```go
package main

import "fmt"

// 数値型のみを許可するインターフェース
type Number interface {
    int | float64
}

// ジェネリック関数に型制約を適用
func Add[T Number](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Add(3, 4))        // 7
    fmt.Println(Add(3.14, 2.71))  // 5.85
}
```

この例では、`Number`インターフェースを定義し、`Add`関数がこのインターフェースを実装する型のみを受け入れるようにしています。

### 型セット

Go 1.18 では、型制約に型セットを使用することができます。型セットは、特定の型の集合を定義するために使用されます。

#### サンプルコード

```go
package main

import "fmt"

// 数値型の集合を定義
type Number interface {
    int | int64 | float64
}

func Sum[T Number](values []T) T {
    var sum T
    for _, v := range values {
        sum += v
    }
    return sum
}

func main() {
    intValues := []int{1, 2, 3, 4, 5}
    floatValues := []float64{1.1, 2.2, 3.3}

    fmt.Println(Sum(intValues))     // 15
    fmt.Println(Sum(floatValues))   // 6.6
}
```

この例では、`Number`インターフェースに複数の数値型を指定し、`Sum`関数がこれらの型のスライスを受け取って合計を計算します。

### 既存コードのリファクタリング

ジェネリクスを使用して既存の非ジェネリックなコードをリファクタリングすることができます。これにより、コードの再利用性が向上し、重複を減らすことができます。

#### サンプルコード

```go
package main

import "fmt"

// 非ジェネリックな最大値関数
func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// ジェネリックな最大値関数
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(MaxInt(3, 5))         // 5
    fmt.Println(Max(3, 5))            // 5
    fmt.Println(Max("apple", "pear")) // pear
}
```

この例では、`MaxInt`関数をジェネリックな`Max`関数にリファクタリングし、任意の型の最大値を求めることができるようにしています。

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/go1.18#generics)
-   [Go Blog: Generics in Go](https://blog.golang.org/generics-next-step)
-   [Go by Example: Generics](https://gobyexample.com/generics)
