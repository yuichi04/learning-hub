## Go 言語の`for`文について

### 概要

`for`文は、繰り返し処理を行うための構文です。Go 言語の`for`文は、他の言語の`while`文や`do-while`文に相当する役割も果たし、シンプルで強力な繰り返し制御を提供します。

### 基本的な構文

Go 言語の`for`文の基本的な構文は次の通りです。

```go
for initialization; condition; post {
    // 繰り返し実行されるコード
}
```

#### サンプルコード

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

この例では、`i`が 0 から 4 までの値を持つ間、`fmt.Println(i)`が繰り返し実行されます。

### 条件のみの`for`文

`for`文は条件のみで使用することもできます。この形は他の言語の`while`文に相当します。

```go
for condition {
    // 繰り返し実行されるコード
}
```

#### サンプルコード

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

この例では、`i`が 5 未満である間、`fmt.Println(i)`が繰り返し実行されます。

### 無限ループ

`for`文の条件を省略すると無限ループになります。

```go
for {
    // 無限に繰り返されるコード
}
```

#### サンプルコード

```go
package main

import "fmt"

func main() {
    i := 0
    for {
        fmt.Println(i)
        i++
        if i >= 5 {
            break
        }
    }
}
```

この例では、`i`が 5 に達するまで無限ループが続きます。`break`文を使ってループを終了しています。

### 範囲（range）を使った`for`文

`range`キーワードを使うと、スライス、配列、マップ、チャネル、文字列などのデータ構造を反復処理できます。

#### スライスや配列の場合

```go
package main

import "fmt"

func main() {
    nums := []int{2, 3, 4}
    for i, num := range nums {
        fmt.Println(i, num)
    }
}
```

この例では、スライス`nums`の各要素を反復処理し、そのインデックスと値を出力します。

#### マップの場合

```go
package main

import "fmt"

func main() {
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Println(k, v)
    }
}
```

この例では、マップ`kvs`の各キーと値を反復処理します。

#### 文字列の場合

```go
package main

import "fmt"

func main() {
    str := "hello"
    for i, c := range str {
        fmt.Printf("%d: %c\n", i, c)
    }
}
```

この例では、文字列`str`の各文字を反復処理し、そのインデックスと文字を出力します。

### `for`文の制御

#### `break`文

`break`文を使ってループを中断し、即座にループの外に出ます。

#### サンプルコード

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println(i)
    }
}
```

この例では、`i`が 5 になるとループが中断されます。

#### `continue`文

`continue`文を使って現在の反復をスキップし、次の反復に進みます。

#### サンプルコード

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println(i)
    }
}
```

この例では、`i`が偶数の場合に現在の反復がスキップされ、`i`が奇数の場合のみ値が出力されます。

### ネストされた`for`文

`for`文はネストすることができます。これにより、二重ループや多重ループを実現できます。

#### サンプルコード

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }
}
```

この例では、二重ループを使って`i`と`j`の組み合わせを出力しています。

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#for)
-   [Go by Example: For](https://gobyexample.com/for)
