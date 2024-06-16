# ローカル環境構築(Mac)

## Go のインストール

### Homebrew がインストールされていることを確認

```shell
brew -v
```

### Homebrew がインストールされていない場合のインストール

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### Go のインストール

```shell
brew install go
```

## パスの追加

```shell
# bash
echo 'export GOPATH=$HOME/go' >> ~/.bash_profile
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
source ~/.bash_profile

# z shell
echo 'export GOPATH=$HOME/go' >> ~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
source ~/.zshrc
```

## インストールの確認

```shell
go version
```

### Go の環境情報を確認

```shell
go env
```

## プロジェクトフォルダの作成

```shell
mkdir my-project && cd $_
```

### プロジェクトディレクトリ構造

```plaintext
my-project/
├── go.mod
└── main.go
```

## モジュールの初期化

```shell
go mod init my-project
```

## 外部パッケージの自動インストール/アンインストール

go mod tidy コマンドは、go.mod ファイルをクリーンアップし、実際に使用されている依存関係のみを保持します。

```shell
go mod tidy
```

### 外部パッケージのインストール例

```shell
# 例: Gin Web Frameworkのインストール
go get -u github.com/gin-gonic/gin
```

## サーバーの起動（ブラウザに`hello world`が表示されるまで）

### エントリーポイントを作成

```shell
touch main.go
```

### サンプルコード

```go
package main

import (
    "fmt"
    "net/http"
)

// helloHandlerは"/"のリクエストを処理
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    // "/"パスにhelloHandlerを割り当てる
    http.HandleFunc("/", helloHandler)
    fmt.Println("Starting server at :8080")

    // サーバーをポート8080で起動
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
```

### ソースコードのコンパイルとプログラムの実行

```shell
go run main.go
```

### サーバーの停止

```plaintext
サーバーを停止するには、実行中のターミナルでCtrl+Cを押します。
```

## ビルドの実行（バイナリデータに変換される）

go build コマンドは、Go のソースコードをコンパイルして実行可能なバイナリファイルを生成します。
生成されたバイナリファイルを実行することで、プログラムを実行できます。

```shell
# `-o`はバイナリデータのファイル名を指定できるオプション
go build -o app main.go
```

### ビルドしたバイナリの実行

```shell
./app
```
