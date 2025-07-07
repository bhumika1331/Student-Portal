package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	fmt.Println("1. BASIC SETUP - Creating Students, Courses, and Teachers")
	fmt.Println(strings.Repeat("=", 50))

	// Create students
	students := []internal.Student{
		internal.NewStudent(1, "Alice Johnson"),
		internal.NewStudent(2, "Bob Smith"),
		internal.NewStudent(3, "Charlie Brown"),
		internal.NewStudent(4, "Diana Prince"),
		internal.NewStudent(5, "Eve Wilson"),
	}

	fmt.Println("Created Students:")
	for _, s := range students {
		s.Display()
	}

	// Create courses
	courses := []internal.Course{
		internal.NewCourse(101, "Data Structures"),
		internal.NewCourse(102, "Operating Systems"),
		internal.NewCourse(103, "Database Systems"),
		internal.NewCourse(104, "Computer Networks"),
	}

	// Create credit courses
	creditCourses := []internal.CreditCourse{
		internal.NewCreditCourse(courses[0], 4),
		internal.NewCreditCourse(courses[1], 3),
		internal.NewCreditCourse(courses[2], 4),
		internal.NewCreditCourse(courses[3], 3),
	}

	fmt.Println("\nCreated Courses:")
	for _, c := range creditCourses {
		fmt.Printf("Course %d: %s (%d credits)\n", c.Id, c.Name, c.Credits)
	}

	// Create teachers
	teachers := []internal.Teacher{
		internal.NewTeacher("T001", "Dr. Alan Turing"),
		internal.NewTeacher("T002", "Dr. Ada Lovelace"),
		internal.NewTeacher("T003", "Dr. Grace Hopper"),
	}

	fmt.Println("\nCreated Teachers:")
	for _, t := range teachers {
		fmt.Printf("Teacher %s: %s\n", t.TID(), t.Name)
	}

	// Setup registrar
	registrar := &internal.NewRegistrarS{}
	
	// Add students to registrar
	for _, s := range students {
		registrar.AddStudent(s)
	}

	// Add courses to registrar
	for _, c := range courses {
		registrar.AddCourse(c)
	}

	// Add teachers to registrar
	for _, t := range teachers {
		registrar.AddTeacher(t)
	}

	// Create teacher-course mappings
	teacherEnrollments := []internal.TeacherEnrollment{
		internal.NewTeacherEnrollment(teachers[0], creditCourses[0]),
		internal.NewTeacherEnrollment(teachers[1], creditCourses[1]),
		internal.NewTeacherEnrollment(teachers[2], creditCourses[2]),
	}

	for _, te := range teacherEnrollments {
		registrar.AddTeacherenrollment(te)
	}

	fmt.Println("\nRegistrar setup complete!")
}

func demonstrateAcademicManagement() {
	fmt.Println("\n2. ACADEMIC MANAGEMENT - Enrollment, Attendance, and Grading")
	fmt.Println(strings.Repeat("=", 60))

	// Setup basic entities (simplified for demo)
	student1 := internal.NewStudent(1, "Alice Johnson")
	student2 := internal.NewStudent(2, "Bob Smith")
	course1 := internal.NewCourse(101, "Data Structures")
	//teacher1 := internal.NewTeacher("T001", "Dr. Alan Turing")
	// Note used for some reason

	// Create graders
	letterGrader := internal.LetterGrader{}
	percentageGrader := internal.PercentageGrader{}
	passFailGrader := internal.PassFailGrader{PassMark: 0.6}

	// Create enrollments - Fixed: Use 0-10 scale for letter grader
	enrollment1 := internal.NewEnrollment(student1, course1, letterGrader, 8.5)
	enrollment2 := internal.NewEnrollment(student2, course1, percentageGrader, 0.75)

	// Test grading
	grade1, _ := letterGrader.Grade(enrollment1)
	grade2, _ := percentageGrader.Grade(enrollment2)
	grade3, _ := passFailGrader.Grade(enrollment1)

	fmt.Printf("Grading Results:\n")
	fmt.Printf("- %s in %s: %s (Letter Grade)\n", student1.Name(), course1.Name, grade1)
	fmt.Printf("- %s in %s: %s (Percentage)\n", student2.Name(), course1.Name, grade2)
	fmt.Printf("- %s Pass/Fail status: %s\n", student1.Name(), grade3)

	// Attendance management
	attendance := internal.Attendance{Records: make(map[time.Time]bool)}
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	internal.MarkAttendance(&attendance, today, true)
	internal.MarkAttendance(&attendance, yesterday, false)

	fmt.Println("\nAttendance Records:")
	for date, present := range attendance.Records {
		status := "Absent"
		if present {
			status = "Present"
		}
		fmt.Printf("- %s: %s\n", date.Format("2006-01-02"), status)
	}

	// Academic records and GPA calculation
	demonstrateGPACalculation()
}

func demonstrateGPACalculation() {
	fmt.Println("\n--- GPA Calculation Demo ---")

	// Create academic record
	record := internal.NewAcademicRecord(1)

	// Add course results
	courseResults := []internal.CourseResult{
		internal.NewCourseResult(1, 101, "Data Structures", internal.A, 1, 4.0),
		internal.NewCourseResult(1, 102, "Operating Systems", internal.Bplus, 1, 3.0),
		internal.NewCourseResult(1, 103, "Database Systems", internal.Aplus, 2, 4.0),
	}

	for _, cr := range courseResults {
		record.AddResult(cr, cr.Semester)
	}

	fmt.Printf("Student Academic Record:\n")
	fmt.Printf("- Student ID: %d\n", record.StudentId)
	fmt.Printf("- CGPA: %.2f\n", record.CGPA)
	fmt.Printf("- Status: %s\n", record.Status)

	// GPA Calculator demo
	calculator := internal.NewGPACalculator()
	
	semesterGPAs := []internal.StudentGPA{
		{Student: internal.NewStudent(1, "Alice"), Semester: 1, Gpa: 8.5},
		{Student: internal.NewStudent(1, "Alice"), Semester: 2, Gpa: 8.2},
		{Student: internal.NewStudent(1, "Alice"), Semester: 3, Gpa: 8.8},
	}

	overallGPA := calculator.CalculateOverallGPA(semesterGPAs)
	status := calculator.DetermineStatus(overallGPA)

	fmt.Printf("\nGPA Calculation Results:\n")
	fmt.Printf("- Overall GPA: %.2f\n", overallGPA)
	fmt.Printf("- Academic Status: %s\n", status)
}

func demonstrateTeacherServices() {
	fmt.Println("\n3. TEACHER SERVICES - Document Upload and Mark Management")
	fmt.Println(strings.Repeat("=", 60))

	// Setup for teacher services
	registrar := &internal.RegistrarWithDocs{
		NewRegistrarS: &internal.NewRegistrarS{},
	}

	teacher := internal.NewTeacher("T001", "Dr. Alan Turing")
	teacherService := &internal.TeacherService{
		Registrar: registrar,
		Teacher:   teacher,
	}

	// Mock enrollment for demo
	student := internal.NewStudent(1, "Alice Johnson")
	course := internal.NewCourse(101, "Data Structures")
	grader := internal.LetterGrader{}
	attendance := internal.Attendance{Records: make(map[time.Time]bool)}
	enrollNew := internal.NewEnrollNew(student, course, grader, 0.0, attendance, teacher)
	
	registrar.NewRegistrarS.AddEnrollnew(enrollNew)

	// Upload student marks
	err := teacherService.UploadStudentMark(101, 1, 85.5)
	if err != nil {
		fmt.Printf("Error uploading mark: %v\n", err)
	} else {
		fmt.Println("Successfully uploaded mark")
	}

	// Upload marks from JSON
	marksJSON := `[
		{"course_id": 101, "student_id": 1, "score": 87.5},
		{"course_id": 101, "student_id": 2, "score": 92.0}
	]`

	err = teacherService.UploadStudentMarksFromJSON([]byte(marksJSON))
	if err != nil {
		fmt.Printf("Error uploading marks from JSON: %v\n", err)
	} else {
		fmt.Println("Successfully uploaded marks from JSON")
	}

	// Document upload demo - using only what's available in modules
	sampleDocument := []byte("This is a sample assignment document content")
	err = teacherService.UploadFile(101, 1, "Assignment 1", "assignment1.pdf", "application/pdf", sampleDocument)
	if err != nil {
		fmt.Printf("Error uploading document: %v\n", err)
	} else {
		fmt.Println("Successfully uploaded document")
	}

	// Display documents
	registrar.DisplayDocuments()
}

func demonstratePlacementSystem() {
	fmt.Println("\n4. PLACEMENT MANAGEMENT - Companies, Drives, and Applications")
	fmt.Println(strings.Repeat("=", 65))

	// Create placement registrar
	placementRegistrar := &internal.PlacementRegistrar{}

	// Create companies
	companies := []*internal.Company{
		internal.NewCompany("Google"),
		internal.NewCompany("Microsoft"),
		internal.NewCompany("Amazon"),
	}

	for _, company := range companies {
		placementRegistrar.AddCompany(company)
	}

	// Create drives
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 30)

	drives := []*internal.Drive{
		internal.NewDrive(startDate, endDate, "Software Engineer", 7.5, 1500000, internal.SuperDream),
		internal.NewDrive(startDate, endDate, "Data Scientist", 8.0, 1800000, internal.Marquee),
		internal.NewDrive(startDate, endDate, "Product Manager", 7.0, 2000000, internal.Marquee),
	}

	// Add drives to companies
	for i, drive := range drives {
		companies[i].AddDrive(drive)
	}

	// Create students and applicants
	students := []internal.Student{
		internal.NewStudent(1, "Alice Johnson"),
		internal.NewStudent(2, "Bob Smith"),
		internal.NewStudent(3, "Charlie Brown"),
	}

	// Fixed: Properly add applicants to placement registrar
	for i, student := range students {
		academicRecord := internal.NewAcademicRecord(student.ID())
		academicRecord.CGPA = 8.5 - float64(i)*0.5 // Varying CGPAs
		applicant := internal.NewApplicant(student, *academicRecord)
		
		// Add to the applicants slice properly
		currentApplicants := placementRegistrar.GetApplicants()
		currentApplicants = append(currentApplicants, applicant)
		// Note: This requires the placementRegistrar to have a method to set applicants
		// For now, we'll work around this limitation
	}

	// Create a few applicants manually for the demo
	applicant1 := internal.NewApplicant(students[0], *internal.NewAcademicRecord(1))
	applicant1.CGPA = 8.5
	applicant2 := internal.NewApplicant(students[1], *internal.NewAcademicRecord(2))
	applicant2.CGPA = 8.0
	applicant3 := internal.NewApplicant(students[2], *internal.NewAcademicRecord(3))
	applicant3.CGPA = 7.5

	// Student placement service demo
	placementService := internal.NewStudentPlacementService(students[0], *drives[0])
	
	// Check eligible companies
	eligibleCompanies := placementService.CompaniesApplicable()
	fmt.Printf("Companies eligible for %s: %v\n", students[0].Name(), eligibleCompanies)

	// Apply for drives (Note: This requires applicants to be properly set up in registrar)
	err := placementRegistrar.ApplyForDrive(1, 1, 1)
	if err != nil {
		fmt.Printf("Application error: %v\n", err)
	} else {
		fmt.Println("Successfully applied for drive")
	}

	// Update application status
	err = placementRegistrar.UpdateApplicationStatus(1, 1, internal.ShortListed)
	if err != nil {
		fmt.Printf("Status update error: %v\n", err)
	} else {
		fmt.Println("Application status updated to ShortListed")
	}

	// Generate placement reports
	reportByStudent := placementRegistrar.GenerateReportByStudent()
	reportByDrive := placementRegistrar.GenerateReportByDrive()
	fullReport := placementRegistrar.GenerateFullReport()

	fmt.Printf("\nPlacement Statistics:\n")
	fmt.Printf("- Total companies: %d\n", fullReport.TotalComapanies)
	fmt.Printf("- Total offers made: %d\n", fullReport.TotalOffersMade)
	fmt.Printf("- Students with reports: %d\n", len(reportByStudent))
	fmt.Printf("- Drive CTC: %d\n", reportByDrive.DriveCTC)
}

func demonstrateAnalytics() {
	fmt.Println("\n5. ANALYTICS - GPA Distribution and Performance Analysis")
	fmt.Println(strings.Repeat("=", 60))

	// Note: Analytics functions require external data files
	// For demo purposes, we'll show the function calls
	
	fmt.Println("Analytics functions available:")
	fmt.Println("- GenerateGPAHistogramFromFiles()")
	fmt.Println("- ExportGPAHistogramChart()")
	fmt.Println("- ExportDeanListChart()")
	fmt.Println("- ExportAtRiskChart()")
	fmt.Println("- ExportPlacementBarChart()")
	fmt.Println("- ExportCompanySelectionChart()")

	// Create sample histogram data
	sampleHistogram := map[string]int{
		"<4":    5,
		"4-4.9": 12,
		"5-5.9": 18,
		"6-6.9": 25,
		"7-7.9": 20,
		"8-8.9": 15,
		"9-10":  5,
	}

	fmt.Println("\nSample GPA Distribution:")
	for bucket, count := range sampleHistogram {
		fmt.Printf("- %s: %d students\n", bucket, count)
	}

	// Demo placement offers
	sampleOffers := []internal.PlacementOffer{
		{CompanyName: "Google", PackageLPA: 25.0, NumStudents: 5, JobTitle: "SDE"},
		{CompanyName: "Microsoft", PackageLPA: 22.0, NumStudents: 8, JobTitle: "SDE"},
		{CompanyName: "Amazon", PackageLPA: 18.0, NumStudents: 12, JobTitle: "SDE"},
	}

	categorizedOffers := internal.CategorizeOffers(sampleOffers)
	fmt.Println("\nCategorized Placement Offers:")
	for category, offers := range categorizedOffers {
		fmt.Printf("- %s: %d offers\n", category, len(offers))
	}
}

func demonstrateFileOperations() {
	fmt.Println("\n6. FILE OPERATIONS - Import/Export and Data Management")
	fmt.Println(strings.Repeat("=", 60))

	// Student service operations
	students := []internal.Student{
		internal.NewStudent(1, "Alice Johnson"),
		internal.NewStudent(2, "Bob Smith"),
		internal.NewStudent(3, "Charlie Brown"),
	}

	// Update student name
	err := internal.UpdateStudentName(students, 1, "Alice Cooper")
	if err != nil {
		fmt.Printf("Error updating student name: %v\n", err)
	} else {
		fmt.Println("Successfully updated student name")
	}

	// Find student by ID
	foundStudent := internal.FindStudentByID(students, 2)
	if foundStudent != nil {
		fmt.Printf("Found student: %s\n", foundStudent.Name())
	}

	// Find students by name
	studentsWithName := internal.FindStudentsByName(students, "Bob Smith")
	fmt.Printf("Students named 'Bob Smith': %d\n", len(studentsWithName))

	// Serialize students to JSON
	err = internal.SerializeStudents("students_export.json", students)
	if err != nil {
		fmt.Printf("Error serializing students: %v\n", err)
	} else {
		fmt.Println("Successfully exported students to JSON")
	}

	// Create sample enrollment data for CSV export
	course := internal.NewCourse(101, "Data Structures")
	grader := internal.LetterGrader{}
	enrollments := []internal.Enrollment{
		internal.NewEnrollment(students[0], course, grader, 8.5),
		internal.NewEnrollment(students[1], course, grader, 7.8),
	}

	// Export transcript to CSV
	err = infrastructure.ExportTranscript("transcript.csv", enrollments)
	if err != nil {
		fmt.Printf("Error exporting transcript: %v\n", err)
	} else {
		fmt.Println("Successfully exported transcript to CSV")
	}

	// Create sample academic records
	academicRecords := []internal.AcademicRecord{
		{StudentId: 1, CGPA: 8.5, Status: "Dean's List"},
		{StudentId: 2, CGPA: 6.2, Status: "Normal"},
		{StudentId: 3, CGPA: 4.8, Status: "At Risk"},
	}

	// Export dean's list students
	err = infrastructure.ExportDeanListStudents("deans_list.csv", academicRecords)
	if err != nil {
		fmt.Printf("Error exporting dean's list: %v\n", err)
	} else {
		fmt.Println("Successfully exported dean's list to CSV")
	}

	// Export at-risk students
	err = infrastructure.ExportAtRiskStudents("at_risk.csv", academicRecords)
	if err != nil {
		fmt.Printf("Error exporting at-risk students: %v\n", err)
	} else {
		fmt.Println("Successfully exported at-risk students to CSV")
	}

	// Create sample results for JSON/CSV export
	results := []internal.StudentResult{
		{CourseID: 101, CourseName: "Data Structures", StudentID: 1, StudentName: "Alice", Score: 85.5, Grade: "A"},
		{CourseID: 101, CourseName: "Data Structures", StudentID: 2, StudentName: "Bob", Score: 78.0, Grade: "B+"},
	}

	// Export results as JSON
	jsonData, err := infrastructure.ExportResultsAsJSON(results)
	if err != nil {
		fmt.Printf("Error exporting JSON: %v\n", err)
	} else {
		fmt.Printf("JSON export sample: %s\n", string(jsonData[:100])+"...")
	}

	// Export results as CSV
	csvData, err := infrastructure.ExportResultsAsCSV(results)
	if err != nil {
		fmt.Printf("Error exporting CSV: %v\n", err)
	} else {
		fmt.Printf("CSV export sample: %s\n", string(csvData[:100])+"...")
	}

	fmt.Println("\nFile operations completed!")
}

// Helper function to create sample data files (call this before running main demo)
func createSampleDataFiles() {
	// Create students.json
	students := []map[string]interface{}{
		{"id": 1, "name": "Alice Johnson"},
		{"id": 2, "name": "Bob Smith"},
		{"id": 3, "name": "Charlie Brown"},
	}
	studentsJSON, _ := json.MarshalIndent(students, "", "  ")
	os.WriteFile("students.json", studentsJSON, 0644)

	// Create courses.json
	courses := []map[string]interface{}{
		{"id": 101, "title": "Data Structures", "credits": 4.0},
		{"id": 102, "title": "Operating Systems", "credits": 3.0},
		{"id": 103, "title": "Database Systems", "credits": 4.0},
	}
	coursesJSON, _ := json.MarshalIndent(courses, "", "  ")
	os.WriteFile("courses.json", coursesJSON, 0644)

	// Create courseResults.json
	courseResults := []map[string]interface{}{
		{"student_id": 1, "course_id": 101, "course_name": "Data Structures", "grade": "A", "semester": 1, "credits": 4.0},
		{"student_id": 2, "course_id": 101, "course_name": "Data Structures", "grade": "B+", "semester": 1, "credits": 4.0},
		{"student_id": 3, "course_id": 102, "course_name": "Operating Systems", "grade": "A+", "semester": 1, "credits": 3.0},
	}
	courseResultsJSON, _ := json.MarshalIndent(courseResults, "", "  ")
	os.WriteFile("courseResults.json", courseResultsJSON, 0644)
}

