package infrastructure

import (
	"encoding/csv"
	"fmt"
	"oops/main/internal"
	"os"
	"strconv"
	"encoding/json"
    "bytes"
)

var deanListStudents []internal.AcademicRecord
var atRiskStudents []internal.AcademicRecord

func ExportTranscript(path string, list []internal.Enrollment) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	_ = w.Write([]string{"Student", "Course", "Grade"})
	for _, e := range list {
		grade, _ := e.Grade(e)
		_ = w.Write([]string{e.Student.Name(), e.Course.Name, grade})
	}
	return w.Error()
}


// func hello() {
// 	var atRiskStudents []internal.AcademicRecord
// 	fmt.Print(atRiskStudents)
// }


func ExportAtRiskStudents(path string, internal []internal.AcademicRecord) error {

	for _, student := range internal {
		if student.Status == "At Risk" {
			atRiskStudents = append(atRiskStudents, student)
		}
	}
	return ExportSummaryReport(path, atRiskStudents)
}

func ExportDeanListStudents(path string, internal []internal.AcademicRecord) error {

	for _, student := range internal {
		if student.Status == "Dean's List" {
			deanListStudents = append(deanListStudents, student)
		}
	}
	return ExportSummaryReport(path, deanListStudents)
}

func ExportSummaryReport(path string, internal []internal.AcademicRecord) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Student_ID", "Student_Name", "Total_Semesters", "Overall_GPA", "Status"}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, student := range internal {
		record := []string{
			strconv.Itoa(student.StudentId),
			//student.Student.Name(),
			strconv.Itoa(len(student.Semesters)),
			fmt.Sprintf("%.2f", student.CGPA),
			student.Status,
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func ExportResultsAsJSON(results []internal.StudentResult) ([]byte, error) {
    return json.MarshalIndent(results, "", "  ")
}


func ExportResultsAsCSV(results []internal.StudentResult) ([]byte, error) {
    var buf bytes.Buffer
    writer := csv.NewWriter(&buf)
    writer.Write([]string{"course_id", "course_name", "student_id", "student_name", "score", "grade"})
    for _, r := range results {
        writer.Write([]string{
            strconv.Itoa(r.CourseID),
            r.CourseName,
            strconv.Itoa(r.StudentID),
            r.StudentName,
            fmt.Sprintf("%.2f", r.Score),
            r.Grade,
        })
    }
    writer.Flush()
    return buf.Bytes(), writer.Error()
}
