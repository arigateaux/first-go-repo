package animals

type Dog struct{}
type Cat struct{}

func (a Dog) Speaks() string {
	return "Woof!"
}

func (a Cat) Speaks() string {
	return "Meow!"
}
