package internal

import (
	"errors"
	"fmt"
	"time"
)

func (ts *TeacherService) UploadFile(
	courseID int,
	studentID int,
	title string,
	filename string,
	mimeType string,
	content []byte,
) error {
	for _, e := range ts.Registrar.NewRegistrarS.enroll {
		if e.Course.Id == courseID &&
			e.Student.ID() == studentID &&
			e.Teacher.TID() == ts.Teacher.TID() {

			doc := Document{
				Title:      title,
				Filename:   filename,
				Content:    content,
				MimeType:   mimeType,
				UploadedAt: time.Now(),
			}

			enrollWithDocs := EnrollnewWithDocs{
				EnrollNew: e,
				Documents: []Document{doc},
			}

			ts.Registrar.EnrollnewWithDocs(enrollWithDocs)

			fmt.Printf("File uploaded: %s (%s) for student %d in course %d by teacher %s\n",
				filename, mimeType, studentID, courseID, ts.Teacher.TID())

			return nil
		}
	}

	return errors.New("no valid enrollment found for this teacher, student, and course")
}
