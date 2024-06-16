## Go 言語の`switch`文について

### 概要

`switch`文は、複数の条件を簡潔にチェックして処理を分岐させるための構文です。`if-else`文の連鎖を置き換えることができ、コードの可読性を向上させます。Go の`switch`文は、柔軟で強力な制御フロー構造を提供し、シンプルな値の比較だけでなく、式や型に基づいた分岐もサポートします。

### 基本的な構文

`switch`文の基本的な構文は次の通りです。

```go
switch expression {
case value1:
    // value1の場合の処理
case value2:
    // value2の場合の処理
default:
    // いずれのケースにも該当しない場合の処理
}
```

#### サンプルコード

```go
package main

import "fmt"

func main() {
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the week")
    case "Friday":
        fmt.Println("End of the week")
    default:
        fmt.Println("Midweek")
    }
}
```

この例では、変数`day`の値に応じて異なるメッセージが出力されます。

### 式を用いた`switch`

`switch`文の`case`には式を使用することができます。これは特定の値だけでなく、条件を使った分岐を可能にします。

#### サンプルコード

```go
package main

import "fmt"

func main() {
    num := 7
    switch {
    case num < 0:
        fmt.Println("Negative number")
    case num == 0:
        fmt.Println("Zero")
    case num > 0:
        fmt.Println("Positive number")
    }
}
```

この例では、`num`の値に基づいて異なるメッセージが出力されます。

### 初期化付き`switch`

`switch`文には、`if`文と同様に初期化ステートメントを含めることができます。初期化ステートメントは、`switch`ブロック内でのみ有効です。

#### サンプルコード

```go
package main

import "fmt"

func main() {
    switch num := 10; {
    case num < 0:
        fmt.Println("Negative number")
    case num == 0:
        fmt.Println("Zero")
    case num > 0:
        fmt.Println("Positive number")
    }
}
```

この例では、初期化ステートメントで`num`を宣言し、その値に基づいて分岐しています。

### 複数のケース

`case`には複数の値を指定することができます。これにより、同じ処理を複数のケースで実行することができます。

#### サンプルコード

```go
package main

import "fmt"

func main() {
    day := "Saturday"
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
```

この例では、`day`が`"Saturday"`または`"Sunday"`の場合に「Weekend」と出力されます。

### フォールスルー

Go の`switch`文はデフォルトでフォールスルーしません。次の`case`に処理を続けるには、明示的に`fallthrough`キーワードを使用します。

#### サンプルコード

```go
package main

import "fmt"

func main() {
    num := 1
    switch num {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two")
    default:
        fmt.Println("Other")
    }
}
```

この例では、「One」と「Two」の両方が出力されます。

### 型スイッチ

型スイッチを使用すると、インターフェース値の動的型に基づいて処理を分岐させることができます。

#### サンプルコード

```go
package main

import "fmt"

func doSomething(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    doSomething(10)       // Integer: 10
    doSomething("hello")  // String: hello
    doSomething(true)     // Unknown type
}
```

この例では、型スイッチを使ってインターフェース値の型に応じた処理を行っています。

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#switch)
-   [Go by Example: Switch](https://gobyexample.com/switch)
