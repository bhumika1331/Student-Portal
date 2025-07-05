package internal

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

func createMockRegistrar() NewRegistrarS {
	student := NewStudent(2, "Bob")
	course := NewCourse(202, "Python")
	creditCourse := NewCreditCourse(course, 4)
	teacher := Teacher{ID: "2", Name: "Prof. Smith"}
	teacherenrollment := NewTeacherEnrollment(teacher, creditCourse)
	enrollment := NewEnrollment(student, course, nil, 0.0)
	attendance := Attendance{
		Records: map[time.Time]bool{
			time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC): true,
		},
	}
	enrollNew, _ := Enroll(enrollment, attendance, teacher, []TeacherEnrollment{teacherenrollment})

	r := NewRegistrarS{}
	r.AddStudent(student)
	r.AddCourse(course)
	r.AddTeacher(teacher)
	r.AddTeacherenrollment(teacherenrollment)
	r.Enroll(enrollment)
	r.AddEnrollnew(enrollNew)

	return r
}

func TestAddFunctions(t *testing.T) {
	reg := NewRegistrarS{}

	s := NewStudent(3, "John")
	c := NewCourse(301, "Java")
	cc := NewCreditCourse(c, 3)
	te := Teacher{ID: "01", Name: "TeacherX"}
	tenroll := NewTeacherEnrollment(te, cc)
	enroll := NewEnrollment(s, c, nil, 85.0)
	att := Attendance{Records: map[time.Time]bool{time.Now(): true}}

	reg.AddStudent(s)
	reg.AddCourse(c)
	reg.AddTeacher(te)
	reg.AddTeacherenrollment(tenroll)
	reg.Enroll(enroll)

	enew, ok := Enroll(enroll, att, te, []TeacherEnrollment{tenroll})
	if !ok {
		t.Error("Expected successful enrollment with teacher mapping")
	}
	reg.AddEnrollnew(enew)
}

func TestEnrollFailure(t *testing.T) {
	student := NewStudent(4, "Fail Student")
	course := NewCourse(404, "FailCourse")
	enroll := NewEnrollment(student, course, nil, 0.0)
	att := Attendance{}
	teacher := Teacher{ID: "99", Name: "Unassigned"}

	enew, ok := Enroll(enroll, att, teacher, []TeacherEnrollment{}) // No match
	if ok {
		t.Error("Expected failure due to unmatched teacher mapping")
	}
	if enew.Teacher.ID != "0" {
		t.Error("Expected empty Enrollnew struct")
	}
}

func TestGiveAndFetchAttendance(t *testing.T) {
	reg := createMockRegistrar()
	now := time.Now()

	// Valid attendance marking
	ok := Giveattendence(&reg, 202, 2, "2", true, now)
	if !ok {
		t.Error("Giveattendence failed on valid data")
	}

	records, ok := FetchAttendance(&reg, 202, 2, "2")
	if !ok {
		t.Error("FetchAttendance failed on valid data")
	}
	if val, exists := records[now]; !exists || !val {
		t.Error("Attendance value incorrect or not set")
	}
}

func TestGiveAttendanceInvalid(t *testing.T) {
	reg := createMockRegistrar()
	ok := Giveattendence(&reg, 999, 999, "999", true, time.Now())
	if ok {
		t.Error("Expected Giveattendence to fail on invalid IDs")
	}
}

func TestFetchAttendanceInvalid(t *testing.T) {
	reg := createMockRegistrar()
	_, ok := FetchAttendance(&reg, 999, 999, "999")
	if ok {
		t.Error("Expected FetchAttendance to fail on invalid data")
	}
}

func TestStudentPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid student ID")
		}
	}()
	_ = NewStudent(0, "Invalid") // Should panic
}

func TestDisplayFunctions(t *testing.T) {
	reg := NewRegistrarS{}
	reg.AddStudent(NewStudent(6, "DisplayTest"))
	reg.AddCourse(NewCourse(606, "DisplayCourse"))

	// Just for coverage; visually nothing is tested
	reg.DisplayStudents()
	reg.DisplayCourses()
}

func TestMarkAttendanceNilMap(t *testing.T) {
	att := Attendance{}
	date := time.Now()
	MarkAttendance(&att, date, true)
	if val, ok := att.Records[date]; !ok || !val {
		t.Error("Expected attendance to be true for newly created map entry")
	}
}

func TestSetGraderAndEnrollments(t *testing.T) {
	reg := Registrar{}
	student := NewStudent(10, "Grader Student")
	course := NewCourse(1001, "Graded Course")
	enrollment := NewEnrollment(student, course, nil, 0.85)

	reg.AddStudent(student)
	reg.AddCourse(course)
	reg.Enroll(enrollment)

	// Set Grader
	grader := PercentageGrader{}
	reg.SetGrader(course.Id, grader)

	enrollments := reg.Enrollments()
	if len(enrollments) != 1 {
		t.Error("Expected 1 enrollment")
	}
	grade, err := enrollments[0].Grade(enrollments[0])
	if err != nil {
		t.Error("Error from grader:", err)
	}
	if grade != "85.0%" {
		t.Errorf("Expected grade '85.0%%', got '%s'", grade)
	}
}
func TestLoadStudents_SuccessAndFailure(t *testing.T) {
	// ------------------------ SUCCESS CASE ------------------------
	studentJSON := `[{"id":7,"name":"JSON Student"}]`
	_ = os.WriteFile("students.json", []byte(studentJSON), 0644)

	reg := &Registrar{}
	reg.LoadStudents()

	if len(reg.students) != 1 || reg.students[0].ID() != 7 || reg.students[0].Name() != "JSON Student" {
		t.Error("Failed to correctly load student from JSON")
	}

	_ = os.Remove("students.json") // Clean up

	// ------------------------ FAILURE: FILE NOT FOUND ------------------------
	if os.Getenv("BE_CRASH_STUDENT1") == "1" {
		reg := &Registrar{}
		reg.LoadStudents()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLoadStudents_SuccessAndFailure")
	cmd.Env = append(os.Environ(), "BE_CRASH_STUDENT1=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); !ok || e.ExitCode() == 0 {
		t.Error("Expected LoadStudents to fail due to missing file")
	}

	// ------------------------ FAILURE: INVALID JSON ------------------------
	_ = os.WriteFile("students.json", []byte("not json"), 0644)

	if os.Getenv("BE_CRASH_STUDENT2") == "1" {
		reg := &Registrar{}
		reg.LoadStudents()
		return
	}
	cmd2 := exec.Command(os.Args[0], "-test.run=TestLoadStudents_SuccessAndFailure")
	cmd2.Env = append(os.Environ(), "BE_CRASH_STUDENT2=1")
	err2 := cmd2.Run()
	if e, ok := err2.(*exec.ExitError); !ok || e.ExitCode() == 0 {
		t.Error("Expected LoadStudents to fail due to malformed JSON")
	}

	_ = os.Remove("students.json")
}

func TestLoadCourses_SuccessAndFailure(t *testing.T) {
	// ------------------------ SUCCESS CASE ------------------------
	courseJSON := `[{"id":707,"title":"JSON Course","credits":3}]`
	_ = os.WriteFile("courses.json", []byte(courseJSON), 0644)

	reg := &Registrar{}
	reg.LoadCourses()

	if len(reg.courses) != 1 || reg.courses[0].Id != 707 || reg.courses[0].Name != "JSON Course" {
		t.Error("Failed to correctly load course from JSON")
	}

	_ = os.Remove("courses.json") // Clean up

	// ------------------------ FAILURE: FILE NOT FOUND ------------------------
	if os.Getenv("BE_CRASH_COURSE1") == "1" {
		reg := &Registrar{}
		reg.LoadCourses()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLoadCourses_SuccessAndFailure")
	cmd.Env = append(os.Environ(), "BE_CRASH_COURSE1=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); !ok || e.ExitCode() == 0 {
		t.Error("Expected LoadCourses to fail due to missing file")
	}

	// ------------------------ FAILURE: INVALID JSON ------------------------
	_ = os.WriteFile("courses.json", []byte("bad json"), 0644)

	if os.Getenv("BE_CRASH_COURSE2") == "1" {
		reg := &Registrar{}
		reg.LoadCourses()
		return
	}
	cmd2 := exec.Command(os.Args[0], "-test.run=TestLoadCourses_SuccessAndFailure")
	cmd2.Env = append(os.Environ(), "BE_CRASH_COURSE2=1")
	err2 := cmd2.Run()
	if e, ok := err2.(*exec.ExitError); !ok || e.ExitCode() == 0 {
		t.Error("Expected LoadCourses to fail due to bad JSON")
	}

	_ = os.Remove("courses.json")
}
