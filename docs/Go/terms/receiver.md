## Go 言語のレシーバ（Receiver）について

### 概要

レシーバは、Go 言語のメソッドが所属する型のインスタンスを指します。レシーバを使用することで、特定の型に関連する関数を定義できます。メソッドは、構造体やカスタム型に対して定義されることが一般的です。

### メソッドの定義

メソッドは、関数定義にレシーバを追加することで定義されます。レシーバは、メソッドの対象となる型を指定します。

#### サンプルコード

```go
package main

import "fmt"

type Rectangle struct {
    Width, Height int
}

// メソッドの定義
func (r Rectangle) Area() int {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area()) // Area: 50
}
```

この例では、`Rectangle`型に`Area`メソッドを定義しています。

### ポインタレシーバと値レシーバ

レシーバには、値レシーバとポインタレシーバの 2 種類があります。

#### 値レシーバ

値レシーバはメソッドが呼び出されたときにその値のコピーを受け取ります。そのため、メソッド内でフィールドを変更しても、元のインスタンスには影響を与えません。

```go
func (r Rectangle) Scale(factor int) {
    r.Width *= factor
    r.Height *= factor
}
```

#### ポインタレシーバ

ポインタレシーバはメソッドが呼び出されたときにそのインスタンスのポインタを受け取ります。そのため、メソッド内でフィールドを変更すると、元のインスタンスにも影響を与えます。

```go
func (r *Rectangle) Scale(factor int) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    rect.Scale(2)
    fmt.Println("Scaled Width:", rect.Width) // Scaled Width: 20
    fmt.Println("Scaled Height:", rect.Height) // Scaled Height: 10
}
```

### 値レシーバとポインタレシーバの選択基準

-   **値レシーバを選ぶ場合**

    -   レシーバが小さな構造体で、コピーのコストが低い場合。
    -   メソッドがレシーバの状態を変更しない場合（イミュータブルな操作）。

-   **ポインタレシーバを選ぶ場合**
    -   レシーバが大きな構造体で、コピーのコストが高い場合。
    -   メソッドがレシーバの状態を変更する場合。
    -   一貫性のため、ほとんどのメソッドがポインタレシーバを使っている場合。

### レシーバの型として使用できるもの

レシーバとして使用できるのは、次の 2 つです。

1. **構造体**

    - 構造体は最も一般的なレシーバです。

2. **カスタム型**
    - 任意のカスタム型もレシーバとして使用できます。

#### サンプルコード（カスタム型）

```go
type MyInt int

func (m MyInt) IsPositive() bool {
    return m > 0
}

func main() {
    var num MyInt = 10
    fmt.Println(num.IsPositive()) // true
}
```

### メソッドチェーン

メソッドチェーンを使うことで、複数のメソッド呼び出しを連続して行うことができます。ポインタレシーバを使うと、元のインスタンスを変更した上でメソッドチェーンを実現できます。

#### サンプルコード

```go
func (r *Rectangle) SetWidth(width int) *Rectangle {
    r.Width = width
    return r
}

func (r *Rectangle) SetHeight(height int) *Rectangle {
    r.Height = height
    return r
}

func main() {
    rect := &Rectangle{}
    rect.SetWidth(10).SetHeight(20)
    fmt.Println("Width:", rect.Width)   // Width: 10
    fmt.Println("Height:", rect.Height) // Height: 20
}
```

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#methods)
-   [Go by Example: Methods](https://gobyexample.com/methods)
