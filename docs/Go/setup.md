# 環境構築

## 本体のインストール

```shell
brew install go
```

## パスの追加（mac）

```shell
# bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
source ~/.bash_profile

# z shell
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
```

## インストールの確認

```shell
go version
```

## プロジェクトフォルダの作成

```shell
mkdir my-project && cd $_
```

## モジュールの初期化

```shell
go mod init my-project
```

## 外部パッケージの自動インストール/アンインストール

```shell
go mod tidy
```

## サーバーの起動（ブラウザに`hello world`が表示されるまで）

### エントリーポイントを作成

```shell
touch main.go
```

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // io.Writer(`http.ResponseWriter`)にフォーマットされた出力を行う
    fmt.Fprintln(w, "Hello, World!")
}

// Goではmain関数が最初に実行されるようになっている
func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Starting server at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
```

### ソースコードのコンパイルとプログラムの実行

実行後、`localhost:8080`にアクセスすると`Hello, World!`が表示される

```shell
go run main.go
```

## ビルドの実行（バイナリデータに変換される）

```shell
# `-o`はバイナリデータのファイル名を指定できるオプション
go build -o app main.go
```

```shell
# `Hello, World`が出力される
./app
```
