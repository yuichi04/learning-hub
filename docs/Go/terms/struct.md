## Go 言語の構造体（Struct）について

### 概要

構造体は、フィールドと呼ばれる名前付きのデータを集めたデータ型です。複数の異なる型のデータをひとまとめにして扱うことができます。Go 言語の構造体は、クラスを持たない言語におけるオブジェクト指向プログラミングの基本単位とも言えます。

### 構造体の定義

構造体を定義するには、`type`キーワードと`struct`キーワードを使用します。

#### サンプルコード

```go
type Person struct {
    Name string
    Age  int
}
```

この例では、`Person`という名前の構造体を定義し、`Name`という文字列型のフィールドと、`Age`という整数型のフィールドを持たせています。

### 構造体のインスタンス化

構造体のインスタンスを作成するには、構造体リテラルを使用します。

#### サンプルコード

```go
p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name) // Alice
fmt.Println(p.Age)  // 30
```

また、フィールドを省略するとデフォルト値（ゼロ値）で初期化されます。

```go
p := Person{Name: "Bob"}
fmt.Println(p.Name) // Bob
fmt.Println(p.Age)  // 0
```

### 構造体のポインタ

構造体のポインタを使用すると、メモリ効率の向上や、メソッドでのレシーバーとしての使用が可能になります。

#### サンプルコード

```go
p := &Person{Name: "Charlie", Age: 25}
fmt.Println(p.Name) // Charlie
fmt.Println(p.Age)  // 25
```

構造体のポインタを使っても、`.`（ドット）演算子を使ってフィールドにアクセスできます。内部的には、`(*p).Name`のように解釈されます。

### 構造体のメソッド

構造体にメソッドを定義することができます。メソッドは特定の型に関連付けられた関数です。

#### サンプルコード

```go
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Greet()) // Hello, my name is Alice
}
```

### ポインタレシーバーと値レシーバー

メソッドには、ポインタレシーバーと値レシーバーの 2 種類があります。ポインタレシーバーは、構造体のフィールドを変更するために使用され、値レシーバーはフィールドの値を変更しません。

#### サンプルコード（ポインタレシーバー）

```go
func (p *Person) HaveBirthday() {
    p.Age++
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    p.HaveBirthday()
    fmt.Println(p.Age) // 31
}
```

### 無名フィールド（埋め込み構造体）

構造体は他の構造体をフィールドとして埋め込むことができます。これにより、ある種の継承のような機能を実現できます。

#### サンプルコード

```go
type Address struct {
    City, State string
}

type Person struct {
    Name    string
    Age     int
    Address // 無名フィールド（埋め込み構造体）
}

func main() {
    p := Person{Name: "Alice", Age: 30, Address: Address{City: "New York", State: "NY"}}
    fmt.Println(p.City)  // New York
    fmt.Println(p.State) // NY
}
```

### タグ付き構造体フィールド

構造体フィールドにはタグを付けることができます。タグは、構造体のフィールドにメタデータを追加するために使用され、特に JSON のマーシャリング/アンマーシャリングやデータベース操作に役立ちます。

#### サンプルコード

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    jsonData, _ := json.Marshal(p)
    fmt.Println(string(jsonData)) // {"name":"Alice","age":30}
}
```

### 構造体の比較

Go 言語では、同じ型の構造体は比較可能です。ただし、全てのフィールドが比較可能である必要があります。

#### サンプルコード

```go
type Point struct {
    X, Y int
}

func main() {
    p1 := Point{X: 1, Y: 2}
    p2 := Point{X: 1, Y: 2}
    fmt.Println(p1 == p2) // true
}
```

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#composite)
-   [Go by Example: Structs](https://gobyexample.com/structs)
