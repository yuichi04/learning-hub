## Go 言語のエラー（Errors）について

### 概要

Go 言語におけるエラーハンドリングは、組み込みの`error`インターフェースを使用します。エラーは通常、関数の戻り値として返され、呼び出し元で処理されます。Go のエラーハンドリングは、シンプルで明示的な方法を提供し、プログラムの信頼性を高めます。

### `error`インターフェース

`error`インターフェースは、単一のメソッド`Error`を持ちます。このメソッドは、エラーメッセージを文字列として返します。

#### `error`インターフェースの定義

```go
type error interface {
    Error() string
}
```

### エラーの生成

標準ライブラリの`errors`パッケージを使用して、エラーを生成できます。

#### サンプルコード

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("an error occurred")
    fmt.Println(err)
}
```

### カスタムエラー

カスタムエラーを作成するために、新しい型を定義し、`Error`メソッドを実装することができます。

#### サンプルコード

```go
package main

import (
    "fmt"
)

type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func main() {
    err := &MyError{Code: 123, Message: "Something went wrong"}
    fmt.Println(err)
}
```

### エラーの返却

関数やメソッドでエラーを返すには、戻り値に`error`を含めます。エラーが発生した場合、`nil`以外の値を返します。

#### サンプルコード

```go
package main

import (
    "fmt"
    "errors"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### エラーチェック

エラーが返された場合、`if`文を使ってエラーをチェックし、適切な処理を行います。

#### サンプルコード

```go
package main

import (
    "fmt"
    "errors"
)

func main() {
    err := errors.New("an error occurred")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

### ラップされたエラー

Go 1.13 以降、`fmt.Errorf`を使ってエラーをラップし、元のエラーを保持することができます。

#### サンプルコード

```go
package main

import (
    "fmt"
    "errors"
)

func doSomething() error {
    return errors.New("original error")
}

func main() {
    err := doSomething()
    if err != nil {
        wrappedErr := fmt.Errorf("failed to do something: %w", err)
        fmt.Println(wrappedErr)
    }
}
```

### エラーの展開

ラップされたエラーを展開するには、`errors.Unwrap`関数や`errors.Is`、`errors.As`関数を使用します。

#### サンプルコード

```go
package main

import (
    "errors"
    "fmt"
)

func doSomething() error {
    return errors.New("original error")
}

func main() {
    err := doSomething()
    if err != nil {
        wrappedErr := fmt.Errorf("failed to do something: %w", err)
        fmt.Println(wrappedErr)

        // Unwrap
        if unwrappedErr := errors.Unwrap(wrappedErr); unwrappedErr != nil {
            fmt.Println("Unwrapped error:", unwrappedErr)
        }

        // Is
        if errors.Is(wrappedErr, err) {
            fmt.Println("Wrapped error matches original error")
        }

        // As
        var myErr *MyError
        if errors.As(wrappedErr, &myErr) {
            fmt.Println("Error is of type MyError:", myErr)
        }
    }
}
```

### エラーハンドリングのベストプラクティス

1. **明示的にエラーチェックを行う**

    - エラーを返す関数の呼び出し後には、必ずエラーチェックを行います。

2. **詳細なエラーメッセージを提供する**

    - エラーメッセージには、問題の詳細を含めるようにします。

3. **エラーのラップと展開**

    - 必要に応じてエラーをラップし、元のエラー情報を保持します。

4. **特定のエラータイプを定義する**
    - カスタムエラータイプを定義し、エラーをより具体的に識別できるようにします。

### 参考文献

-   [公式ドキュメント](https://golang.org/doc/effective_go.html#errors)
-   [Go by Example: Errors](https://gobyexample.com/errors)
-   [Go Blog: Error Handling and Go](https://blog.golang.org/error-handling-and-go)
