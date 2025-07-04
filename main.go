package main

import (
	"fmt"
	"oops/main/internal"
	"oops/main/infrastructure"
	// "oops/main/internal/admin"
)

var courseResults []internal.CourseResult

func main() {
	regitrar := internal.Registrar{}

	regitrar.LoadCourses()
	regitrar.DisplayCourses()

	regitrar.LoadStudents()
	regitrar.DisplayStudents()
	
	courseResults = infrastructure.LoadCourseResults()

	fmt.Println("======================")
	fmt.Println()

	// Teacher 2 views Bob
	// ts2 := &internal.TeacherService{
	// 	Teacher:   teacher2,
	// 	Registrar: registrar,
	// }
	// fmt.Println("Teacher:", teacher2.Name)
	// ts2.DisplayAttendance(102, 1002)
	// fmt.Println()

	
}

