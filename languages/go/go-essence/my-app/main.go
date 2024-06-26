package main

import "fmt"

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

func main() {
	dog := Dog{}
	DoSpeak(&dog)

	cat := Cat{}
	DoSpeak(&cat)
}
