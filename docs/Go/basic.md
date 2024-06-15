# Golang の基本文法

## 1. 変数

データを保存し、プログラム内で操作するための基本的な方法です。

```go
package main

import "fmt"

func main() {
    var a int = 10 // 明示的な型指定
    var b = 20     // 型推論
    c := 30        // 短縮変数宣言

    fmt.Println(a, b, c)
}

// まとめて宣言する場合
var (
    a int
    b string
    c float64
)
```

## 2. 定数

物理定数や設定値など、変更されない固定値を定義するために使用します。

```go
package main

import "fmt"

const Pi = 3.14

func main() {
    fmt.Println(Pi)
}
```

## 3. 基本的なデータ型

様々な種類のデータを表現するために使用します。<br>
Go の基本的なデータ型には、整数、浮動小数点数、文字列、ブール型があります。

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = 3.14
    var s string = "Hello"
    var b bool = true

    fmt.Println(i, f, s, b)
}
```

## 4. 配列

固定長のデータセットを扱う場合に使用します。

```go
package main

import "fmt"

func main() {
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr)
}
```

## 5. スライス

リストやダイナミックな配列など、可変長のデータセットを扱う場合に使用します。

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)
}
```

## 6. マップ

辞書やハッシュマップなど、キーと値のペアを管理するために使用します。

```go
package main

import "fmt"

func main() {
    m := map[string]int{"one": 1, "two": 2}
    fmt.Println(m)
}
```

## 7. 条件分岐

条件に基づいて異なる処理を実行する場合に使用します。

```go
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}
```

## 8. ループ

リストの各要素を処理する場合など、同じ処理を繰り返し実行する場合に使用します。

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

## 9. 関数

再利用可能なコードブロックを定義する場合に使用します。例えば、特定の処理を行うコードを関数として定義しておくと便利です。

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(2, 3)
    fmt.Println(result)
}
```

## 10. 構造体

ユーザー情報など、複数のデータフィールドを一つのエンティティとして扱う場合に使用します。

```go
package main

import "fmt"

type User struct {
    Name  string
    Email string
}

func main() {
    user := User{Name: "John Doe", Email: "john@example.com"}
    fmt.Println(user)
}
```

## 11. メソッド

ユーザー情報の表示など、構造体に関連する機能を定義する場合に使用します。

```go
package main

import "fmt"

type User struct {
    Name  string
    Email string
}

func (u User) GetDetails() string {
    return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

func main() {
    user := User{Name: "John Doe", Email: "john@example.com"}
    fmt.Println(user.GetDetails())
}
```

## 12. インターフェース

異なる種類のエンティティに共通のメソッドを提供する場合など、異なる型の間で共通の動作を定義する場合に使用します。

```go
package main

import "fmt"

type Describer interface {
    Describe() string
}

type User struct {
    Name  string
    Email string
}

func (u User) Describe() string {
    return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

func main() {
    var d Describer
    user := User{Name: "John Doe", Email: "john@example.com"}
    d = user
    fmt.Println(d.Describe())
}
```

## 13. エラーハンドリング

ファイルの読み込み時にエラーが発生した場合など、エラーが発生した場合にそのエラーを処理するために使用します。

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

## 14. ゴルーチンとチャンネル

複数のタスクを同時に実行する場合など、並行処理を実現するために使用します。

### ゴルーチン

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

### チャンネルの使用

チャンネルを使ってゴルーチン間でデータをやり取りする場合。

```go
package main

import (
    "fmt"
)

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total
}

func main() {
    a := []int{1, 2, 3, 4, 5}
    c := make(chan int)
    go sum(a, c)
    result := <-c
    fmt.Println(result)
}
```
