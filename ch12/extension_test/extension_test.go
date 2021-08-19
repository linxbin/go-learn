package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

type Dog struct {
	Pet
}

type Cat struct {
	Pet
}

func (p *Pet) Jump() {
	fmt.Print("jump")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)

	dog.SpeakTo("chao")

	cat := &Cat{}
	cat.Jump()
}
