package main

import "fmt"

func main() {
	// Esspresso 構造体を Beverage インターフェースとして受け取る
	var drink Beverage = NewEspresso()

	// Mocha コンストラクタは Beverage インターフェースを受け取って、その中の Description と Cost を取得できる
	drink = NewMocha(drink)
	fmt.Println(drink.GetDescription(), drink.Cost())

	drink = NewWhip(drink)
	fmt.Println(drink.GetDescription(), drink.Cost())

}
