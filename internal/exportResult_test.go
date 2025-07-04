package internal

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func (ts *TeacherService) ExportMarksForCourseJSON(courseID int) ([]byte, error) {
	results, err := ts.GetCourseResults(courseID)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(results, "", "  ")
}

func (ts *TeacherService) ExportMarksForCourseCSV(courseID int) (string, error) {
	results, err := ts.GetCourseResults(courseID)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	// Write header
	writer.Write([]string{"StudentID", "StudentName", "CourseID", "CourseName", "Score", "Grade"})
	// Write data rows
	for _, r := range results {
		writer.Write([]string{
			strconv.Itoa(r.StudentID),
			r.StudentName,
			strconv.Itoa(r.CourseID),
			r.CourseName,
			fmt.Sprintf("%.2f", r.Score),
			r.Grade, // Use directly, no .String()
		})

	}
	writer.Flush()
	return buf.String(), nil
}

func TestExportMarksForCourseJSON(t *testing.T) {
	registrar := &RegistrarWithDocs{NewRegistrar: &NewRegistrar{}}
	teacher := NewTeacher("T1", "Alice")
	course := NewCourse(101, "Math")
	grader := LetterGrader{}

	// Add teacher and course mapping
	registrar.NewRegistrar.AddTeacher(teacher)
	registrar.NewRegistrar.AddTeacherenrollment(NewTeacherEnrollment(teacher, CreditCourse{Course: course, Credit: 4}))

	// Enroll students 1-50 in course 101 with teacher T1 and assign scores
	for i := 1; i <= 50; i++ {
		student := NewStudent(i, fmt.Sprintf("Student%d", i))
		enroll := NewEnrollNew(student, course, grader, float64(50+i), Attendance{}, teacher)
		registrar.NewRegistrar.Enrollnew(enroll)
	}

	ts := &TeacherService{Registrar: registrar, Teacher: teacher}

	// Export as JSON
	data, err := ts.ExportMarksForCourseJSON(101)
	if err != nil {
		t.Fatalf("export failed: %v", err)
	}

	var results []StudentResult
	if err := json.Unmarshal(data, &results); err != nil {
		t.Fatalf("failed to unmarshal exported JSON: %v", err)
	}
	if len(results) != 50 {
		t.Errorf("expected 50 results, got %d", len(results))
	}
	for _, r := range results {
		if r.CourseID != 101 {
			t.Errorf("unexpected course id in result: %d", r.CourseID)
		}
	}
}

func TestExportMarksForCourseCSV(t *testing.T) {
	registrar := &RegistrarWithDocs{NewRegistrar: &NewRegistrar{}}
	teacher := NewTeacher("T1", "Alice")
	course := NewCourse(101, "Math")
	grader := LetterGrader{}

	// Add teacher and course mapping
	registrar.NewRegistrar.AddTeacher(teacher)
	registrar.NewRegistrar.AddTeacherenrollment(NewTeacherEnrollment(teacher, CreditCourse{Course: course, Credit: 4}))

	// Enroll students 1-50 in course 101 with teacher T1 and assign scores
	for i := 1; i <= 50; i++ {
		student := NewStudent(i, fmt.Sprintf("Student%d", i))
		enroll := NewEnrollNew(student, course, grader, float64(50+i), Attendance{}, teacher)
		registrar.NewRegistrar.Enrollnew(enroll)
	}

	ts := &TeacherService{Registrar: registrar, Teacher: teacher}

	// Export as CSV
	csvData, err := ts.ExportMarksForCourseCSV(101)
	if err != nil {
		t.Fatalf("export failed: %v", err)
	}
	fmt.Println(csvData)
	if !strings.Contains(csvData, "StudentID") || !strings.Contains(csvData, "Math") {
		t.Errorf("CSV export missing headers or course name")
	}
	// Optionally, check for the number of rows
	lines := strings.Split(strings.TrimSpace(csvData), "\n")
	if len(lines) != 51 { // 1 header + 50 students
		t.Errorf("expected 51 lines in CSV, got %d", len(lines))
	}
}
