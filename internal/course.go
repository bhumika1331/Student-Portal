package internal

type Course struct {
	Id   int
	Name string
}

type CreditCourse struct {
	Course
	Credits int
}

func NewCreditCourse(c Course, credits int) CreditCourse {
	return CreditCourse{Course: c, Credits: credits}
}

func NewCourse(id int, name string) Course {
	return Course{Id: id, Name: name}
}
