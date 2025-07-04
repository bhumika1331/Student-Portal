package main

import (
	"fmt"
	"oops/main/infrastructure"
	"oops/main/internal"
	"time"
)

var courseResults []internal.CourseResult

func main() {

	registrar := internal.Registrar{}

	registrar.LoadCourses()
	registrar.DisplayCourses()

	registrar.LoadStudents()
	registrar.DisplayStudents()
	courseResults = infrastructure.LoadCourseResults()

	fmt.Println("======================")
	fmt.Println()

	fmt.Print(courseResults)
	Drive := internal.NewDrive(time.Date(2025, time.July, 4, 14, 30, 0, 0, time.UTC), time.Date(2025, time.July, 18, 14, 30, 0, 0, time.UTC), "Java Developer", 5.0, 50000, internal.Dream)
	placReg := internal.PlacementRegistrar{}
	comp := internal.Company{}
	placReg.AddCompany(comp)
	comp.AddDrive(Drive)
	fmt.Println(placReg)

}
