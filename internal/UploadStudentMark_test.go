package internal

import (
	"fmt"
	"os"
	"testing"
)

func TestUploadStudentMarksFromJSON(t *testing.T) {
	// Setup: Registrar, Teacher, Course, Grader
	registrar := &RegistrarWithDocs{NewRegistrarS: &NewRegistrarS{}}
	teacher := NewTeacher("T1", "Alice")
	course := NewCourse(101, "Math")
	grader := LetterGrader{}

	// Add teacher and course mapping
	registrar.NewRegistrarS.AddTeacher(teacher)
	registrar.NewRegistrarS.AddTeacherenrollment(NewTeacherEnrollment(teacher, CreditCourse{Course: course, Credit: 4}))

	// Enroll students 1-50 in course 101 with teacher T1
	for i := 1; i <= 50; i++ {
		student := NewStudent(i, fmt.Sprintf("Student%d", i))
		enroll := NewEnrollNew(student, course, grader, 0, Attendance{}, teacher)
		registrar.NewRegistrarS.Enrollnew(enroll)
	}

	ts := &TeacherService{Registrar: registrar, Teacher: teacher}

	// Load marks from JSON file
	data, err := os.ReadFile("marks.json")
	if err != nil {
		t.Fatalf("failed to read marks.json: %v", err)
	}
	err = ts.UploadStudentMarksFromJSON(data)
	if err != nil {
		t.Errorf("unexpected error uploading marks: %v", err)
	}

	// Check that marks are uploaded for all enrolled students
	for i := 1; i <= 50; i++ {
		found := false
		for _, e := range registrar.NewRegistrarS.enroll {
			if e.Student.ID() == i && e.Course.Id == 101 && e.Teacher.TID() == "T1" {
				if e.score == 0 {
					t.Errorf("score not uploaded for student %d", i)
				}
				found = true
				break
			}
		}
		if !found {
			t.Errorf("enrollment not found for student %d", i)
		}
	}
}
