package main

import (
	"fmt"
)

type Value int

func (v Value) Add(n Value) Value {
	fmt.Println(v)
	return v + n
}

/*
コンストラクタ
Goにはコンストラクタがない。初期処理としてstructの初期化を行い、そのstructのポインタを返したいような場合は、
Goでは習慣としてNewを接頭語につけた関数を用意する
*/
type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

/*
インターフェース
structは実態を持つデータだが、interfaceはメソッドを持つ型のインターフェースを定義できる
*/
type Speaker interface {
	Speak() error
}

type Dog struct{}

func (d *Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct{}

func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}

/*
defer
deferは後処理を記述できる。Goにおける後処理とは関数を抜ける際に実行される処理をdeferで宣言することを意味する。
deferは関数スコープで実行されることに注意
*/
// func doSomething(dir string) {
// f, err := os.Open("data.txt")
// if err != nil {
// 	log.Fatal(err)
// }
// defer f.Close()

// var b [512]byte
// n, err := f.Read(b[:])
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(string(b[:n]))

// deferは指定された順とは逆順に実行される
// ディレクトリとファイルを生成して処理を行った後、ファイルを閉じてディレクトリを削除するような後処理も簡単に実装できる
// defer fmt.Println("6")
// defer fmt.Println("5")
// defer fmt.Println("4")
// fmt.Println("1")
// fmt.Println("2")
// fmt.Println("3")

// err := os.Mkdir(dir, 0755)
// if err != nil {
// 	return err
// }
// defer os.RemoveAll(dir)

// f, err := os.Create(filepath.Join(dir, "data.txt"))
// if err != nil {
// 	return err
// }
// defer f.Close()
// }

// deferには無名関数を渡す
func doSomething() {
	var n = 1
	defer func() {
		fmt.Println(n) // nはキャプチャされていないため2が出力される
	}()
	n = 2
}

// deferに通常の関数を渡す
func doSomething2() {
	var n = 1
	defer fmt.Println(n) // nがキャプチャされているため1が出力される
	n = 2
}

// doSomethingと同じ処理
func doSomethingSame() {
	var n = 1
	n = 2

	func() {
		fmt.Println(n) // 2
	}()
}

// doSomething2と同じ処理
func doSomething2Same() {
	// var n = 1
	// n = 2

	fmt.Println(1) // deferが実行された時のnは1なので実質これと同じ
}

func main() {
	dog := Dog{}
	DoSpeak(&dog)

	cat := Cat{}
	DoSpeak(&cat)

	doSomething()
	doSomethingSame()

	doSomething2()
	doSomething2Same()
}
