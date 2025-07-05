package internal

import "fmt"

type Student struct {
	id   int
	name string
}

func NewStudent(id int, name string) Student {
	if id <= 0 {
		panic("Student id must be positive")
	}
	return Student{id: id, name: name}
}

func (s Student) Name() string { return s.name }
func (s Student) ID() int      { return s.id }

func (s Student) Display() {
	fmt.Printf("Student #%d : %s\n", s.id, s.name)
}
