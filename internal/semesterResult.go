package internal

import "errors"

type AlphabeticGrade uint8

const (
	O AlphabeticGrade = iota
	Aplus
	A
	Bplus
	B
	C
	F
)

var alphabeticGradeStrings = map[AlphabeticGrade]string{
	O:     "O",
	Aplus: "A+",
	A:     "A",
	Bplus: "B+",
	B:     "B",
	C:     "C",
	F:     "F",
}

func (a AlphabeticGrade) String() string {
	return alphabeticGradeStrings[a]
}

type SemesterResult struct {
	Semester  int                  `json:"semester"`
	StudentId int                  `json:"student_id"`
	Courses   map[int]CourseResult `json:"courses"`
	SGPA      float64              `json:"sgpa"`
}

func NewSemesterResult(studentId, semester int) *SemesterResult {
	return &SemesterResult{
		StudentId: studentId,
		Semester:  semester,
		Courses:   make(map[int]CourseResult),
	}
}

func (pr *SemesterResult) SetSemester(sem int) error {
	if pr.Semester == 0 {
		pr.Semester = sem
	}

	return errors.New("cannot change an already set semester")

}

func (sr *SemesterResult) AddCourseResult(result CourseResult) {
	sr.Courses[result.CourseId] = result
	sr.calculateSGPA()
}

func (sr *SemesterResult) calculateSGPA() {
	var totalPoints, totalCredits float64

	for _, result := range sr.Courses {
		gradePoints := sr.getGradePoints(result.Grade)
		totalPoints += gradePoints * result.Credits
		totalCredits += result.Credits
	}

	if totalCredits > 0 {
		sr.SGPA = totalPoints / totalCredits
	}
}

func (sr *SemesterResult) getGradePoints(grade AlphabeticGrade) float64 {
	gradePoints := map[AlphabeticGrade]float64{
		O:     10.0,
		Aplus: 9.0,
		A:     8.0,
		Bplus: 7.0,
		B:     6.0,
		C:     5.0,
		F:     0.0,
	}
	return gradePoints[grade]
}
