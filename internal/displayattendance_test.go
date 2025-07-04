package internal

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func setupDisplayAttendanceTest() (*TeacherService, int, int, string, time.Time) {
	studentID := 1001
	courseID := 101
	teacherID := "T001"
	student := NewStudent(studentID, "Alice")
	course := NewCourse(courseID, "Math")
	teacher := NewTeacher(teacherID, "Prof. Smith")
	att := Attendance{Records: make(map[time.Time]bool)}
	date := time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC)
	att.Records[date] = true
	gr := PercentageGrader{}
	reg := &NewRegistrarS{}
	enroll := NewEnrollNew(student, course, gr, 0.85, att, teacher)
	reg.Enrollnew(enroll)
	reg.AddTeacher(teacher)
	regWithDocs := &RegistrarWithDocs{NewRegistrarS: reg}
	ts := &TeacherService{Registrar: regWithDocs, Teacher: teacher}
	return ts, courseID, studentID, teacherID, date
}

func TestDisplayAttendance_Found(t *testing.T) {
	ts, courseID, studentID, _, date := setupDisplayAttendanceTest()

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ts.DisplayAttendance(courseID, studentID)

	w.Close()
	os.Stdout = stdout
	outputBytes := make([]byte, 1024)
	n, _ := r.Read(outputBytes)
	output := string(outputBytes[:n])

	expectedHeader := fmt.Sprintf("Attendance for student %d in course %d:\n", studentID, courseID)
	expectedRecord := fmt.Sprintf("%s -> true\n", date.Format("2006-01-02"))

	if !strings.Contains(output, expectedHeader) {
		t.Errorf("expected header %q in output, got %q", expectedHeader, output)
	}
	if !strings.Contains(output, expectedRecord) {
		t.Errorf("expected attendance record %q in output, got %q", expectedRecord, output)
	}
}

func TestDisplayAttendance_NotFound(t *testing.T) {
	ts, courseID, _, _, _ := setupDisplayAttendanceTest()

	nonExistentStudentID := 9999

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ts.DisplayAttendance(courseID, nonExistentStudentID)

	w.Close()
	os.Stdout = stdout
	outputBytes := make([]byte, 1024)
	n, _ := r.Read(outputBytes)
	output := string(outputBytes[:n])

	expected := "No attendance records found.\n"
	if !strings.Contains(output, expected) {
		t.Errorf("expected %q in output, got %q", expected, output)
	}
}
