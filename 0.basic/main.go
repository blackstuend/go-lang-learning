package main

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) String() string {
	return p.FirstName + p.LastName
}

func main() {
	person := Person{"hello", "kitty"}

	Persons := make([]person, 0)
}
