package simple

import "fmt"

type Factory interface {
	Print(words string) string
}

func NewFactory(typ int) Factory {
	switch typ {
	case 1:
		return &A{}
	case 2:
		return &B{}
	}
	return nil
}

type A struct{}

func (a *A) Print(words string) string {
	return fmt.Sprintf("A: %s", words)
}

type B struct{}

func (b *B) Print(words string) string {
	return fmt.Sprintf("B: %s", words)
}
