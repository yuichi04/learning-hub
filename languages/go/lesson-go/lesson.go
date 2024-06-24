// // Structオリエンテッド
// package main

// import "fmt"

// type Vertex struct {
// 	X, Y int
// }

// func (v Vertex) Area() int {
// 	return v.X * v.Y
// }

// // ポインタレシーバー
// func (v *Vertex) Scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

// func Area(v Vertex) int {
// 	return v.X * v.Y
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(Area(v))
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

// // コンストラクタ
// package main

// import "fmt"

// type Vertex struct {
// 	x, y int
// }

// func (v Vertex) Area() int {
// 	return v.x * v.y
// }

// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

// func main() {
// 	v := New(3, 4)
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }
