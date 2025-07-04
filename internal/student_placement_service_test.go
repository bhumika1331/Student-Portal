package internal

import (
	"testing"
	"time"
)

func createTestPlacementService() *StudentPlacementService {
	student := NewStudent(1, "Alice")
	ar := AcademicRecord{StudentId: 1, CGPA: 9.0}
	applicant := Applicant{Student: student, AcademicRecord: ar}
	drive := NewDrive(time.Now(), time.Now().Add(24*time.Hour), "Engineer", 8.0, 1000000, Day)
	company := NewCompany("TestCorp")
	company.drives = []*Drive{drive}
	pr := PlacementRegistrar{
		companies:    []*Company{company},
		applications: []*Application{},
		applicants:   []*Applicant{&applicant},
	}
	return &StudentPlacementService{
		student:            student,
		drive:              *drive,
		applicant:          applicant,
		PlacementRegistrar: pr,
		Company:            *company,
	}
}

func TestStudentPlacementService_Apply(t *testing.T) {
	service := createTestPlacementService()
	service.PlacementRegistrar.applicants = []*Applicant{&service.applicant}
	service.PlacementRegistrar.companies = []*Company{&service.Company}
	service.PlacementRegistrar.applications = nil
	service.drive = service.drive
	service.Company = service.Company
	err := service.Apply()
	if err != nil {
		t.Errorf("Apply returned error: %v", err)
	}
}

func TestStudentPlacementService_CompaniesApplicable(t *testing.T) {
	service := createTestPlacementService()
	applicable := service.CompaniesApplicable()
	if len(applicable) == 0 || applicable[0] != "TestCorp" {
		t.Errorf("expected TestCorp to be applicable, got %v", applicable)
	}
}

func TestStudentPlacementService_CompaniesApplied(t *testing.T) {
	service := createTestPlacementService()
	applied := service.CompaniesApplied()
	if applied == nil {
		t.Error("expected CompaniesApplied to return a value")
	}
}

func TestStudentPlacementService_GetDrive(t *testing.T) {
	service := createTestPlacementService()
	if service.GetDrive().ID() != service.drive.ID() {
		t.Errorf("expected drive ID %d, got %d", service.drive.ID(), service.GetDrive().ID())
	}
}

func TestStudentPlacementService_ViewOfferDetails(t *testing.T) {
	service := createTestPlacementService()
	// Just ensure it doesn't panic
	service.ViewOfferDetails("TestCorp", service.drive.ID())
}

func TestStudentPlacementService_ViewShortlistStatus(t *testing.T) {
	service := createTestPlacementService()
	// Just ensure it doesn't panic
	service.ViewShortlistStatus(service.student.ID())
}

func TestDriveNotification_Send(t *testing.T) {
	drive := Drive{id: 1, roleName: "Engineer"}
	n := NewDriveNotification(drive)
	returned := n.Send().(Drive)
	if returned.id != drive.id || returned.roleName != drive.roleName {
		t.Errorf("DriveNotification did not return correct drive")
	}
}

func TestResultNotification_Send(t *testing.T) {
	n := NewResultNotification(5)
	if n.Send() != 5 {
		t.Errorf("ResultNotification did not return correct days left")
	}
}
