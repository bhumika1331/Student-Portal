package internal

import (
	"testing"
	"time"
)

func TestPlacementRegistrar_AddCompany(t *testing.T) {
	t.Run("should add single company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 1, name: "TestCo"}
		pr.AddCompany(c)

		if len(pr.companies) != 1 {
			t.Error("AddCompany failed to add company")
		}
		if pr.companies[0] != c {
			t.Error("AddCompany did not store correct company")
		}
	})

	t.Run("should add multiple companies", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c1 := &Company{id: 1, name: "TestCo1"}
		c2 := &Company{id: 2, name: "TestCo2"}
		pr.AddCompany(c1)
		pr.AddCompany(c2)

		if len(pr.companies) != 2 {
			t.Error("AddCompany failed to add multiple companies")
		}
	})

	t.Run("should add nil company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		pr.AddCompany(nil)

		if len(pr.companies) != 1 {
			t.Error("AddCompany should add nil company")
		}
		if pr.companies[0] != nil {
			t.Error("AddCompany should store nil company")
		}
	})

	t.Run("should add duplicate companies", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 1, name: "TestCo"}
		pr.AddCompany(c)
		pr.AddCompany(c)

		if len(pr.companies) != 2 {
			t.Error("AddCompany should allow duplicate companies")
		}
	})
}

func TestPlacementRegistrar_GetCompanyByID(t *testing.T) {
	t.Run("should get existing company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 1, name: "TestCo"}
		pr.AddCompany(c)

		got, err := pr.GetCompanyByID(1)
		if err != nil || got != c {
			t.Error("GetCompanyByID failed for existing company")
		}
	})

	t.Run("should fail for non-existing company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		_, err := pr.GetCompanyByID(999)
		if err == nil {
			t.Error("GetCompanyByID should fail for missing company")
		}
	})

	t.Run("should handle zero ID", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 0, name: "TestCo"}
		pr.AddCompany(c)

		got, err := pr.GetCompanyByID(0)
		if err != nil || got != c {
			t.Error("GetCompanyByID should handle zero ID")
		}
	})

	t.Run("should handle negative ID", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: -1, name: "TestCo"}
		pr.AddCompany(c)

		got, err := pr.GetCompanyByID(-1)
		if err != nil || got != c {
			t.Error("GetCompanyByID should handle negative ID")
		}
	})

	t.Run("should return first match for duplicate IDs", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c1 := &Company{id: 1, name: "First"}
		c2 := &Company{id: 1, name: "Second"}
		pr.AddCompany(c1)
		pr.AddCompany(c2)

		got, err := pr.GetCompanyByID(1)
		if err != nil || got != c1 {
			t.Error("GetCompanyByID should return first match")
		}
	})
}

func TestPlacementRegistrar_UpdateCompany(t *testing.T) {
	t.Run("should update existing company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 2, name: "Old"}
		pr.AddCompany(c)

		c2 := &Company{id: 2, name: "New"}
		err := pr.UpdateCompany(c2)
		if err != nil {
			t.Error("UpdateCompany failed")
		}

		got, _ := pr.GetCompanyByID(2)
		if got.Name() != "New" {
			t.Error("UpdateCompany did not update company")
		}
	})

	t.Run("should fail for non-existing company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 999, name: "New"}
		err := pr.UpdateCompany(c)
		if err == nil {
			t.Error("UpdateCompany should fail for non-existing company")
		}
	})

	t.Run("should handle nil company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		err := pr.UpdateCompany(nil)
		if err == nil {
			t.Error("UpdateCompany should fail for nil company")
		}
	})

	t.Run("should replace pointer reference", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c1 := &Company{id: 2, name: "Old"}
		pr.AddCompany(c1)

		c2 := &Company{id: 2, name: "New"}
		pr.UpdateCompany(c2)

		got, _ := pr.GetCompanyByID(2)
		if got == c1 {
			t.Error("UpdateCompany should replace pointer reference")
		}
		if got != c2 {
			t.Error("UpdateCompany should set new pointer reference")
		}
	})
}

func TestPlacementRegistrar_AddDriveToCompany(t *testing.T) {
	t.Run("should add drive to existing company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 3, name: "DriveCo"}
		pr.AddCompany(c)

		d := &Drive{id: 301}
		err := pr.AddDriveToCompany(3, d)
		if err != nil {
			t.Error("AddDriveToCompany failed")
		}

		if len(c.drives) != 1 || c.drives[0] != d {
			t.Error("Drive not added to company")
		}
	})

	t.Run("should fail for non-existing company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		d := &Drive{id: 301}
		err := pr.AddDriveToCompany(999, d)
		if err == nil {
			t.Error("AddDriveToCompany should fail for non-existing company")
		}
	})

	t.Run("should add multiple drives", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 3, name: "DriveCo"}
		pr.AddCompany(c)

		d1 := &Drive{id: 301}
		d2 := &Drive{id: 302}
		pr.AddDriveToCompany(3, d1)
		pr.AddDriveToCompany(3, d2)

		if len(c.drives) != 2 {
			t.Error("AddDriveToCompany should add multiple drives")
		}
	})

	t.Run("should add nil drive", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		c := &Company{id: 3, name: "DriveCo"}
		pr.AddCompany(c)

		err := pr.AddDriveToCompany(3, nil)
		if err != nil {
			t.Error("AddDriveToCompany should handle nil drive")
		}

		if len(c.drives) != 1 || c.drives[0] != nil {
			t.Error("AddDriveToCompany should add nil drive")
		}
	})
}

func TestPlacementRegistrar_ApplicantByID(t *testing.T) {
	t.Run("should find existing applicant", func(t *testing.T) {
		a := NewApplicant(Student{id: 10, name: "Test"}, AcademicRecord{})
		pr := &PlacementRegistrar{applicants: []*Applicant{a}}

		got, err := pr.ApplicantByID(10)
		if err != nil || got != a {
			t.Error("ApplicantByID failed")
		}
	})

	t.Run("should fail for non-existing applicant", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		_, err := pr.ApplicantByID(999)
		if err == nil {
			t.Error("ApplicantByID should fail for missing applicant")
		}
	})

	t.Run("should handle zero ID", func(t *testing.T) {
		a := NewApplicant(Student{id: 0, name: "Test"}, AcademicRecord{})
		pr := &PlacementRegistrar{applicants: []*Applicant{a}}

		got, err := pr.ApplicantByID(0)
		if err != nil || got != a {
			t.Error("ApplicantByID should handle zero ID")
		}
	})

	t.Run("should handle negative ID", func(t *testing.T) {
		a := NewApplicant(Student{id: -1, name: "Test"}, AcademicRecord{})
		pr := &PlacementRegistrar{applicants: []*Applicant{a}}

		got, err := pr.ApplicantByID(-1)
		if err != nil || got != a {
			t.Error("ApplicantByID should handle negative ID")
		}
	})

	t.Run("should return first match for duplicate IDs", func(t *testing.T) {
		a1 := NewApplicant(Student{id: 1, name: "First"}, AcademicRecord{})
		a2 := NewApplicant(Student{id: 1, name: "Second"}, AcademicRecord{})
		pr := &PlacementRegistrar{applicants: []*Applicant{a1, a2}}

		got, err := pr.ApplicantByID(1)
		if err != nil || got != a1 {
			t.Error("ApplicantByID should return first match")
		}
	})
}

func TestPlacementRegistrar_DriveByID(t *testing.T) {
	t.Run("should find existing drive", func(t *testing.T) {
		d := &Drive{id: 401}
		c := &Company{id: 40, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}}

		got, err := pr.DriveByID(40, 401)
		if err != nil || got != d {
			t.Error("DriveByID failed")
		}
	})

	t.Run("should fail for non-existing drive", func(t *testing.T) {
		c := &Company{id: 40, name: "Test", drives: []*Drive{}}
		pr := &PlacementRegistrar{companies: []*Company{c}}

		_, err := pr.DriveByID(40, 999)
		if err == nil {
			t.Error("DriveByID should fail for missing drive")
		}
	})

	t.Run("should fail for non-existing company", func(t *testing.T) {
		pr := &PlacementRegistrar{}
		_, err := pr.DriveByID(999, 401)
		if err == nil {
			t.Error("DriveByID should fail for missing company")
		}
	})

	t.Run("should handle zero IDs", func(t *testing.T) {
		d := &Drive{id: 0}
		c := &Company{id: 0, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}}

		got, err := pr.DriveByID(0, 0)
		if err != nil || got != d {
			t.Error("DriveByID should handle zero IDs")
		}
	})

	t.Run("should handle negative IDs", func(t *testing.T) {
		d := &Drive{id: -1}
		c := &Company{id: -1, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}}

		got, err := pr.DriveByID(-1, -1)
		if err != nil || got != d {
			t.Error("DriveByID should handle negative IDs")
		}
	})
}

func TestPlacementRegistrar_ApplyForDrive(t *testing.T) {
	t.Run("should apply for eligible drive", func(t *testing.T) {
		a := NewApplicant(Student{id: 50, name: "Test"}, AcademicRecord{CGPA: 9.0})
		d := NewDrive(time.Now(), time.Now(), "Dev", 8.0, 100, Dream)
		c := &Company{id: 50, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}, applicants: []*Applicant{a}}

		err := pr.ApplyForDrive(50, 50, d.ID())
		if err != nil {
			t.Errorf("ApplyForDrive failed: %v", err)
		}

		if len(pr.applications) != 1 {
			t.Error("Application not added to registrar")
		}
		if len(d.Applications()) != 1 {
			t.Error("Application not added to drive")
		}
		if len(a.DrivesAppliedFor()) != 1 {
			t.Error("Drive not added to applicant")
		}
	})

	t.Run("should fail for duplicate application", func(t *testing.T) {
		a := NewApplicant(Student{id: 50, name: "Test"}, AcademicRecord{CGPA: 9.0})
		d := NewDrive(time.Now(), time.Now(), "Dev", 8.0, 100, Dream)
		c := &Company{id: 50, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}, applicants: []*Applicant{a}}

		pr.ApplyForDrive(50, 50, d.ID())
		err := pr.ApplyForDrive(50, 50, d.ID())
		if err == nil {
			t.Error("ApplyForDrive should fail for duplicate application")
		}
	})

	t.Run("should fail for ineligible applicant", func(t *testing.T) {
		a := NewApplicant(Student{id: 51, name: "Bad"}, AcademicRecord{CGPA: 5.0})
		d := NewDrive(time.Now(), time.Now(), "Dev", 8.0, 100, Dream)
		c := &Company{id: 50, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}, applicants: []*Applicant{a}}

		err := pr.ApplyForDrive(51, 50, d.ID())
		if err == nil {
			t.Error("ApplyForDrive should fail for ineligible applicant")
		}
	})

	t.Run("should fail for non-existing applicant", func(t *testing.T) {
		d := NewDrive(time.Now(), time.Now(), "Dev", 8.0, 100, Dream)
		c := &Company{id: 50, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}, applicants: []*Applicant{}}

		err := pr.ApplyForDrive(999, 50, d.ID())
		if err == nil {
			t.Error("ApplyForDrive should fail for non-existing applicant")
		}
	})

	t.Run("should fail for non-existing drive", func(t *testing.T) {
		a := NewApplicant(Student{id: 50, name: "Test"}, AcademicRecord{CGPA: 9.0})
		c := &Company{id: 50, name: "Test", drives: []*Drive{}}
		pr := &PlacementRegistrar{companies: []*Company{c}, applicants: []*Applicant{a}}

		err := pr.ApplyForDrive(50, 50, 999)
		if err == nil {
			t.Error("ApplyForDrive should fail for non-existing drive")
		}
	})

	t.Run("should fail for non-existing company", func(t *testing.T) {
		a := NewApplicant(Student{id: 50, name: "Test"}, AcademicRecord{CGPA: 9.0})
		pr := &PlacementRegistrar{companies: []*Company{}, applicants: []*Applicant{a}}

		err := pr.ApplyForDrive(50, 999, 401)
		if err == nil {
			t.Error("ApplyForDrive should fail for non-existing company")
		}
	})
}

func TestPlacementRegistrar_UpdateApplicationStatus(t *testing.T) {
	t.Run("should update existing application status", func(t *testing.T) {
		a := NewApplicant(Student{id: 60}, AcademicRecord{})
		app := &Application{id: 1, driveId: 600, Applicant: a, status: Applied}
		pr := &PlacementRegistrar{applications: []*Application{app}}

		err := pr.UpdateApplicationStatus(60, 600, Selected)
		if err != nil {
			t.Error("UpdateApplicationStatus failed")
		}

		if app.status != Selected {
			t.Error("Application status not updated")
		}
	})

	t.Run("should fail for non-existing application", func(t *testing.T) {
		pr := &PlacementRegistrar{applications: []*Application{}}
		err := pr.UpdateApplicationStatus(999, 600, Selected)
		if err == nil {
			t.Error("UpdateApplicationStatus should fail for non-existing application")
		}
	})

	t.Run("should handle all status transitions", func(t *testing.T) {
		a := NewApplicant(Student{id: 60}, AcademicRecord{})
		app := &Application{id: 1, driveId: 600, Applicant: a, status: Applied}
		pr := &PlacementRegistrar{applications: []*Application{app}}

		statuses := []ApplicationStatus{ShortListed, Cleared, Selected, Rejected}
		for _, status := range statuses {
			err := pr.UpdateApplicationStatus(60, 600, status)
			if err != nil {
				t.Errorf("UpdateApplicationStatus failed for status %d", status)
			}
			if app.status != status {
				t.Errorf("Status not updated to %d", status)
			}
		}
	})

	t.Run("should match both student ID and drive ID", func(t *testing.T) {
		a1 := NewApplicant(Student{id: 60}, AcademicRecord{})
		a2 := NewApplicant(Student{id: 61}, AcademicRecord{})
		app1 := &Application{id: 1, driveId: 600, Applicant: a1, status: Applied}
		app2 := &Application{id: 2, driveId: 601, Applicant: a2, status: Applied}
		pr := &PlacementRegistrar{applications: []*Application{app1, app2}}

		err := pr.UpdateApplicationStatus(60, 600, Selected)
		if err != nil {
			t.Error("UpdateApplicationStatus failed")
		}

		if app1.status != Selected {
			t.Error("First application status not updated")
		}
		if app2.status != Applied {
			t.Error("Second application status should not be updated")
		}
	})
}

func TestPlacementRegistrar_AllDrives(t *testing.T) {
	t.Run("should return all drives from all companies", func(t *testing.T) {
		d1 := &Drive{id: 701}
		d2 := &Drive{id: 702}
		d3 := &Drive{id: 703}
		c1 := &Company{id: 70, drives: []*Drive{d1, d2}}
		c2 := &Company{id: 71, drives: []*Drive{d3}}
		pr := &PlacementRegistrar{companies: []*Company{c1, c2}}

		drives := pr.AllDrives()
		if len(drives) != 3 {
			t.Errorf("Expected 3 drives, got %d", len(drives))
		}
	})

	t.Run("should return empty slice for no companies", func(t *testing.T) {
		pr := &PlacementRegistrar{companies: []*Company{}}
		drives := pr.AllDrives()

		if len(drives) != 0 {
			t.Error("AllDrives should return empty slice for no companies")
		}
	})

	t.Run("should handle companies with no drives", func(t *testing.T) {
		c1 := &Company{id: 70, drives: []*Drive{}}
		c2 := &Company{id: 71, drives: []*Drive{}}
		pr := &PlacementRegistrar{companies: []*Company{c1, c2}}

		drives := pr.AllDrives()
		if len(drives) != 0 {
			t.Error("AllDrives should return empty slice for companies with no drives")
		}
	})

	t.Run("should handle mixed companies", func(t *testing.T) {
		d1 := &Drive{id: 701}
		c1 := &Company{id: 70, drives: []*Drive{d1}}
		c2 := &Company{id: 71, drives: []*Drive{}}
		pr := &PlacementRegistrar{companies: []*Company{c1, c2}}

		drives := pr.AllDrives()
		if len(drives) != 1 || drives[0] != d1 {
			t.Error("AllDrives should handle mixed companies correctly")
		}
	})
}

func TestPlacementRegistrar_GenerateReportByDrive(t *testing.T) {
	t.Run("should generate report for drive", func(t *testing.T) {
		d := &Drive{id: 801, ctc: 100}
		c := &Company{id: 80, name: "Test", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}}

		rep := pr.GenerateReportByDrive()
		if rep.drive != d || rep.company != c || rep.driveCTC != 100 {
			t.Error("GenerateReportByDrive failed")
		}
	})

	t.Run("should use last drive when multiple drives exist", func(t *testing.T) {
		d1 := &Drive{id: 801, ctc: 100}
		d2 := &Drive{id: 802, ctc: 200}
		c := &Company{id: 80, name: "Test", drives: []*Drive{d1, d2}}
		pr := &PlacementRegistrar{companies: []*Company{c}}

		rep := pr.GenerateReportByDrive()
		if rep.drive != d2 || rep.driveCTC != 200 {
			t.Error("GenerateReportByDrive should use last drive")
		}
	})

	t.Run("should handle zero values", func(t *testing.T) {
		d := &Drive{id: 0, ctc: 0}
		c := &Company{id: 0, name: "", drives: []*Drive{d}}
		pr := &PlacementRegistrar{companies: []*Company{c}}

		rep := pr.GenerateReportByDrive()
		if rep.driveCTC != 0 {
			t.Error("GenerateReportByDrive should handle zero CTC")
		}
	})
}

func TestPlacementRegistrar_GenerateFullReport(t *testing.T) {
	t.Run("should generate full report", func(t *testing.T) {
		a := NewApplicant(Student{id: 100}, AcademicRecord{})
		d := NewDrive(time.Now(), time.Now(), "Dev", 8.0, 100, Dream)
		app := &Application{id: 1, driveId: d.ID(), Applicant: a, status: Selected}
		c := &Company{id: 100, drives: []*Drive{d}}
		pr := &PlacementRegistrar{
			applicants:   []*Applicant{a},
			companies:    []*Company{c},
			applications: []*Application{app},
		}

		rep := pr.GenerateFullReport()
		if rep.totalComapanies != 1 || rep.totalOffersMade != 1 {
			t.Error("GenerateFullReport failed")
		}
	})

	t.Run("should handle no applications", func(t *testing.T) {
		c := &Company{id: 100, drives: []*Drive{}}
		pr := &PlacementRegistrar{
			applicants:   []*Applicant{},
			companies:    []*Company{c},
			applications: []*Application{},
		}

		rep := pr.GenerateFullReport()
		if rep.totalComapanies != 1 || rep.totalOffersMade != 0 {
			t.Error("GenerateFullReport should handle no applications")
		}
	})

	t.Run("should count offers by category", func(t *testing.T) {
		a1 := NewApplicant(Student{id: 100}, AcademicRecord{})
		a2 := NewApplicant(Student{id: 101}, AcademicRecord{})
		d1 := NewDrive(time.Now(), time.Now(), "Dev1", 8.0, 100, Dream)
		d2 := NewDrive(time.Now(), time.Now(), "Dev2", 8.0, 100, Day)
		app1 := &Application{id: 1, driveId: d1.ID(), Applicant: a1, status: Selected}
		app2 := &Application{id: 2, driveId: d2.ID(), Applicant: a2, status: Selected}
		c := &Company{id: 100, drives: []*Drive{d1, d2}}
		pr := &PlacementRegistrar{
			applicants:   []*Applicant{a1, a2},
			companies:    []*Company{c},
			applications: []*Application{app1, app2},
		}

		rep := pr.GenerateFullReport()
		if rep.totalOffersByCatagory[Dream] != 1 {
			t.Error("GenerateFullReport should count Dream offers")
		}
		if rep.totalOffersByCatagory[Day] != 1 {
			t.Error("GenerateFullReport should count Day offers")
		}
	})
}

func TestPlacementRegistrar_GenerateReportByStudent(t *testing.T) {
	t.Run("should generate report for student with offers", func(t *testing.T) {
		a := NewApplicant(Student{id: 90, name: "Test"}, AcademicRecord{CGPA: 9.0})
		d := NewDrive(time.Now(), time.Now(), "Dev", 8.0, 100, Dream)
		app := &Application{id: 1, driveId: d.ID(), Applicant: a, status: Selected}
		d.AppendApplication(app)
		a.AddDrivesAppliedFor(d)
		pr := &PlacementRegistrar{
			applicants:   []*Applicant{a},
			companies:    []*Company{{id: 90, drives: []*Drive{d}}},
			applications: []*Application{app},
		}

		reports := pr.GenerateReportByStudent()
		if len(reports) != 1 {
			t.Errorf("Expected 1 report, got %d", len(reports))
		}

		report := reports[0]
		if report.applicant != a {
			t.Error("Report applicant should match")
		}
		if report.ctcForFinalOffer != 100 {
			t.Errorf("Expected CTC 100, got %d", report.ctcForFinalOffer)
		}
	})

	t.Run("should handle student with no offers", func(t *testing.T) {
		a := NewApplicant(Student{id: 91, name: "NoOffers"}, AcademicRecord{CGPA: 8.0})
		d := NewDrive(time.Now(), time.Now(), "Dev", 7.0, 100, Dream)
		a.AddDrivesAppliedFor(d)
		pr := &PlacementRegistrar{
			applicants:   []*Applicant{a},
			companies:    []*Company{{id: 91, drives: []*Drive{d}}},
			applications: []*Application{},
		}

		reports := pr.GenerateReportByStudent()
		if len(reports) != 1 {
			t.Errorf("Expected 1 report, got %d", len(reports))
		}

		report := reports[0]
		if report.finalOffer != nil {
			t.Error("Final offer should be nil for no offers")
		}
		if report.ctcForFinalOffer != 0 {
			t.Error("CTC should be 0 for no offers")
		}
	})

	t.Run("should identify eligible roles", func(t *testing.T) {
		a := NewApplicant(Student{id: 92, name: "Eligible"}, AcademicRecord{CGPA: 8.5})
		d1 := NewDrive(time.Now(), time.Now(), "Dev1", 8.0, 100, Dream)
		d2 := NewDrive(time.Now(), time.Now(), "Dev2", 9.0, 150, Dream)
		d3 := NewDrive(time.Now(), time.Now(), "Dev3", 7.0, 80, Day)
		a.AddDrivesAppliedFor(d1)
		pr := &PlacementRegistrar{
			applicants: []*Applicant{a},
			companies: []*Company{
				{id: 1, drives: []*Drive{d1, d2, d3}},
			},
			applications: []*Application{},
		}

		reports := pr.GenerateReportByStudent()
		if len(reports) != 1 {
			t.Errorf("Expected 1 report, got %d", len(reports))
		}

		report := reports[0]
		if len(report.eligibileRoles) != 2 {
			t.Errorf("Expected 2 eligible roles, got %d", len(report.eligibileRoles))
		}
	})

	t.Run("should handle multiple students", func(t *testing.T) {
		a1 := NewApplicant(Student{id: 93, name: "Student1"}, AcademicRecord{CGPA: 8.0})
		a2 := NewApplicant(Student{id: 94, name: "Student2"}, AcademicRecord{CGPA: 7.0})
		d1 := NewDrive(time.Now(), time.Now(), "Dev", 7.5, 100, Dream)
		a1.AddDrivesAppliedFor(d1)
		a2.AddDrivesAppliedFor(d1)
		pr := &PlacementRegistrar{
			applicants: []*Applicant{a1, a2},
			companies: []*Company{
				{id: 1, drives: []*Drive{d1}},
			},
			applications: []*Application{},
		}

		reports := pr.GenerateReportByStudent()
		if len(reports) != 2 {
			t.Errorf("Expected 2 reports, got %d", len(reports))
		}

		studentIDs := make(map[int]bool)
		for _, report := range reports {
			studentIDs[report.applicant.ID()] = true
		}

		if !studentIDs[93] || !studentIDs[94] {
			t.Error("Both students should be represented in reports")
		}
	})
}
