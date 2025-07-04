package internal

import (
	"log"
	"os"
	"testing"
	"time"
)

// --- Sample Student Constructor for Tests ---
func createSampleStudents() []Student {
	return []Student{
		NewStudent(1, "Alice"),
		NewStudent(2, "Bob"),
		NewStudent(3, "Charlie"),
	}
}

// --- Test NewStudent ---
func TestNewStudent_Valid(t *testing.T) {
	s := NewStudent(10, "Test")
	if s.ID() != 10 || s.Name() != "Test" {
		t.Errorf("expected ID=10, Name='Test'; got ID=%d, Name=%s", s.ID(), s.Name())
	}
}

func TestNewStudent_Invalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for id <= 0, but no panic occurred")
		}
	}()
	NewStudent(0, "Invalid")
}

// --- Test Display ---
func TestDisplay(t *testing.T) {
	s := NewStudent(42, "TestName")
	// not asserting anything here since Display only prints to stdout
	s.Display()
}

// --- Test UpdateStudentName ---
func TestUpdateStudentName(t *testing.T) {
	students := createSampleStudents()
	err := UpdateStudentName(students, 2, "Bobby")
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if students[1].Name() != "Bobby" {
		t.Errorf("expected name to be 'Bobby', got: %s", students[1].Name())
	}

	err = UpdateStudentName(students, 999, "Ghost")
	if err == nil {
		t.Error("expected error for non-existent student ID")
	}
}

// --- Test FindStudentByID ---
func TestFindStudentByID(t *testing.T) {
	students := createSampleStudents()

	s := FindStudentByID(students, 3)
	if s == nil || s.Name() != "Charlie" {
		t.Error("expected to find Charlie")
	}

	s = FindStudentByID(students, 999)
	if s != nil {
		t.Error("expected nil for non-existent student")
	}
}

// --- Test FindStudentsByName ---
func TestFindStudentsByName(t *testing.T) {
	students := createSampleStudents()
	students = append(students, NewStudent(4, "Alice")) // duplicate name

	found := FindStudentsByName(students, "Alice")
	if len(found) != 2 {
		t.Errorf("expected 2 students named Alice, got: %d", len(found))
	}

	found = FindStudentsByName(students, "Unknown")
	if len(found) != 0 {
		t.Errorf("expected 0 results for unknown name, got: %d", len(found))
	}
}

// --- Test DeleteStudentByID ---
func TestDeleteStudentByID(t *testing.T) {
	students := createSampleStudents()

	newList, err := DeleteStudentByID(students, 2)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if len(newList) != len(students)-1 {
		t.Errorf("expected length %d, got %d", len(students)-1, len(newList))
	}
	for _, s := range newList {
		if s.ID() == 2 {
			t.Error("student with ID 2 was not deleted")
		}
	}

	newList, err = DeleteStudentByID(newList, 999)
	if err == nil {
		t.Error("expected error when deleting non-existent student")
	}
}

// --- Test SerializeStudents ---
func TestSerializeStudents(t *testing.T) {
	students := createSampleStudents()

	err := SerializeStudents("students.json", students) // No return value, just ensure no panic
	if err != nil {
		t.Fatalf("Could not create the json: %v", err)
	}
}

// --- Test DeserializeStudents ---
// func TestDeserializeStudents(t *testing.T) {
// 	jsonData := `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`

// 	students, err := DeserializeStudents([]byte(jsonData))
// 	if err != nil {
// 		t.Errorf("deserialization failed: %v", err)
// 	}
// 	if len(students) != 2 || students[0].Name() != "Alice" {
// 		t.Error("unexpected deserialized student data")
// 	}

// 	badData := []byte(`{invalid json}`)
// 	_, err = DeserializeStudents(badData)
// 	if err == nil {
// 		t.Error("expected error for invalid JSON")
// 	}

// 	_, err = DeserializeStudents([]byte(""))
// 	if err == nil {
// 		t.Error("expected error for empty data")
// 	}
// }

// --- Test StudentPlacementService ---
func TestStudentPlacementService(t *testing.T) {
	// Create a sample student and applicant using constructors
	student := NewStudent(1, "Alice")
	ar := AcademicRecord{StudentId: 1, CGPA: 9.0}
	applicant := Applicant{Student: student, AcademicRecord: ar}

	// Create a drive and company using constructors
	drive := NewDrive(
		time.Now(), time.Now().Add(24*time.Hour), "Engineer", 8.0, 1000000, Day,
	)
	company := NewCompany("TestCorp")
	company.drives = []*Drive{drive}

	// Create a placement registrar with the company and applicant
	pr := PlacementRegistrar{
		companies:    []*Company{company},
		applications: []*Application{},
		applicants:   []*Applicant{&applicant},
	}

	// Create the service
	service := &StudentPlacementService{
		student:            student,
		drive:              *drive,
		applicant:          applicant,
		PlacementRegistrar: pr,
		Company:            *company,
	}

	// Test GetDrive
	if service.GetDrive().ID() != drive.ID() {
		t.Errorf("expected drive ID %d, got %d", drive.ID(), service.GetDrive().ID())
	}

	// Test CompaniesApplicable
	applicable := service.CompaniesApplicable()
	if len(applicable) == 0 || applicable[0] != "TestCorp" {
		t.Errorf("expected TestCorp to be applicable, got %v", applicable)
	}

	// Test CompaniesApplied (should be empty initially)
	applied := service.CompaniesApplied()
	if appliedSlice, ok := applied.([]string); ok {
		if len(appliedSlice) != 0 {
			t.Errorf("expected no companies applied, got %v", appliedSlice)
		}
	} else {
		t.Errorf("expected CompaniesApplied to return []string, got %T", applied)
	}

	// Test Apply (simulate application)
	service.PlacementRegistrar.applicants = []*Applicant{&service.applicant}
	service.PlacementRegistrar.companies = []*Company{&service.Company}
	service.PlacementRegistrar.applications = nil
	service.drive = *drive
	service.Company = *company

	// The Apply method should not panic and should return an error or nil
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Apply panicked: %v", r)
		}
	}()
	_ = service.Apply()
}

// --- Test Notifications ---
func TestDriveNotification(t *testing.T) {
	drive := Drive{id: 1, roleName: "Engineer"}
	n := NewDriveNotification(drive)
	returned := n.Send().(Drive)
	if returned.id != drive.id || returned.roleName != drive.roleName {
		t.Errorf("DriveNotification did not return correct drive")
	}
}

func TestResultNotification(t *testing.T) {
	n := NewResultNotification(5)
	if n.Send() != 5 {
		t.Errorf("ResultNotification did not return correct days left")
	}
}

// --- Additional Edge/Negative Tests for student_service.go ---

func TestUpdateStudentName_EmptySlice(t *testing.T) {
	err := UpdateStudentName([]Student{}, 1, "Test")
	if err == nil {
		t.Error("expected error for empty slice, got nil")
	}
}

func TestFindStudentByID_EmptySlice(t *testing.T) {
	res := FindStudentByID([]Student{}, 1)
	if res != nil {
		t.Error("expected nil for empty slice")
	}
}

func TestFindStudentsByName_EmptySlice(t *testing.T) {
	res := FindStudentsByName([]Student{}, "Test")
	if len(res) != 0 {
		t.Errorf("expected 0, got %d", len(res))
	}
}

func TestFindStudentsByName_CaseSensitivity(t *testing.T) {
	students := []Student{NewStudent(1, "Alice")}
	res := FindStudentsByName(students, "alice")
	if len(res) != 0 {
		t.Error("expected 0 for case-sensitive mismatch")
	}
}

func TestDeleteStudentByID_EmptySlice(t *testing.T) {
	newList, err := DeleteStudentByID([]Student{}, 1)
	if err == nil {
		t.Error("expected error for empty slice")
	}
	if len(newList) != 0 {
		t.Errorf("expected 0, got %d", len(newList))
	}
}

func TestDeleteStudentByID_DeleteOnlyStudent(t *testing.T) {
	students := []Student{NewStudent(1, "Solo")}
	newList, err := DeleteStudentByID(students, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(newList) != 0 {
		t.Errorf("expected 0, got %d", len(newList))
	}
}

// func TestSerializeStudents_InvalidFilename(t *testing.T) {
// 	students := createSampleStudents()
// 	err := SerializeStudents("/invalid_path/\x00students.json", students)
// 	if err == nil {
// 		t.Error("expected error for invalid filename")
// 	}
// }

func TestDeserializeStudents_EmptyData(t *testing.T) {
	students, err := DeserializeStudents([]byte(""))
	if err == nil {
		t.Error("expected error for empty data")
	}
	if students != nil && len(students) != 0 {
		t.Errorf("expected empty slice, got %v", students)
	}
}

func TestSerializeStudents_Cleanup(t *testing.T) {
	students := createSampleStudents()
	filename := "test_students_cleanup.json"
	err := SerializeStudents(filename, students)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	// Clean up
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("failed to clean up test file: %v", err)
	}
}
