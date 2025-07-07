package main

import (
	"fmt"
	"oops/main/infrastructure"
	"oops/main/internal"
	"time"
)

var courseResults []internal.CourseResult

func main() {
	// =============================
	//   STUDENT PORTAL DEMO (REAL LOGIC)
	// =============================
	fmt.Println("\n===== STUDENT PORTAL DEMO (REAL LOGIC) =====")

	// 1. View grades and academic transcripts (already real)
	fmt.Println("\n-- View Grades and Academic Transcript --")
	student := internal.NewStudent(101, "Alice")
	course := internal.NewCourse(201, "Algorithms")
	grader := internal.PercentageGrader{}
	enrollment := internal.NewEnrollment(student, course, grader, 88.5)
	grade, _ := grader.Grade(enrollment)
	fmt.Printf("Student: %s, Course: %s, Grade: %s\n", student.Name(), course.Name, grade)
	enrollments := []internal.Enrollment{enrollment}
	err := infrastructure.ExportTranscript("student_transcript.csv", enrollments)
	if err == nil {
		fmt.Println("Transcript exported to student_transcript.csv")
	}

	// 2. Track attendance (real logic: mark and fetch)
	fmt.Println("\n-- Track Attendance --")
	att := internal.Attendance{}
	internal.MarkAttendance(&att, time.Now(), true)
	fmt.Printf("Attendance marked for %s. Records: %v\n", student.Name(), att.Records)

	// 3. Placement eligibility, shortlist status, and offer details (real logic)
	fmt.Println("\n-- Placement Eligibility, Shortlist Status, and Offers --")
	company := internal.NewCompany("DreamTech")
	drive := internal.NewDrive(time.Now(), time.Now().AddDate(0,0,7), "SDE", 7.0, 1200000, internal.Dream)
	company.AddDrive(drive)
	appRecord := internal.NewAcademicRecord(student.ID())
	applicant := internal.NewApplicant(student, *appRecord)
	eligible := drive.Eligibility().CheckEligibility(applicant)
	fmt.Printf("Eligible for drive '%s': %v\n", drive.RoleName(), eligible)
	// Simulate offer: add applicant to PlacementRegistrar, check offers
	placementRegistrar := internal.PlacementRegistrar{}
	placementRegistrar.AddCompany(company)
	placementRegistrar.applicants = append(placementRegistrar.applicants, applicant)
	fmt.Printf("%s registered for drive '%s' at %s\n", student.Name(), drive.RoleName(), company.Name())
	// Check offers (real logic: getFinalOffer)
	offerCTC, offerErr := applicant.getFinalOffer()
	if offerErr != nil {
		fmt.Println("Offer Details:", offerErr)
	} else {
		fmt.Printf("Offer Details: Final CTC: %d\n", offerCTC)
	}

	// 4. Register for placement drives
	fmt.Println("\n-- Register for Placement Drives --")
	placementRegistrar := internal.PlacementRegistrar{}
	placementRegistrar.AddCompany(company)
	fmt.Printf("%s registered for drive '%s' at %s\n", student.Name(), drive.RoleName(), company.Name())

	// --- STUDENT & TEACHER DEMO ---
	fmt.Println("=== Student & Teacher Demo ===")
	student1 := internal.NewStudent(1, "Alice")
	student2 := internal.NewStudent(2, "Bob")
	student1.Display()
	student2.Display()

	teacher1 := internal.NewTeacher("T1", "Dr. Smith")
	teacher2 := internal.NewTeacher("T2", "Prof. Johnson")
	fmt.Printf("Teacher: %s (%s)\n", teacher1.Name, teacher1.ID)
	fmt.Printf("Teacher: %s (%s)\n", teacher2.Name, teacher2.ID)

	// --- COURSE & ENROLLMENT DEMO ---
	fmt.Println("\n=== Course & Enrollment Demo ===")
	course1 := internal.NewCourse(101, "Mathematics")
	course2 := internal.NewCourse(102, "Physics")
	creditCourse1 := internal.NewCreditCourse(course1, 4)
	creditCourse2 := internal.NewCreditCourse(course2, 3)

	teacherEnrollments := []internal.TeacherEnrollment{
		internal.NewTeacherEnrollment(teacher1, creditCourse1),
		internal.NewTeacherEnrollment(teacher2, creditCourse2),
	}

	percentageGrader := internal.PercentageGrader{}
	enrollment1 := internal.NewEnrollment(student1, course1, percentageGrader, 85.0)
	enrollment2 := internal.NewEnrollment(student2, course2, percentageGrader, 92.0)

	attendance := internal.Attendance{}
	enrollNew1, ok1 := internal.Enroll(enrollment1, attendance, teacher1, teacherEnrollments)
	enrollNew2, ok2 := internal.Enroll(enrollment2, attendance, teacher2, teacherEnrollments)
	fmt.Printf("Enrollment 1 success: %v, %+v\n", ok1, enrollNew1)
	fmt.Printf("Enrollment 2 success: %v, %+v\n", ok2, enrollNew2)

	// --- GPA CALCULATION DEMO ---
	fmt.Println("\n=== GPA Calculation Demo ===")
	semesters := []internal.StudentGPA{
		{Student: student1, Semester: 1, Gpa: 8.5},
		{Student: student1, Semester: 2, Gpa: 9.0},
	}
	gpaCalc := internal.NewGPACalculator()
	overallGPA := gpaCalc.CalculateOverallGPA(semesters)
	status := gpaCalc.DetermineStatus(overallGPA)
	fmt.Printf("Student: %s, Overall GPA: %.2f, Status: %s\n", student1.Name(), overallGPA, status)

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

	// --- ATTENDANCE DEMO ---
	fmt.Println("\n=== Attendance Demo ===")
	att := internal.Attendance{}
	date := time.Now()
	internal.MarkAttendance(&att, date, true)
	fmt.Printf("Attendance on %v: %v\n", date.Format("2006-01-02"), att.Records[date])

	// --- PLACEMENT & DRIVES DEMO ---
	fmt.Println("\n=== Placement & Drives Demo ===")
	company := internal.NewCompany("TechCorp")
	drive := internal.NewDrive(time.Now(), time.Now().AddDate(0, 0, 7), "Backend Developer", 7.0, 1200000, internal.Dream)
	company.AddDrive(drive)
	placementRegistrar := internal.PlacementRegistrar{}
	placementRegistrar.AddCompany(company)
	fmt.Printf("Company: %s, Drives: %d\n", company.Name(), len(company.Drives()))

	// --- ACADEMIC RECORDS DEMO ---
	fmt.Println("\n=== Academic Records Demo ===")
	ar := internal.NewAcademicRecord(student1.ID())
	cr := internal.NewCourseResult(student1.ID(), course1.Id, course1.Name, internal.A, 1, 4)
	ar.AddResult(cr, 1)
	fmt.Printf("Academic Record for Student %d: CGPA: %.2f, Status: %s\n", ar.StudentId, ar.CGPA, ar.Status)

	// --- APPLICATIONS DEMO ---
	fmt.Println("\n=== Applications Demo ===")
	applicant := internal.NewApplicant(student1, *ar)
	applicant.AddDrivesAppliedFor(drive)
	fmt.Printf("Applicant %s applied for %d drives.\n", applicant.Name(), len(applicant.DrivesAppliedFor()))

	// --- INFRASTRUCTURE DEMO: Load Course Results ---
	fmt.Println("\n=== Infrastructure Demo: Load Course Results ===")
	courseResults := infrastructure.LoadCourseResults()
	fmt.Printf("Loaded %d course results from file.\n", len(courseResults))

	// --- INFRASTRUCTURE DEMO: Export Transcript ---
	fmt.Println("\n=== Infrastructure Demo: Export Transcript ===")
	enrollments := []internal.Enrollment{
		internal.NewEnrollment(student1, course1, percentageGrader, 85.0),
		internal.NewEnrollment(student2, course2, percentageGrader, 92.0),
	}
	err := infrastructure.ExportTranscript("transcript.csv", enrollments)
	if err != nil {
		fmt.Println("Failed to export transcript:", err)
	} else {
		fmt.Println("Transcript exported to transcript.csv")
	}

	// --- INFRASTRUCTURE DEMO: Export At-Risk and Dean's List Students ---
	fmt.Println("\n=== Infrastructure Demo: Export At-Risk and Dean's List Students ===")
	ar1 := *internal.NewAcademicRecord(student1.ID())
	ar1.Status = "At Risk"
	ar2 := *internal.NewAcademicRecord(student2.ID())
	ar2.Status = "Dean's List"
	records := []internal.AcademicRecord{ar1, ar2}
	err = infrastructure.ExportAtRiskStudents("at_risk_students.csv", records)
	if err != nil {
		fmt.Println("Failed to export at-risk students:", err)
	} else {
		fmt.Println("At-risk students exported to at_risk_students.csv")
	}
	err = infrastructure.ExportDeanListStudents("dean_list_students.csv", records)
	if err != nil {
		fmt.Println("Failed to export dean's list students:", err)
	} else {
		fmt.Println("Dean's list students exported to dean_list_students.csv")
	}

	// --- INFRASTRUCTURE DEMO: Export Results as JSON and CSV ---
	fmt.Println("\n=== Infrastructure Demo: Export Results as JSON and CSV ===")
	studentResults := []internal.StudentResult{
		{CourseID: course1.Id, CourseName: course1.Name, StudentID: student1.ID(), StudentName: student1.Name(), Score: 85.0, Grade: "A"},
		{CourseID: course2.Id, CourseName: course2.Name, StudentID: student2.ID(), StudentName: student2.Name(), Score: 92.0, Grade: "A+"},
	}
	jsonBytes, err := infrastructure.ExportResultsAsJSON(studentResults)
	if err != nil {
		fmt.Println("Failed to export results as JSON:", err)
	} else {
		fmt.Println("Results as JSON:\n", string(jsonBytes))
	}
	csvBytes, err := infrastructure.ExportResultsAsCSV(studentResults)
	if err != nil {
		fmt.Println("Failed to export results as CSV:", err)
	} else {
		fmt.Println("Results as CSV:\n", string(csvBytes))
	}

	// Run GPA Histogram Analysis
	var hist map[string]int
	hist, err = internal.GenerateGPAHistogramFromFiles("courseResults.json", "students.json")
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

	// =============================
	//   ANALYTICS ENGINE DEMO
	// =============================
	fmt.Println("\n===== ANALYTICS ENGINE DEMO =====")

	// 1. Generate GPA Trends
	fmt.Println("\n-- Generate GPA Trends --")
	hist, err = internal.GenerateGPAHistogramFromFiles("courseResults.json", "students.json")
	if err == nil {
		err = internal.ExportGPAHistogramChart(hist, "gpa_histogram.png")
		if err == nil {
			fmt.Println("GPA Trends Chart Generated: gpa_histogram.png")
		}
	}

	// 2. Generate Dean's List
	fmt.Println("\n-- Generate Dean's List --")
	err = internal.ExportDeanListChart("courseResults.json", "students.json", "dean_list.png")
	if err == nil {
		fmt.Println("Dean List Chart Generated: dean_list.png")
	}

	// 3. At-Risk Student Detection
	fmt.Println("\n-- At-Risk Student Detection --")
	err = internal.ExportAtRiskChart("courseResults.json", "students.json", "at_risk_students.png")
	if err == nil {
		fmt.Println("At-Risk Chart Generated: at_risk_students.png")
	}

	// 4. Placement Offer Analysis
	fmt.Println("\n-- Placement Offer Analysis --")
	offers, err = internal.LoadOffers("placement_offers.json")
	if err == nil {
		categorized := internal.CategorizeOffers(offers)
		err = internal.ExportCategorizedOffers("placement_chart.json", categorized)
		if err == nil {
			fmt.Println("Placement offers categorized and saved: placement_chart.json")
		}
	}
	err = internal.ExportPlacementBarChart("placement_chart.json", "placement_chart.png")
	if err == nil {
		fmt.Println("Placement Chart Generated: placement_chart.png")
	}

	// 5. Company-wise Selection Metrics
	fmt.Println("\n-- Company-wise Selection Metrics --")
	err = internal.ExportCompanySelectionChart("placement_offers.json", "company_selection.png", "company_selection.json")
	if err == nil {
		fmt.Println("Company Selection Chart Generated: company_selection.png")
	}

	// =============================
	//   TEACHER PORTAL DEMO (REAL LOGIC)
	// =============================
	fmt.Println("\n===== TEACHER PORTAL DEMO (REAL LOGIC) =====")

	// 1. Upload student marks using configurable grading strategies (real logic)
	fmt.Println("\n-- Upload Student Marks (Strategy Pattern) --")
	teacher := internal.NewTeacher("T100", "Dr. Smith")
	studentA := internal.NewStudent(201, "Bob")
	studentB := internal.NewStudent(202, "Carol")
	courseA := internal.NewCourse(301, "Data Structures")
	percentageGrader := internal.PercentageGrader{}
	passFailGrader := internal.PassFailGrader{}
	letterGrader := internal.LetterGrader{}
	// Setup RegistrarWithDocs and TeacherService for real logic
	regWithDocs := &internal.RegistrarWithDocs{NewRegistrarS: &internal.NewRegistrarS{}}
	regWithDocs.AddTeacher(teacher)
	creditCourse := internal.NewCreditCourse(courseA, 4)
	regWithDocs.AddTeacherenrollment(internal.NewTeacherEnrollment(teacher, creditCourse))
	att := internal.Attendance{}
	enrollNewA := internal.NewEnrollNew(studentA, courseA, percentageGrader, 78.0, att, teacher)
	enrollNewB := internal.NewEnrollNew(studentB, courseA, passFailGrader, 62.0, att, teacher)
	regWithDocs.AddEnrollnew(enrollNewA)
	regWithDocs.AddEnrollnew(enrollNewB)
	teacherService := internal.TeacherService{Registrar: regWithDocs, Teacher: teacher}
	// Upload marks
	uploadMarkErrA := teacherService.UploadStudentMark(courseA.Id, studentA.ID(), 78.0)
	uploadMarkErrB := teacherService.UploadStudentMark(courseA.Id, studentB.ID(), 62.0)
	if uploadMarkErrA != nil {
		fmt.Println("Upload mark failed for Bob:", uploadMarkErrA)
	}
	if uploadMarkErrB != nil {
		fmt.Println("Upload mark failed for Carol:", uploadMarkErrB)
	}
	// Change grading policy mid-semester
	regWithDocs.SetGrader(courseA.Id, letterGrader)
	fmt.Println("Changed grading policy for Data Structures to Letter Grader.")

	// 2. Upload assignments, lecture notes, or reading material (real upload already done above)
	fmt.Println("\n-- Upload Assignments and Materials (Real Upload) --")
	dummyContent := []byte("This is the content of Assignment2.docx")
	uploadErr := teacherService.UploadFile(courseA.Id, studentA.ID(), "Assignment 2", "Assignment2.docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document", dummyContent)
	if uploadErr != nil {
		fmt.Println("Upload failed:", uploadErr)
	} else {
		fmt.Println("Upload succeeded!")
		regWithDocs.DisplayDocuments()
	}

	// 3. Mark attendance by session or batch (real logic)
	fmt.Println("\n-- Mark Attendance by Session --")
	internal.MarkAttendance(&att, time.Now(), true)
	teacherService.DisplayAttendance(courseA.Id, studentA.ID())
	teacherService.DisplayAttendance(courseA.Id, studentB.ID())

	// 4. Export result summaries (CSV, JSON) (real logic)
	fmt.Println("\n-- Export Result Summaries --")
	results, resErr := teacherService.GetCourseResults(courseA.Id)
	if resErr != nil {
		fmt.Println("Failed to get course results:", resErr)
	} else {
		jsonBytes, err := infrastructure.ExportResultsAsJSON(results)
		if err == nil {
			fmt.Println("Results as JSON:\n", string(jsonBytes))
		}
		csvBytes, err := infrastructure.ExportResultsAsCSV(results)
		if err == nil {
			fmt.Println("Results as CSV:\n", string(csvBytes))
		}
	}

	// =============================
	//   ADMINISTRATION PORTAL DEMO (REAL LOGIC)
	// =============================
	fmt.Println("\n===== ADMINISTRATION PORTAL DEMO (REAL LOGIC) =====")

	// 1. Enroll new students and teachers (real logic)
	fmt.Println("\n-- Enroll New Students and Teachers --")
	adminRegistrar := internal.Registrar{}
	newStudent := internal.NewStudent(301, "David")
	newTeacher := internal.NewTeacher("T200", "Prof. Lee")
	adminRegistrar.AddStudent(newStudent)
	adminRegistrar.AddCourse(courseA)
	adminRegistrar.AddCourse(courseB)
	adminRegistrar.AddCourse(course)
	adminRegistrar.AddStudent(student)
	adminRegistrar.AddStudent(studentA)
	adminRegistrar.AddStudent(studentB)
	adminRegistrar.DisplayStudents()
	adminRegistrar.DisplayCourses()

	// 2. Create and update courses, map them to grading strategies (real logic)
	fmt.Println("\n-- Create/Update Courses and Map Grading Strategies --")
	adminRegistrar.SetGrader(courseA.Id, percentageGrader)
	adminRegistrar.SetGrader(courseB.Id, letterGrader)
	fmt.Println("Mapped grading strategies to courses.")

	// 3. Modify grading policies mid-semester via strategy pattern (real logic)
	fmt.Println("\n-- Modify Grading Policy Mid-Semester --")
	adminRegistrar.SetGrader(courseA.Id, letterGrader)
	fmt.Println("Changed grading policy for Data Structures to Letter Grader.")

	// 4. Generate consolidated academic reports (real logic)
	fmt.Println("\n-- Generate Consolidated Academic Reports --")
	academicRecord := internal.NewAcademicRecord(newStudent.ID())
	cr1 := internal.NewCourseResult(newStudent.ID(), courseA.Id, courseA.Name, internal.A, 1, 4)
	cr2 := internal.NewCourseResult(newStudent.ID(), courseB.Id, courseB.Name, internal.Bplus, 1, 3)
	academicRecord.AddResult(cr1, 1)
	academicRecord.AddResult(cr2, 1)
	fmt.Printf("Academic Record for %s: CGPA: %.2f, Status: %s\n", newStudent.Name(), academicRecord.CGPA, academicRecord.Status)

	// =============================
	//   PLACEMENT CELL PORTAL DEMO (REAL LOGIC)
	// =============================
	fmt.Println("\n===== PLACEMENT CELL PORTAL DEMO (REAL LOGIC) =====")

	// 1. Add/update company listings with eligibility criteria and drive details (real logic)
	fmt.Println("\n-- Add/Update Company Listings and Drives --")
	placementCell := internal.PlacementRegistrar{}
	company1 := internal.NewCompany("InnovateX")
	drive1 := internal.NewDrive(time.Now(), time.Now().AddDate(0, 0, 10), "ML Engineer", 8.0, 1800000, internal.SuperDream)
	company1.AddDrive(drive1)
	placementCell.AddCompany(company1)
	fmt.Printf("Added company: %s with drive: %s (CTC: %d)\n", company1.Name(), drive1.RoleName(), drive1.CTC())

	// 2. Manage job profiles, compensation packages, and application deadlines (real logic)
	fmt.Println("\n-- Manage Job Profiles, Compensation, Deadlines --")
	drive1.SetCTC(2000000)
	drive1.SetEndDate(time.Now().AddDate(0, 0, 14))
	fmt.Printf("Updated drive: %s, New CTC: %d, New Deadline: %v\n", drive1.RoleName(), drive1.CTC(), drive1.EndDate())

	// 3. Track student registrations, shortlists, and final selections (real logic)
	fmt.Println("\n-- Track Student Registrations, Shortlists, Selections --")
	placementApplicant := internal.NewApplicant(newStudent, *internal.NewAcademicRecord(newStudent.ID()))
	placementApplicant.AddDrivesAppliedFor(drive1)
	placementCell.applicants = append(placementCell.applicants, placementApplicant)
	fmt.Printf("%s registered for drive: %s\n", newStudent.Name(), drive1.RoleName())
	// Shortlist and selection logic would require status update; here we just show registration and can check offers
	offerCTC, offerErr := placementApplicant.getFinalOffer()
	if offerErr != nil {
		fmt.Println("Final Selection/Offer:", offerErr)
	} else {
		fmt.Printf("Final Selection/Offer: CTC: %d\n", offerCTC)
	}

	// 4. Generate placement statistics, offer reports, recruiter engagement summaries (real logic)
	fmt.Println("\n-- Generate Placement Statistics and Reports --")
	reportByDrive := placementCell.GenerateReportByDrive()
	fmt.Printf("Report By Drive: Company: %s, Drive: %s, CTC: %d, Selected: %d\n", reportByDrive.company.Name(), reportByDrive.drive.RoleName(), reportByDrive.driveCTC, reportByDrive.noOfSelectedStudents)
	fullReport := placementCell.GenerateFullReport()
	fmt.Printf("Total Companies: %d, Total Offers: %d\n", fullReport.totalComapanies, fullReport.totalOffersMade)

}
