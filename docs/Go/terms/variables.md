## Go 言語の変数（Variables）について

### 概要

Go 言語の変数は、値を保持するための名前付きストレージです。変数を宣言するには`var`キーワードを使用し、静的型付けの特性を持ちます。型推論もサポートしており、コンパイル時に変数の型が決定されます。

### 変数の宣言と初期化

#### 基本的な宣言方法

```go
var name string
var age int
```

この例では、`name`という文字列型の変数と、`age`という整数型の変数を宣言しています。これらの変数はデフォルト値で初期化されます（`string`は空文字列、`int`は 0）。

#### 宣言と初期化を同時に行う

```go
var name string = "Alice"
var age int = 30
```

この例では、変数を宣言すると同時に初期値を設定しています。

#### 型推論を用いた宣言と初期化

```go
name := "Alice"
age := 30
```

`:=`を使うことで、型を明示的に指定せずに変数を宣言し初期化できます。この方法では、右辺の値から型が自動的に推論されます。

### 複数変数の宣言

#### 一行で複数の変数を宣言

```go
var x, y int
```

#### 一行で複数の変数を宣言と初期化

```go
var a, b int = 1, 2
```

#### 型推論を用いた複数の変数の宣言と初期化

```go
a, b := 1, 2
```

### 変数のゼロ値

Go では、変数を宣言した際に初期値を設定しない場合、各型に応じたゼロ値が設定されます。

-   数値型（`int`, `float`など）：0
-   ブール型：`false`
-   文字列型：空文字列 `""`
-   ポインタ、スライス、マップ、チャネル、インターフェース：`nil`

#### サンプルコード

```go
var x int
var y bool
var z string

fmt.Println(x) // 0
fmt.Println(y) // false
fmt.Println(z) // ""
```

### 定数

Go では、変数とは別に定数を宣言できます。定数は、`const`キーワードを使って宣言し、一度値を設定すると変更できません。

#### サンプルコード

```go
const Pi = 3.14
const Greeting = "Hello, World"
```

### シャドーイング

シャドーイングとは、内側のスコープで宣言された変数が外側のスコープの変数を隠すことを言います。

#### サンプルコード

```go
var x int = 10
if true {
    var x int = 20
    fmt.Println(x) // 20
}
fmt.Println(x) // 10
```

この例では、内側のスコープで宣言された`x`が外側の`x`をシャドーイングしています。

### 複数の戻り値

Go の関数は、複数の戻り値を返すことができます。これを利用して、複数の変数に一度に値を代入できます。

#### サンプルコード

```go
func swap(x, y string) (string, string) {
    return y, x
}

a, b := swap("hello", "world")
fmt.Println(a, b) // "world hello"
```

### ブランク識別子

ブランク識別子 `_` は、不要な変数を無視するために使用されます。

#### サンプルコード

```go
func getValues() (int, int) {
    return 1, 2
}

a, _ := getValues()
fmt.Println(a) // 1
```

### 変数のスコープ

変数のスコープは、宣言された場所によって決まります。グローバルスコープ、パッケージスコープ、関数スコープ、ブロックスコープがあります。

#### サンプルコード

```go
var globalVar = "I am global"

func main() {
    var localVar = "I am local"
    fmt.Println(globalVar) // I am global
    fmt.Println(localVar)  // I am local
}
```

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#variables)
-   [Go by Example: Variables](https://gobyexample.com/variables)
