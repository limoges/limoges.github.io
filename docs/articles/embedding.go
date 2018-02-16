package main

import "fmt"

type Humanoid interface {
	Climb() string
}

type Monkey struct {
	// Empty
}

func (m Monkey) Climb() string {
	return "Monkey climbs"
}

type Human struct {
	Monkey
}

func (h Human) Climb() string {
	return "Human climbs"
}

func (h Human) Run() string {
	return "Human runs"
}

type Monkeyman struct {
	Human
	Monkey
}

func (m Monkeyman) Run() string {
	return fmt.Sprintf("Superhuman runs while %s", m.Human.Climb())
}

func main() {
	var humanoid Humanoid
	humanoid = Monkey{}
	fmt.Println("humanoid<Monkey>.Climb():", humanoid.Climb())

	humanoid = Human{}
	fmt.Println(" humanoid<Human>.Climb():", humanoid.Climb())

	monkeyman := Monkeyman{}
	fmt.Println("        monkeyman.Run():", monkeyman.Run())

}
