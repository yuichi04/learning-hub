# 環境構築

## 本体のインストール

```shell
brew install go
```

## ディレクトリの作成

```shell
mkdir myproject
cd myproject
```

## モジュールの初期化

```shell
go mod init myproject
```

## サーバーの起動

<u>例）HTTP サーバーを起動し、ブラウザに「Hello, World!」を出力する</u>

### 1. プログラムの作成

#### ファイルを作成

```shell
touch main.go
```

#### ファイルの内容

```go
package main

import (
    "fmt"
    "net/http"
)
// ハンドラ関数を定義
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // ハンドラ関数をルートパスに割り当て
    http.HandleFunc("/", helloHandler)
    // サーバーをポート8080で起動
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
```

### 2. ソースコードのコンパイルとプログラムの実行

```shell
go run main.go
```
