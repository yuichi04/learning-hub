# Go 言語の基礎情報

## 特徴

Go 言語（Golang）は、Google によって開発されたオープンソースのプログラミング言語です。2007 年にロバート・グリースマー、ロブ・パイク、ケン・トンプソンによって設計され、2009 年に最初のバージョンがリリースされました。Go 言語の主な特徴は以下の通りです。

1. **シンプルで効率的な構文**

    - シンプルな文法と静的型付けにより、コードの可読性と保守性が高い。
    - コンパイル速度が速く、効率的なバイナリを生成。

2. **並行処理のサポート**

    - Go 言語は、ゴルーチンとチャネルを使った並行処理をネイティブにサポートし、高いパフォーマンスを実現。

3. **豊富な標準ライブラリ**

    - 幅広い用途に対応する豊富な標準ライブラリが提供されており、外部ライブラリへの依存を最小限に抑えられる。

4. **ガベージコレクション**

    - 自動メモリ管理を行うガベージコレクターを搭載しており、メモリ管理が容易。

5. **クロスプラットフォーム**

    - Windows、Linux、macOS など多くのプラットフォームで動作し、コンパイルされたバイナリは依存関係が少ない。

6. **コミュニティとエコシステム**
    - 活発なコミュニティと豊富なサードパーティライブラリがあり、継続的な開発が行われている。

これらの特徴により、Go 言語は Web サーバー、ネットワークツール、分散システム、クラウドサービスなど、さまざまな用途で利用されています。

---

## 用語説明

### [Module](./terms/module.md)

Go のモジュールは、パッケージの集まりです。`go.mod`ファイルで管理され、依存関係を指定します。

#### サンプルコード

```go
// go.mod
module example.com/mymodule

go 1.20
```

#### ユースケース

プロジェクトのルートディレクトリで`go mod init example.com/mymodule`を実行してモジュールを初期化します。

### [Package](./terms/package.md)

パッケージは関連する Go ソースコードファイルの集まりです。パッケージ名はディレクトリ名に対応します。

#### サンプルコード

```go
// mathutil/mathutil.go
package mathutil

func Add(a, b int) int {
    return a + b
}
```

#### ユースケース

パッケージを使うことでコードを整理し、再利用可能なユニットに分割できます。

### [Variables](./terms/variables.md)

変数は`var`キーワードまたは短縮記法`:=`を使って宣言します。型推論もサポートしています。

#### サンプルコード

```go
var name string = "Go"
age := 10
```

#### ユースケース

データを保存して操作するために変数を使用します。

### [Pointer](./terms/pointer.md)

ポインタは値のメモリアドレスを保持します。`*`と`&`を使って操作します。

#### サンプルコード

```go
var x int = 10
var p *int = &x
*p = 20
```

#### ユースケース

ポインタを使うことで、関数が呼び出し元の変数を直接変更できます。

### [Shadowing](./terms/shadowing.md)

変数のシャドーイングは、外側のスコープの変数を内側のスコープで同じ名前の変数が隠すことです。

#### サンプルコード

```go
x := 10
if true {
    x := 20
    fmt.Println(x) // 20
}
fmt.Println(x) // 10
```

#### ユースケース

特定のスコープでのみ有効な変数を使いたい場合にシャドーイングが発生します。

### [Slice](./terms/slice.md)

スライスは動的サイズの配列のようなデータ構造です。

#### サンプルコード

```go
numbers := []int{1, 2, 3, 4}
numbers = append(numbers, 5)
```

#### ユースケース

可変長のデータを格納するのに便利です。

### [Map](./terms/map.md)

マップはキーと値のペアを格納するデータ構造です。

#### サンプルコード

```go
ages := map[string]int{"Alice": 30, "Bob": 25}
ages["Charlie"] = 35
```

#### ユースケース

キーによって値に素早くアクセスできるため、検索や集計に便利です。

### [Struct](./terms/struct.md)

構造体はフィールドの集合を持つデータ型です。

#### サンプルコード

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
```

#### ユースケース

関連するデータを一つの単位として扱いたいときに使用します。

### [Receiver](./terms/receiver.md)

メソッドのレシーバーは、メソッドが属する型を指定します。

#### サンプルコード

```go
type Rectangle struct {
    Width, Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}
```

#### ユースケース

特定の型に関連する関数を定義する際に使用します。

### [Function](./terms/function.md)

関数は`func`キーワードで定義します。

#### サンプルコード

```go
func add(a int, b int) int {
    return a + b
}
```

#### ユースケース

コードを再利用可能な単位に分割します。

### [Closure](./terms/closure.md)

クロージャは外部の変数を参照する関数です。

#### サンプルコード

```go
func incrementer() func() int {
    var x int
    return func() int {
        x++
        return x
    }
}

inc := incrementer()
fmt.Println(inc()) // 1
fmt.Println(inc()) // 2
```

#### ユースケース

関数が生成された環境を保持し続ける必要がある場合に使用します。

### [Interface](./terms/interface.md)

インターフェースはメソッドの集合を定義します。

#### サンプルコード

```go
type Speaker interface {
    Speak() string
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}
```

#### ユースケース

異なる型が同じ操作を提供できるようにします。

### [If](./terms/if.md)

条件付き実行には`if`文を使います。

#### サンプルコード

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is not greater than 10")
}
```

#### ユースケース

条件に応じて異なるコードを実行します。

### [Switch](./terms/switch.md)

複数の条件を扱う場合は`switch`文を使います。

#### サンプルコード

```go
switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("End of the week")
default:
    fmt.Println("Midweek")
}
```

#### ユースケース

複数の分岐を明確に記述できます。

### [For](./terms/for.md)

ループを実行するために`for`文を使います。

#### サンプルコード

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

#### ユースケース

一定回数の繰り返し処理を行います。

### [Errors](./terms/errors.md)

エラーは組み込みの`error`インターフェースを使って処理します。

#### サンプルコード

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

#### ユースケース

関数のエラー状態を呼び出し元に伝えるために使用します。

### [Generics](./terms/generics.md)

#### 基

礎情報
Go 1.18 から導入されたジェネリクスは、型に依存しない関数やデータ構造を定義できます。

#### サンプルコード

```go
func Map[T any](slice []T, f func(T) T) []T {
    result := make([]T, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}
```

#### ユースケース

型に依存しない再利用可能なコードを記述します。

### [Unit Test (Table Driven, Coverage)](./terms/unit-test.md)

ユニットテストは、`testing`パッケージを使って記述します。テーブル駆動型テストは入力と期待される出力を表にしてテストします。

#### サンプルコード

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b int
        want int
    }{
        {1, 2, 3},
        {2, 2, 4},
        {10, 5, 15},
    }

    for _, tt := range tests {
        got := add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
```

#### ユースケース

さまざまな入力に対して関数の動作を検証します。

### [Logger](./terms/unit-test.md)

ログを記録するために`log`パッケージを使います。

#### サンプルコード

```go
import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    logger.Println("This is an info message")
}
```

#### ユースケース

アプリケーションの動作状況を記録してデバッグや運用を支援します。
