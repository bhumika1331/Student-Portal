package main

import (
	"fmt"
	"strings"
	"time"

	"oops/main/infrastructure"
	"oops/main/internal"
)

func main() {
	fmt.Println("=== Student Management System Demo ===\n")

	// Phase 1: Setup and Basic Entity Creation
	demonstrateBasicSetup()

	// Phase 2: Academic Management  
	demonstrateAcademicManagement()

	// Phase 3: Teacher Services and Document Management
	demonstrateTeacherServices()

	// Phase 4: Placement Management
	demonstratePlacementSystem()

	// Phase 5: Analytics and Reporting
	demonstrateAnalytics()

	// Phase 6: File Operations
	demonstrateFileOperations()

	fmt.Println("\n=== Demo Complete ===")
}

func demonstrateBasicSetup() {
	fmt.Println("1. BASIC SETUP - Loading Data and Creating System")
	fmt.Println(strings.Repeat("=", 50))

	// Setup registrar and load data from JSON files
	registrar := &internal.NewRegistrarS{}
	
	// Load students from JSON file
	registrar.LoadStudents()
	fmt.Println("✅ Loaded students from students.json")
	
	// Load courses from JSON file  
	registrar.LoadCourses()
	fmt.Println("✅ Loaded courses from courses.json")

	// Only hardcode teachers (not imported from JSON)
	teachers := []internal.Teacher{
		internal.NewTeacher("T001", "Dr. Alan Turing"),
		internal.NewTeacher("T002", "Dr. Ada Lovelace"), 
		internal.NewTeacher("T003", "Dr. Grace Hopper"),
		internal.NewTeacher("T004", "Dr. Tim Berners-Lee"),
		internal.NewTeacher("T005", "Dr. Linus Torvalds"),
		internal.NewTeacher("T006", "Dr. Margaret Hamilton"),
		internal.NewTeacher("T007", "Dr. Donald Knuth"),
		internal.NewTeacher("T008", "Dr. Barbara Liskov"),
		internal.NewTeacher("T009", "Dr. John McCarthy"),
		internal.NewTeacher("T010", "Dr. Edsger Dijkstra"),
	}

	fmt.Println("\nCreated Teachers:")
	for _, t := range teachers {
		registrar.AddTeacher(t)
		fmt.Printf("Teacher %s: %s\n", t.TID(), t.Name)
	}

	// Display loaded data summary
	registrar.DisplayStudents()
	registrar.DisplayCourses()

	fmt.Println("\nRegistrar setup complete with imported data!")
}

func demonstrateAcademicManagement() {
	fmt.Println("\n2. ACADEMIC MANAGEMENT - Using Imported Course Results")
	fmt.Println(strings.Repeat("=", 60))

	// Load course results from JSON file (not hardcoded)
	courseResults := infrastructure.LoadCourseResults()
	fmt.Printf("✅ Loaded %d course results from courseResults.json\n", len(courseResults))

	// Demo grading systems with sample enrollments
	student1 := internal.NewStudent(1, "Alice Johnson") 
	course1 := internal.NewCourse(101, "Data Structures")

	letterGrader := internal.LetterGrader{}
	percentageGrader := internal.PercentageGrader{}
	passFailGrader := internal.PassFailGrader{PassMark: 0.6}

	enrollment1 := internal.NewEnrollment(student1, course1, letterGrader, 8.5)
	enrollment2 := internal.NewEnrollment(student1, course1, percentageGrader, 0.75)

	grade1, _ := letterGrader.Grade(enrollment1)
	grade2, _ := percentageGrader.Grade(enrollment2) 
	grade3, _ := passFailGrader.Grade(enrollment1)

	fmt.Printf("\nGrading System Demo:\n")
	fmt.Printf("- Letter Grade (8.5/10): %s\n", grade1)
	fmt.Printf("- Percentage Grade (0.75): %s\n", grade2)
	fmt.Printf("- Pass/Fail Status: %s\n", grade3)

	// Attendance management demo
	attendance := internal.Attendance{Records: make(map[time.Time]bool)}
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	internal.MarkAttendance(&attendance, today, true)
	internal.MarkAttendance(&attendance, yesterday, false)

	fmt.Println("\nAttendance Demo:")
	for date, present := range attendance.Records {
		status := "Absent"
		if present {
			status = "Present"
		}
		fmt.Printf("- %s: %s\n", date.Format("2006-01-02"), status)
	}

	// Process imported course results into academic records
	demonstrateGPACalculationWithImportedData(courseResults)
}

func demonstrateGPACalculationWithImportedData(courseResults []internal.CourseResult) {
	fmt.Println("\n--- GPA Calculation with Imported Data ---")

	// Create academic records from imported course results
	records := make(map[int]*internal.AcademicRecord)
	
	for _, cr := range courseResults {
		if records[cr.StudentId] == nil {
			records[cr.StudentId] = internal.NewAcademicRecord(cr.StudentId)
		}
		records[cr.StudentId].AddResult(cr, cr.Semester)
	}

	// Show sample academic records
	fmt.Printf("Sample Academic Records from Imported Data:\n")
	count := 0
	for studentID, record := range records {
		if count >= 5 { // Show first 5 students
			break
		}
		fmt.Printf("Student %d: CGPA=%.2f, Status=%s\n", 
			studentID, record.CGPA, record.Status)
		count++
	}

	// GPA Calculator demo with real data
	calculator := internal.NewGPACalculator()
	if len(records) > 0 {
		// Use first student's data
		for studentID, record := range records {
			student := internal.NewStudent(studentID, fmt.Sprintf("Student %d", studentID))
			semesterGPAs := []internal.StudentGPA{
				{Student: student, Semester: 1, Gpa: record.CGPA},
				{Student: student, Semester: 2, Gpa: record.CGPA},
			}
			
			overallGPA := calculator.CalculateOverallGPA(semesterGPAs)
			status := calculator.DetermineStatus(overallGPA)
			
			fmt.Printf("\nGPA Calculator Validation:\n")
			fmt.Printf("- Calculated GPA: %.2f\n", overallGPA)
			fmt.Printf("- Determined Status: %s\n", status)
			break // Just demo with first student
		}
	}
}

func demonstrateTeacherServices() {
	fmt.Println("\n3. TEACHER SERVICES - Document Upload and Mark Management")
	fmt.Println(strings.Repeat("=", 60))

	registrar := &internal.RegistrarWithDocs{
		NewRegistrarS: &internal.NewRegistrarS{},
	}

	// Load students and courses from JSON
	registrar.LoadStudents()
	registrar.LoadCourses()

	// Create teachers (not imported from JSON)
	teachers := []internal.Teacher{
		internal.NewTeacher("T001", "Dr. Alan Turing"),
		internal.NewTeacher("T002", "Dr. Ada Lovelace"),
		internal.NewTeacher("T003", "Dr. Grace Hopper"),
	}

	for _, teacher := range teachers {
		registrar.AddTeacher(teacher)
	}

	// Create sample students and courses for demo (using first few from loaded data)
	sampleStudents := []internal.Student{
		internal.NewStudent(1, "Alice Johnson"),
		internal.NewStudent(2, "Bob Smith"), 
		internal.NewStudent(3, "Charlie Brown"),
	}

	sampleCourses := []internal.Course{
		internal.NewCourse(101, "Data Structures"),
		internal.NewCourse(102, "Operating Systems"),
	}

	// Create enrollments for teacher services demo
	registrar.CreateEnrollmentsForAllStudents(sampleStudents, sampleCourses, teachers)
	fmt.Printf("✅ Created enrollments for teacher services demo\n")

	// Teacher service demo
	teacherService := &internal.TeacherService{
		Registrar: registrar,
		Teacher:   teachers[0],
	}

	// Upload marks demo
	err := teacherService.UploadStudentMark(101, 1, 85.5)
	if err != nil {
		fmt.Printf("Error uploading mark: %v\n", err)
	} else {
		fmt.Println("✅ Successfully uploaded individual mark")
	}

	// Bulk marks upload demo
	marksJSON := `[
		{"course_id": 101, "student_id": 1, "score": 87.5},
		{"course_id": 101, "student_id": 2, "score": 92.0}
	]`

	err = teacherService.UploadStudentMarksFromJSON([]byte(marksJSON))
	if err != nil {
		fmt.Printf("Error uploading bulk marks: %v\n", err)
	} else {
		fmt.Println("✅ Successfully uploaded bulk marks")
	}

	// Document upload demo
	sampleDocument := []byte("Sample assignment content")
	err = teacherService.UploadFile(101, 1, "Assignment 1", "assignment1.pdf", "application/pdf", sampleDocument)
	if err != nil {
		fmt.Printf("Error uploading document: %v\n", err)
	} else {
		fmt.Println("✅ Successfully uploaded document")
	}

	registrar.DisplayDocuments()
}

func demonstratePlacementSystem() {
	fmt.Println("\n4. PLACEMENT MANAGEMENT - Companies, Drives, and Applications")
	fmt.Println(strings.Repeat("=", 65))

	placementRegistrar := &internal.PlacementRegistrar{}

	// Create companies (not imported from JSON - business logic)
	companyNames := []string{
		"Google", "Microsoft", "Amazon", "Meta", "Apple",
		"Netflix", "Adobe", "Salesforce", "Oracle", "IBM",
		"TCS", "Infosys", "Wipro", "Accenture", "Capgemini",
	}

	companies := make([]*internal.Company, len(companyNames))
	for i, name := range companyNames {
		companies[i] = internal.NewCompany(name)
		placementRegistrar.AddCompany(companies[i])
	}
	fmt.Printf("✅ Created %d companies\n", len(companies))

	// Create placement drives (not imported from JSON)
	startDate := time.Now()
	endDate := startDate.AddDate(0, 1, 0)

	driveConfigs := []struct {
		companyIndex int
		role         string
		minGPA       float64
		ctc          int
		category     internal.JobCategory
	}{
		{0, "Software Engineer", 8.5, 3000000, internal.Marquee},
		{1, "SDE-2", 8.0, 2800000, internal.Marquee},
		{2, "SDE-1", 7.5, 2500000, internal.SuperDream},
		{3, "Software Engineer", 8.8, 3200000, internal.Marquee},
		{4, "iOS Developer", 8.0, 2900000, internal.Marquee},
		{10, "Associate", 6.0, 900000, internal.Dream},
		{11, "Systems Engineer", 6.5, 950000, internal.Dream},
		{12, "Project Engineer", 6.0, 800000, internal.Dream},
	}

	for _, config := range driveConfigs {
		drive := internal.NewDrive(startDate, endDate, config.role, config.minGPA, config.ctc, config.category)
		companies[config.companyIndex].AddDrive(drive)
		fmt.Printf("✅ %s: %s (%.1f LPA, min GPA: %.1f)\n", 
			companies[config.companyIndex].Name(), config.role,
			float64(config.ctc)/100000, config.minGPA)
	}

	// Create applicants using imported student data
	createApplicantsFromImportedData(placementRegistrar)

	// Demo placement process
	fmt.Println("\n--- Placement Process Demo ---")
	err := placementRegistrar.ApplyForDrive(1, 1, 1)
	if err != nil {
		fmt.Printf("Application error: %v\n", err)
	} else {
		fmt.Println("✅ Successfully applied for drive")
	}

	err = placementRegistrar.UpdateApplicationStatus(1, 1, internal.ShortListed)
	if err != nil {
		fmt.Printf("Status update error: %v\n", err)
	} else {
		fmt.Println("✅ Application status updated to ShortListed")
	}

	// Generate reports
	reportByStudent := placementRegistrar.GenerateReportByStudent()
	reportByDrive := placementRegistrar.GenerateReportByDrive()
	fullReport := placementRegistrar.GenerateFullReport()

	fmt.Printf("\nPlacement Statistics:\n")
	fmt.Printf("- Total companies: %d\n", fullReport.TotalComapanies)
	fmt.Printf("- Total offers made: %d\n", fullReport.TotalOffersMade)
	fmt.Printf("- Students with reports: %d\n", len(reportByStudent))
	fmt.Printf("- Drive CTC: ₹%.1f LPA\n", float64(reportByDrive.DriveCTC)/100000)
}

func createApplicantsFromImportedData(placementRegistrar *internal.PlacementRegistrar) {
	// Load course results to calculate CGPAs
	courseResults := infrastructure.LoadCourseResults()
	
	// Build academic records from imported data
	records := make(map[int]*internal.AcademicRecord)
	for _, cr := range courseResults {
		if records[cr.StudentId] == nil {
			records[cr.StudentId] = internal.NewAcademicRecord(cr.StudentId)
		}
		records[cr.StudentId].AddResult(cr, cr.Semester)
	}

	// Create applicants from first 20 students with academic records
	count := 0
	for studentID, record := range records {
		if count >= 20 { // Limit for demo
			break
		}
		
		student := internal.NewStudent(studentID, fmt.Sprintf("Student %d", studentID))
		applicant := internal.NewApplicant(student, *record)
		placementRegistrar.AddApplicant(applicant)
		
		fmt.Printf("✅ Added applicant: Student %d (CGPA: %.1f, Status: %s)\n", 
			studentID, record.CGPA, record.Status)
		count++
	}
}

func demonstrateAnalytics() {
	fmt.Println("\n5. ANALYTICS - Real Data Processing and Visualization")
	fmt.Println(strings.Repeat("=", 60))

	// Use imported data for analytics (not hardcoded)
	histogram, err := internal.GenerateGPAHistogramFromFiles("courseResults.json", "students.json")
	if err != nil {
		fmt.Printf("Error generating histogram: %v\n", err)
		return
	}

	fmt.Println("✅ Generated GPA histogram from imported data!")
	fmt.Println("\nGPA Distribution:")
	for bucket, count := range histogram {
		fmt.Printf("- %s: %d students\n", bucket, count)
	}

	// Export analytics charts
	err = internal.ExportGPAHistogramChart(histogram, "gpa_histogram.png")
	if err != nil {
		fmt.Printf("Error exporting histogram chart: %v\n", err)
	} else {
		fmt.Println("✅ Generated GPA histogram chart: gpa_histogram.png")
	}

	err = internal.ExportDeanListChart("courseResults.json", "students.json", "deans_list.png")
	if err != nil {
		fmt.Printf("Error generating Dean's List chart: %v\n", err)
	} else {
		fmt.Println("✅ Generated Dean's List chart: deans_list.png")
	}

	err = internal.ExportAtRiskChart("courseResults.json", "students.json", "at_risk.png")
	if err != nil {
		fmt.Printf("Error generating At-Risk chart: %v\n", err)
	} else {
		fmt.Println("✅ Generated At-Risk students chart: at_risk.png")
	}

	// Load and process placement offers from JSON
	offers, err := internal.LoadOffers("placementOffers.json")
	if err != nil {
		fmt.Printf("Error loading placement offers: %v\n", err)
		return
	}

	categorizedOffers := internal.CategorizeOffers(offers)
	fmt.Println("\nPlacement Offers by Category (from imported data):")
	for category, categoryOffers := range categorizedOffers {
		fmt.Printf("- %s: %d offers\n", category, len(categoryOffers))
	}

	// Export placement analytics
	err = internal.ExportCategorizedOffers("categorized_offers.json", categorizedOffers)
	if err != nil {
		fmt.Printf("Error exporting categorized offers: %v\n", err)
	} else {
		fmt.Println("✅ Exported categorized offers: categorized_offers.json")
	}
}

func demonstrateFileOperations() {
	fmt.Println("\n6. FILE OPERATIONS - Export and Data Management")
	fmt.Println(strings.Repeat("=", 60))

	// File operations demo with sample data (not imported data)
	sampleStudents := []internal.Student{
		internal.NewStudent(1, "Alice Johnson"),
		internal.NewStudent(2, "Bob Smith"),
		internal.NewStudent(3, "Charlie Brown"),
	}

	// Student management operations
	err := internal.UpdateStudentName(sampleStudents, 1, "Alice Cooper")
	if err != nil {
		fmt.Printf("Error updating student name: %v\n", err)
	} else {
		fmt.Println("✅ Successfully updated student name")
	}

	foundStudent := internal.FindStudentByID(sampleStudents, 2)
	if foundStudent != nil {
		fmt.Printf("✅ Found student: %s\n", foundStudent.Name())
	}

	studentsWithName := internal.FindStudentsByName(sampleStudents, "Bob Smith")
	fmt.Printf("✅ Students named 'Bob Smith': %d\n", len(studentsWithName))

	// Export operations
	err = internal.SerializeStudents("exported_students.json", sampleStudents)
	if err != nil {
		fmt.Printf("Error serializing students: %v\n", err)
	} else {
		fmt.Println("✅ Successfully exported students to JSON")
	}

	// Sample academic records for export demo
	academicRecords := []internal.AcademicRecord{
		{StudentId: 1, CGPA: 8.5, Status: "Dean's List"},
		{StudentId: 2, CGPA: 6.2, Status: "Normal"},
		{StudentId: 3, CGPA: 4.8, Status: "At Risk"},
	}

	err = infrastructure.ExportDeanListStudents("deans_list_export.csv", academicRecords)
	if err != nil {
		fmt.Printf("Error exporting dean's list: %v\n", err)
	} else {
		fmt.Println("✅ Successfully exported dean's list to CSV")
	}

	err = infrastructure.ExportAtRiskStudents("at_risk_export.csv", academicRecords)
	if err != nil {
		fmt.Printf("Error exporting at-risk students: %v\n", err)
	} else {
		fmt.Println("✅ Successfully exported at-risk students to CSV")
	}

	// Sample results for export demo
	results := []internal.StudentResult{
		{CourseID: 101, CourseName: "Data Structures", StudentID: 1, StudentName: "Alice", Score: 85.5, Grade: "A"},
		{CourseID: 101, CourseName: "Data Structures", StudentID: 2, StudentName: "Bob", Score: 78.0, Grade: "B+"},
	}

	jsonData, err := infrastructure.ExportResultsAsJSON(results)
	if err != nil {
		fmt.Printf("Error exporting results as JSON: %v\n", err)
	} else {
		fmt.Printf("✅ JSON export completed (%d bytes)\n", len(jsonData))
	}

	csvData, err := infrastructure.ExportResultsAsCSV(results)
	if err != nil {
		fmt.Printf("Error exporting results as CSV: %v\n", err)
	} else {
		fmt.Printf("✅ CSV export completed (%d bytes)\n", len(csvData))
	}

	fmt.Println("\n✅ All file operations completed!")
}

