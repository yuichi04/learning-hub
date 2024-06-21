package main

import (
	"fmt"
	"math"
)

var taxRate float64 = 10.0

func calculateTotalPrice(price float64, totalTax float64, round bool) int {
	total := price + totalTax
	if round {
		return int(math.Round(total))
	}
	return int(total)
}

func calculateSubtotalPrice(price float64, quantity int) float64 {
	return price * float64(quantity)
}

func calculateTotalTax(price float64, round bool) float64 {
	total := price * (taxRate / 100)
	if round {
		return math.Round(total)
	}
	return total
}

func main() {
	var price float64
	fmt.Print("商品の価格を入力して下さい: ")
	_, err := fmt.Scan(&price)
	if err != nil {
		fmt.Println("入力エラー:", err)
		return
	}

	var quantity int
	fmt.Print("商品の数量を入力して下さい: ")
	_, err = fmt.Scan(&quantity)
	if err != nil {
		fmt.Println("数量の入力ラー:", err)
		return
	}

	var roundChoice string
	fmt.Print("小数点以下を四捨五入しますか？（y/n）")
	_, err = fmt.Scan(&roundChoice)
	if err != nil {
		fmt.Println("入力エラー:", err)
		return
	}

	round := roundChoice == "y" || roundChoice == "Y"

	subtotalPrice := calculateSubtotalPrice(price, quantity)
	totalTax := calculateTotalTax(subtotalPrice, round)
	totalPrice := calculateTotalPrice(subtotalPrice, totalTax, round)
	fmt.Printf("消費税： %v\n", totalTax)
	fmt.Printf("合計金額（税込）: %v\n", totalPrice)
}
