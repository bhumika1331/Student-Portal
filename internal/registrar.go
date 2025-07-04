package internal

import (
	"encoding/json"
	"fmt"
	"log"

	//"oops/main/internal"

	//"oops/main/analytics"
	//"oops/main/domain"
	"os"
)

type Registrar struct {
	students    []Student
	courses     []Course
	enrollments []Enrollment
	graders     map[int]Grader
}

func NewRegistrar() *Registrar {
	return &Registrar{graders: make(map[int]Grader)}
}

type NewRegistrarS struct {
	Registrar
	teacher    []Teacher
	teachermap []TeacherEnrollment
	enroll     []EnrollNew
}

type RegistrarWithDocs struct {
	*NewRegistrarS
	enrollWithDocs []EnrollnewWithDocs
}

func (r *RegistrarWithDocs) EnrollnewWithDocs(e EnrollnewWithDocs) {
	r.enrollWithDocs = append(r.enrollWithDocs, e)
}

func (r *RegistrarWithDocs) DisplayDocuments() {
	for _, e := range r.enrollWithDocs {
		fmt.Printf("Student: %s (ID %d)\n", e.Student.Name(), e.Student.ID())
		for _, doc := range e.Documents {
			fmt.Printf(" - %s (%s)\n", doc.Title, doc.Filename)
		}
	}
}

func (r *NewRegistrarS) AddTeacher(t Teacher) {
	r.teacher = append(r.teacher, t)
}

func (r *NewRegistrarS) AddTeacherenrollment(te TeacherEnrollment) {
	r.teachermap = append(r.teachermap, te)
}

func (r *Registrar) AddStudent(s Student) {
	r.students = append(r.students, s)
}

func (r *Registrar) AddCourse(c Course) {
	r.courses = append(r.courses, c)
}

func (r *Registrar) Enroll(e Enrollment) {
	r.enrollments = append(r.enrollments, e)
}

func (r *NewRegistrarS) Enrollnew(e EnrollNew) {
	r.enroll = append(r.enroll, e)
}

func (r *Registrar) SetGrader(courseID int, g Grader) {
	for i, e := range r.enrollments {
		if e.Course.Id == courseID {
			r.enrollments[i].Grader = g
		}
	}
}

func (r *Registrar) Enrollments() []Enrollment {
	return r.enrollments
}

type StudentData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type courseData struct {
	ID      int     `json:"id"`
	Name    string  `json:"title"`
	Credits float64 `json:"credits"`
}

func (regis *Registrar) LoadCourses() {
	var coursesData []courseData
	data, err := os.ReadFile("courses.json")
	if err != nil {
		log.Fatal("Failed to read courses.json:", err)
	}
	err = json.Unmarshal(data, &coursesData)
	if err != nil {
		log.Fatal("Failed to unmarshal courses:", err)
	}
	for _, course := range coursesData {
		regis.AddCourse(NewCourse(course.ID, course.Name))
	}
}

func (r *Registrar) LoadStudents() {
	var students []StudentData
	data, err := os.ReadFile("students.json")
	if err != nil {
		log.Fatal("Failed to read students.json:", err)
	}

	err = json.Unmarshal(data, &students)
	if err != nil {
		log.Fatal("Failed to unmarshal students:", err)
	}
	r.students = make([]Student, 0, len(students))
	for _, sd := range students {
		student := NewStudent(sd.ID, sd.Name)
		r.AddStudent(student)
	}
}

func (r *Registrar) DisplayStudents() {
	for _, sd := range r.students {
		fmt.Printf("#%d : %s\n", sd.ID(), sd.Name())
	}
}

func (r *Registrar) DisplayCourses() {
	log.Println("Courses...")
	for _, cr := range r.courses {
		log.Printf("#%d : %s\n", cr.Id, cr.Name)
	}
}
