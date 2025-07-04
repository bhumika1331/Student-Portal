package internal

import (
	"testing"
	"time"
)

func setupTeacherTestEnv() (*TeacherService, int, int, string) {
	// IDs
	studentID := 1001
	courseID := 101
	teacherID := "T001"

	// Core objects
	student := NewStudent(studentID, "Alice")
	course := NewCourse(courseID, "Math")
	teacher := NewTeacher(teacherID, "Prof. Smith")
	att := Attendance{Records: make(map[time.Time]bool)}
	gr := PercentageGrader{}

	// Registrar setup
	reg := &NewRegistrar{}
	enroll := NewEnrollNew(student, course, gr, 0.85, att, teacher)
	reg.Enrollnew(enroll)
	reg.AddTeacher(teacher)
	regWithDocs := &RegistrarWithDocs{NewRegistrar: reg}

	ts := &TeacherService{Registrar: regWithDocs, Teacher: teacher}
	return ts, courseID, studentID, teacherID
}

func TestTeacherUploadFile(t *testing.T) {
	ts, courseID, studentID, _ := setupTeacherTestEnv()
	err := ts.UploadFile(courseID, studentID, "Assignment 1", "assignment1.pdf", "application/pdf", []byte("dummy"))
	if err != nil {
		t.Fatalf("expected upload to succeed, got error: %v", err)
	}
	// Check document attached
	found := false
	for _, e := range ts.Registrar.enrollWithDocs {
		if e.Student.ID() == studentID && e.Course.Id == courseID {
			for _, doc := range e.Documents {
				if doc.Title == "Assignment 1" && doc.Filename == "assignment1.pdf" {
					found = true
				}
			}
		}
	}
	if !found {
		t.Error("uploaded document not found in registrar")
	}
}

// func TestTeacherUploadStudentMarksFromJSON(t *testing.T) {
// 	ts, courseID, studentID, _ := setupTeacherTestEnv()

// 	// Prepare JSON input for bulk upload
// 	marks := []StudentMarkInput{
// 		{CourseID: courseID, StudentID: studentID, Score: 0.88},
// 	}
// 	jsonData, err := json.Marshal(marks)
// 	if err != nil {
// 		t.Fatalf("failed to marshal marks: %v", err)
// 	}

// 	// Call the bulk upload method
// 	err = ts.UploadStudentMarksFromJSON(jsonData)
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	// Assert that the score was updated
// 	found := false
// 	for _, e := range ts.Registrar.enroll {
// 		if e.Course.Id == courseID && e.Student.ID() == studentID {
// 			found = true
// 			if e.score != 0.88 {
// 				t.Fatalf("got score %v want 0.88", e.score)
// 			}
// 			grade, _ := e.Grader.Grade(e.Enrollment)
// 			if grade != "88.0%" {
// 				t.Fatalf("got grade %s want 88.0%%", grade)
// 			}
// 		}
// 	}
// 	if !found {
// 		t.Fatal("enrollment not found")
// 	}
// }

// func TestTeacherGetCourseResults(t *testing.T) {
// 	ts, courseID, studentID, _ := setupTeacherTestEnv()

// 	results, err := ts.GetCourseResults(courseID)
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}
// 	if len(results) != 1 {
// 		t.Fatalf("expected 1 result, got %d", len(results))
// 	}
// 	got := results[0]
// 	if got.CourseID != courseID {
// 		t.Errorf("got CourseID %d, want %d", got.CourseID, courseID)
// 	}
// 	if got.StudentID != studentID {
// 		t.Errorf("got StudentID %d, want %d", got.StudentID, studentID)
// 	}
// 	if got.Score != 0.85 {
// 		t.Errorf("got Score %v, want 0.85", got.Score)
// 	}
// 	if got.Grade != "85.0%" {
// 		t.Errorf("got Grade %s, want 85.0%%", got.Grade)
// 	}
// }
