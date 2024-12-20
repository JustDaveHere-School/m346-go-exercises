package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}

	classA := Class{
		Name: "Klasse A",
		Students: []Student{
			{"VornameA1", "NachnameA1"},
			{"VornameA2", "NachnameA2"},
			{"VornameA3", "NachnameA3"},
		},
	}

	classB := Class{
		Name: "Klasse B",
		Students: []Student{
			{"VornameB1", "NachnameB1"},
			{"VornameB2", "NachnameB2"},
			{"VornameB3", "NachnameB3"},
		},
	}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int][]Class{
		104: {classA, classB},
		117: {classA},
		346: {classB},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println(classA, "\n", "\n", classB, "\n", modules)
}
