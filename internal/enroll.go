package internal

import (
	"fmt"
)

type Enrollment struct {
	Student
	Course
	Grader
	score float64
}

type EnrollNew struct {
	Enrollment            // Embedding Enrollment to include student, course, grader, and score
	Attend     Attendance // Attendance records associated with the enrollment
	Teacher               // Embedding Teacher to associate with the enrollment
}

type TeacherEnrollment struct {
	Teacher
	CreditCourse
}

func NewEnrollNew(st Student, c Course, g Grader, score float64, attend Attendance, t Teacher) EnrollNew {
	return EnrollNew{
		Enrollment: Enrollment{
			Student: st,
			Course:  c,
			Grader:  g,
			score:   score,
		},
		Attend:  attend,
		Teacher: t,
	}
}

func NewTeacherEnrollment(t Teacher, c CreditCourse) TeacherEnrollment {
	return TeacherEnrollment{Teacher: t, CreditCourse: c}
}

func NewEnrollment(st Student, c Course, g Grader, score float64) Enrollment {
	return Enrollment{Student: st, Course: c, Grader: g, score: score}
}

// func enrollnew(st Student, c Course, g Grader, score float64, attend Attendance, t Teacher) Enrollnew {
// 	return Enrollnew{Enrollment: Enrollment{Student: st, Course: c, Grader: g, score: score}, Attend: attend, Teacher: t}
// }

func (e Enrollment) String() string {
	r, _ := e.Grade(e)
	return fmt.Sprintf("%s ->  %s : %s", e.Student.Name(), e.Course.Name, r)
}
