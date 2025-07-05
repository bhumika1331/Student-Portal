package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Registrar struct {
	students    []Student
	courses     []Course
	enrollments []Enrollment
}

// NewRegistrar creates a new Registrar instance
type NewRegistrarS struct {
	Registrar                      // Embedding Registrar to extend its functionality
	teacher    []Teacher           // List of teachers
	Teachermap []TeacherEnrollment // list of Map of teachers with their courses
	enroll     []EnrollNew         // List of enrollments (students with courses) with additional teacher and attendance information
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

// function to adppend teacher into register
func (r *NewRegistrarS) AddTeacher(t Teacher) {
	r.teacher = append(r.teacher, t)
}

// function to add teacher with course in teacher map
func (r *NewRegistrarS) AddTeacherenrollment(te TeacherEnrollment) {
	r.Teachermap = append(r.Teachermap, te)
}

// function to add student into register
func (r *Registrar) AddStudent(s Student) {
	r.students = append(r.students, s)
}

// function to add course into register
func (r *Registrar) AddCourse(c Course) {
	r.courses = append(r.courses, c)
}

// old variable of enrollment which maintain only student with courses
func (r *Registrar) Enroll(e Enrollment) {
	r.enrollments = append(r.enrollments, e)
}

func (r *NewRegistrarS) Enrollnew(e EnrollNew) {
	r.enroll = append(r.enroll, e)
}

// new variable to maintain enrollments with additional information like teacher and attendance
func (r *NewRegistrarS) AddEnrollnew(e EnrollNew) {
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
