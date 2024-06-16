## Go 言語の関数（Function）について

### 概要

関数は、特定のタスクを実行するために一連の命令をまとめたものです。関数を使用すると、コードを再利用しやすくし、プログラムの可読性と保守性を向上させることができます。Go 言語の関数は、宣言、呼び出し、引数、戻り値、匿名関数、クロージャ、可変長引数などをサポートします。

### 関数の宣言

関数は`func`キーワードを使って宣言します。関数名、引数リスト、戻り値の型、関数本体を含みます。

#### サンプルコード

```go
func add(a int, b int) int {
    return a + b
}
```

この例では、`add`という名前の関数を宣言し、2 つの整数引数を取り、その和を整数として返します。

### 関数の呼び出し

宣言された関数は、名前と引数を使って呼び出します。

#### サンプルコード

```go
result := add(3, 5)
fmt.Println(result) // 8
```

### 複数の戻り値

Go の関数は、複数の値を返すことができます。これはエラーハンドリングや追加の情報を返す際に非常に便利です。

#### サンプルコード

```go
func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result) // 5
}
```

### 名前付き戻り値

戻り値に名前を付けることで、戻り値を初期化し、関数内で使用することができます。`return`ステートメントのみで戻り値を返すことができます。

#### サンプルコード

```go
func swap(a, b int) (x, y int) {
    x = b
    y = a
    return
}

x, y := swap(1, 2)
fmt.Println(x, y) // 2 1
```

### 可変長引数

可変長引数を使用すると、任意の数の引数を関数に渡すことができます。可変長引数は、引数リストの最後に配置されます。

#### サンプルコード

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

result := sum(1, 2, 3, 4)
fmt.Println(result) // 10
```

### 関数リテラル（匿名関数）

Go では、関数リテラル（匿名関数）を使用して関数を変数に代入したり、他の関数に引数として渡すことができます。

#### サンプルコード

```go
add := func(a, b int) int {
    return a + b
}

fmt.Println(add(3, 4)) // 7
```

### クロージャ

クロージャは、関数リテラルがその定義された環境を保持する機能です。クロージャは、外部の変数にアクセスし、それを変更することができます。

#### サンプルコード

```go
func incrementer() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

inc := incrementer()
fmt.Println(inc()) // 1
fmt.Println(inc()) // 2
```

### デファード関数呼び出し

`defer`キーワードを使って、関数の終了時に特定の関数を呼び出すことができます。これは、リソースの解放やクリーンアップ処理に役立ちます。

#### サンプルコード

```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}

// 出力結果:
// Hello
// World
```

### 再帰関数

関数は自分自身を呼び出すことができます。これは再帰と呼ばれ、特定の問題を解決するのに便利です。

#### サンプルコード

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5)) // 120
```

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#functions)
-   [Go by Example: Functions](https://gobyexample.com/functions)
