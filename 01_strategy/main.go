package main

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}

func (f FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings!")
}

type FlyNoWay struct{}

func (f FlyNoWay) Fly() {
	fmt.Println("I can't fly.")
}

type Duck struct {
	flyBehavior FlyBehavior
}

// ここの引数 Duck が値レシーバーだと main で呼んだときにカモは飛べないまま
func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

// こっちはポインタレシーバーである必要はないけど、慣例的にポインタレシーバーに統一するらしい
func (d *Duck) PeformFly() {
	d.flyBehavior.Fly()
}

// ModelDuck は初期状態では飛べない
func NewModelDuck() *Duck {
	return &Duck{
		flyBehavior: FlyNoWay{},
	}
}

func main() {
	duck := NewModelDuck()

	fmt.Print("Before upgrading: ")
	duck.PeformFly()

	duck.SetFlyBehavior(FlyWithWings{})

	fmt.Print("After upgrading: ")
	duck.PeformFly()
}
