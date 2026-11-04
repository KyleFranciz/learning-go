package structsex

type person1 struct {
	name string
	age  int
}

var person2 struct {
	// anonymous struct
	name string
	age  int
}

func Comparison() bool {
	// name struct
	p := person1{
		name: "Kyle",
		age:  26,
	}

	// the values become comparable if the values of the fields are the same
	person2.name = "Kyle"
	person2.age = 26

	// only named and anonymous structs are comparable
	return p == person2
}
