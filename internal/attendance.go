package internal

import (
	"time"
)

type Attendance struct {
	Records map[time.Time]bool
}

func markAttendance(attendance *Attendance, date time.Time, present bool) {
	if attendance.Records == nil {
		attendance.Records = make(map[time.Time]bool)
	}
	attendance.Records[date] = present
}

// need to check time data type parameter passing
func Giveattendence(r *NewRegistrarS, courseID int, studentID int, TeacherID string, attendence bool, t time.Time) {
	for i := range r.enroll {
		e := &r.enroll[i]
		if e.Course.Id == courseID && e.Student.ID() == studentID && e.Teacher.TID() == TeacherID {
			markAttendance(&e.Attend, t, attendence)
			return
		}
	}
}
