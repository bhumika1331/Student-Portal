package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"oops/main/internal"
	"os"
)


type courseResultData struct {
	StudentID  int     `json:"student_id"`
	CourseID   int     `json:"course_id"`
	CourseName string  `json:"course_name"`
	Grade      string  `json:"grade"`
	Semester   int     `json:"semester"`
	Credits    float64 `json:"credits"`
}

func parseGrade(gradeStr string) (internal.AlphabeticGrade, error) {
	gradeMap := map[string]internal.AlphabeticGrade{
		"O":  internal.O,
		"A+": internal.Aplus,
		"A":  internal.A,
		"B+": internal.Bplus,
		"B":  internal.B,
		"C":  internal.C,
		"F":  internal.F,
	}

	if grade, exists := gradeMap[gradeStr]; exists {
		return grade, nil
	}
	return internal.F, fmt.Errorf("invalid grade: %s", gradeStr)
}

func LoadCourseResults() []internal.CourseResult {
	var resultsData []courseResultData
	data, err := os.ReadFile("courseResults.json")
	if err != nil {
		log.Fatal("Failed to read courseResults.json:", err)
	}

	err = json.Unmarshal(data, &resultsData)
	if err != nil {
		log.Fatal("Failed to unmarshal course results:", err)
	}

	var results []internal.CourseResult
	for _, rd := range resultsData {
		grade, err := parseGrade(rd.Grade)
		if err != nil {
			log.Printf("Warning: %v for student ID %d", err, rd.StudentID)
			continue
		}
		courseResult := internal.NewCourseResult(rd.StudentID, rd.CourseID, rd.CourseName, grade, rd.Semester, rd.Credits)
		results = append(results, courseResult)
	}
	return results
}
/*
func (ar *AcademicRecordGenerator) GetAtRiskStudents() []internal.AcademicRecord { 
	var atRisk []internal.AcademicRecord
	for _, student := range ar.internal {
		if student.Status == "At Risk" {
			atRisk = append(atRisk, student)
		}
	}
	return atRisk
}

func (ar *AcademicRecordGenerator) GetDeanListStudents() []internal.AcademicRecord{ // Here i have added (ar *AcademicRecordGenerator to try to map it with GPAReportGenerator)
	var deanList []internal.AcademicRecord
	for _, student := range ar.internal {
		if student.Status == "Dean's List" {
			deanList = append(deanList, student)
		}
	}
	return deanList
}
*/
