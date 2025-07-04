package internal

import (
	"testing"
	"time"
)

func TestMarkAttendance(t *testing.T) {
	registrar := &NewRegistrarS{}
	teacher := NewTeacher("T1", "Alice")
	student := NewStudent(1, "Bob")
	course := NewCourse(101, "Math")
	grader := LetterGrader{}

	enroll := NewEnrollNew(student, course, grader, 0, Attendance{}, teacher)
	registrar.Enrollnew(enroll)

	date := time.Now()
	Giveattendence(registrar, 101, 1, "T1", true, date)

	found := false
	for _, e := range registrar.enroll {
		if e.Attend.Records[date] {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("attendance not marked correctly")
	}
}
