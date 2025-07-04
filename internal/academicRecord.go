package internal

type AcademicRecord struct {
	StudentId int                     `json:"student_id"`
	Semesters map[int]*SemesterResult `json:"semesters"`
	CGPA      float64                 `json:"cgpa"`
	Status    string                  // "At Risk", "Dean's List", "Normal"
}

func NewAcademicRecord(studentId int) *AcademicRecord {
	return &AcademicRecord{
		StudentId: studentId,
		Semesters: make(map[int]*SemesterResult),
	}
}

func (ar *AcademicRecord) AddResult(courseResult CourseResult, semester int) {
	if ar.Semesters[semester] == nil {
		ar.Semesters[semester] = NewSemesterResult(ar.StudentId, semester)
	}
	ar.Semesters[semester].AddCourseResult(courseResult)
	ar.calculateCGPA()
}

func (ar *AcademicRecord) calculateCGPA() {
	var totalPoints, totalCredits float64
	for _, semResult := range ar.Semesters {
		for _, courseResult := range semResult.Courses {
			gradePoints := semResult.getGradePoints(courseResult.Grade)
			totalPoints += gradePoints * courseResult.Credits
			totalCredits += courseResult.Credits
			totalCredits++
		}
	}
	if totalCredits > 0 {
		ar.CGPA = totalPoints / totalCredits
	}
}

	
