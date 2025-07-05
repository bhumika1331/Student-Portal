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
	placReg.AddCompany(&comp)
	comp.AddDrive(Drive)
	fmt.Println(placReg)

	// Run GPA Histogram Analysis
	var hist map[string]int
	hist, err := internal.GenerateGPAHistogramFromFiles("courseResults.json", "students.json")
	if err != nil {
		fmt.Println("Failed to generate GPA histogram:", err)
	} else {
		//fmt.Println("GPA Histogram generated:", hist)
		err = internal.ExportGPAHistogramChart(hist, "gpa_histogram.png")
		if err != nil {
			fmt.Println("Failed to export histogram chart:", err)
		} else {
			fmt.Println("GPA Trends Chart Generated")
		}
	}
	// Dean List Chart
	if err := internal.ExportDeanListChart("courseResults.json", "students.json", "dean_list.png"); err != nil {
		fmt.Println("Dean List chart export failed:", err)
	} else {
		fmt.Println("Dean List Chart Generated.")
	}

	// At-Risk Students Chart
	if err := internal.ExportAtRiskChart("courseResults.json", "students.json", "at_risk_students.png"); err != nil {
		fmt.Println("At-Risk chart export failed:", err)
	} else {
		fmt.Println("At-Risk Chart Generated.")
	}

	// Placement offer categorization
	offers, err := internal.LoadOffers("placement_offers.json")
	if err != nil {
		fmt.Println("Failed to load offers:", err)
	} else {
		categorized := internal.CategorizeOffers(offers)
		err = internal.ExportCategorizedOffers("placement_chart.json", categorized)
		if err != nil {
			fmt.Println("Export failed:", err)
		}
		//} else {
		//	fmt.Print("Offers categorized and saved to placement_chart.json")
		//}
	}

	// Export placement bar chart
	err = internal.ExportPlacementBarChart("placement_chart.json", "placement_chart.png")
	if err != nil {
		fmt.Println("Placement chart export failed:", err)
	} else {
		fmt.Println("Placement Chart Generated.")
	}

	//Company Wise Selection Metrics
	if err := internal.ExportCompanySelectionChart("placement_offers.json", "company_selection.png", "company_selection.json"); err != nil {
		fmt.Println("Company Selection chart export failed:", err)
	} else {
		fmt.Println("Company Selection Chart Generated.")
	}

}
