## Go 言語のユニットテスト（Unit Test）について

### 概要

ユニットテストは、ソフトウェアの個々のコンポーネント（ユニット）をテストするための手法です。Go 言語では、標準ライブラリの`testing`パッケージを使用してユニットテストを記述できます。テーブル駆動テストは、さまざまな入力データセットに対して同じテストコードを実行するための効果的な方法です。テストカバレッジは、テストがソースコードのどれだけをカバーしているかを測定する手段です。

### ユニットテストの基本

Go のユニットテストは、`*_test.go`ファイルに記述され、`func TestXxx(t *testing.T)`という形式の関数で定義されます。テスト関数は`go test`コマンドで実行されます。

#### サンプルコード

```go
package main

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    expected := 3
    if result != expected {
        t.Errorf("Add(1, 2) = %d; want %d", result, expected)
    }
}
```

この例では、`Add`関数のテストを`TestAdd`関数で定義し、期待される結果と実際の結果を比較しています。

### テーブル駆動テスト

テーブル駆動テストは、異なる入力データセットに対して同じテストコードを実行するためのパターンです。テストケースを構造体のスライスとして定義し、ループを使ってテストを実行します。

#### サンプルコード

```go
package main

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"1+1", 1, 1, 2},
        {"2+2", 2, 2, 4},
        {"3+3", 3, 3, 6},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

この例では、テストケースを構造体のスライスとして定義し、`t.Run`を使用して各テストケースを実行しています。

### テストカバレッジ

テストカバレッジは、テストがソースコードのどれだけをカバーしているかを測定します。Go では、`go test`コマンドに`-cover`フラグを付けることでテストカバレッジを確認できます。

#### コマンド例

```sh
go test -cover
```

さらに、カバレッジの詳細なレポートを生成するには、`-coverprofile`フラグを使用します。

#### コマンド例

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

これにより、カバレッジレポートが生成され、ブラウザで確認できるようになります。

### テストのベストプラクティス

1. **小さなテスト関数**

    - 各テスト関数は 1 つの機能や挙動をテストするようにします。

2. **テーブル駆動テストの使用**

    - 入力データセットが複数ある場合、テーブル駆動テストを使用して効率的にテストを実行します。

3. **テストカバレッジの向上**

    - テストカバレッジを測定し、不足している部分のテストを追加します。

4. **`t.Run`の使用**

    - `t.Run`を使用して、個々のテストケースを名前付きで実行し、詳細なエラーメッセージを提供します。

5. **モックとスタブの利用**
    - 外部依存関係を持つコードをテストする際には、モックやスタブを使用して依存関係をシミュレートします。

### モックの使用

モックは、依存関係をシミュレートするためのテストダブルです。`testing`パッケージを使用して、モックを作成できます。

#### サンプルコード

```go
package main

import (
    "fmt"
    "testing"
)

// 依存関係のインターフェース
type Greeter interface {
    Greet() string
}

// 依存関係のモック
type MockGreeter struct{}

func (m *MockGreeter) Greet() string {
    return "Hello, Test!"
}

// テスト対象の関数
func SayHello(g Greeter) string {
    return g.Greet()
}

func TestSayHello(t *testing.T) {
    mock := &MockGreeter{}
    result := SayHello(mock)
    expected := "Hello, Test!"
    if result != expected {
        t.Errorf("SayHello() = %s; want %s", result, expected)
    }
}
```

この例では、`Greeter`インターフェースを定義し、それを実装するモック`MockGreeter`を作成しています。テストでは、このモックを使用して依存関係をシミュレートしています。

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#testing)
-   [Go by Example: Testing](https://gobyexample.com/testing)
-   [Testing Techniques in Go](https://blog.golang.org/cover)
