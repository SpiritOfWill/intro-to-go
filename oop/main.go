package main

/*

1. Encapsulation - Go encapsulates things at the package level. Names that start with a lowercase letter are only visible within that package.
2. Composition, inheritance - No multiple inheritance.
3. Polymorphism

https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540

*/

import "fmt"

type Animal struct{}

func (*Animal) Walk() string {
	return "🚶🏻‍♂"
}

func (*Animal) Die() string {
	return "💀"
}

type Dog struct {
	*Animal // anonymous composition via embedding!!!
}

func NewDog() *Dog {
	return &Dog{}
}

func (*Dog) Bark() string {
	return "bark"
}

type Cat struct {
	*Animal

	lives uint
}

func NewCat() *Cat {
	return &Cat{
		lives: 9,
	}
}

func (*Cat) Meow() string {
	return "meow"
}

// overriding:
func (c *Cat) Die() string {
	c.lives--

	if c.lives == 0 {
		return "💀"
	}

	return fmt.Sprintf("left %d lives", c.lives)
}

// Walker - just interface(not related to inheritance, added for usability)
type Walker interface {
	Walk() string
}

func main() {
	dog := NewDog()
	cat := NewCat()

	animals := []Walker{dog, cat}

	for _, animal := range animals {
		fmt.Printf("%T Walk(): %s\n", animal, animal.Walk())
	}

	fmt.Println()

	fmt.Println("🐶 Bark():", dog.Bark())
	fmt.Println("😼 Meow():", cat.Meow())

	fmt.Println()

	// no way to find out from whom .Die() method was inherited:
	fmt.Println("🐶 Die():", dog.Die())
	fmt.Println("😼 Die():", cat.Die())
	fmt.Println("😼 Die():", cat.Die())

	// but parent method is not overridden and available:
	fmt.Println("cat.Animal.Die():", cat.Animal.Die())
}
