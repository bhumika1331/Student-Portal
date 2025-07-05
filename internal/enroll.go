package internal

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

// made enrool function to check with teacher maping with course (changes may be like checking in main while adding or passing teacher map from rigister to enroll function)
func Enroll(oe Enrollment, att Attendance, t Teacher, teachermap []TeacherEnrollment) (EnrollNew, bool) {
	for _, te := range teachermap {
		if te.Teacher.ID == t.TID() && te.Course.Id == oe.Course.Id {
			return EnrollNew{Enrollment: oe, Attend: att, Teacher: t}, true
		}
	}
	return EnrollNew{}, false
}

/*
=======
// func enrollnew(st Student, c Course, g Grader, score float64, attend Attendance, t Teacher) Enrollnew {
// 	return Enrollnew{Enrollment: Enrollment{Student: st, Course: c, Grader: g, score: score}, Attend: attend, Teacher: t}
// }


// made enrool function to check with teacher maping with course (changes may be like checking in main while adding or passing teacher map from rigister to enroll function)
// func Enroll(st Student, c Course, g Grader, score float64, att Attendance, t Teacher, teachermap []Teacherenrollment) (Enrollnew, bool) {
func Enroll(oe Enrollment, att Attendance, t Teacher, teachermap []Teacherenrollment) (Enrollnew, bool) {
	for _, te := range teachermap {
		if te.Teacher.ID() == t.ID() && te.Course.Id == oe.Course.Id {
			return Enrollnew{Enrollment: oe, Attend: att, Teacher: t}, true
		}
	}
	return Enrollnew{}, false
}

/*
func (e Enrollment) String() string {
	r, _ := e.Grade(e)
	return fmt.Sprintf("%s ->  %s : %s", e.Student.Name(), e.Course.Name, r)
}
*/
