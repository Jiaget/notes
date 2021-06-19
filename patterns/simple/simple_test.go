package simple

import "testing"

func TestA(t *testing.T) {
	product := NewFactory(1)
	s := product.Print("It's A")
	if s != "A: It's A" {
		t.Fatal("testA failed")
	}
}

func TestB(t *testing.T) {
	product := NewFactory(2)
	s := product.Print("It's B")
	if s != "B: It's B" {
		t.Fatal("testB failed")
	}
}
