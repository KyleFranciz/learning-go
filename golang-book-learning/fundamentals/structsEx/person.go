package structsex

type Person struct {
	name   string
	age    int
	weight int
	race   string
}

func CreateNewPerson(name string, age int, weight int, race string) *Person {
	// func creates a new person points to it and returns the adress of the newly created person
	return &Person{
		name,
		age,
		weight,
		race,
	}
}

func StructUseCaseEx() Person {
	// this returns a person struct w/ the feilds at their Zero value
	kyle := Person{} // only age and weight will show
	return kyle
}
