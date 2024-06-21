package main

import "fmt"

func main() {
	var price float64
	fmt.Print("商品の価格を入力して下さい: ")
	fmt.Scan(&price)

	// 消費税率10%
	taxRate := 0.10

	// 消費税額を計算
	taxAmount := price * taxRate

	// 最終的な支払額を計算
	finalAmount := price + taxAmount

	fmt.Printf("商品価格: %.2f円\n", price)
	fmt.Printf("消費税額: %.2f円\n", taxAmount)
	fmt.Printf("支払額合計: %.2f円\n", finalAmount)
}
