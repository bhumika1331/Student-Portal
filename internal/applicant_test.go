package internal

import (
	"testing"
)

func TestNewApplicant(t *testing.T) {
	t.Run("should create new applicant with valid data", func(t *testing.T) {
		student := Student{id: 1, name: "Alice"}
		record := AcademicRecord{CGPA: 8.5}
		applicant := NewApplicant(student, record)

		if applicant.Student.id != 1 {
			t.Errorf("Expected student id 1, got %d", applicant.Student.id)
		}
		if applicant.Student.name != "Alice" {
			t.Errorf("Expected student name Alice, got %s", applicant.Student.name)
		}
		if applicant.AcademicRecord.CGPA != 8.5 {
			t.Errorf("Expected CGPA 8.5, got %f", applicant.AcademicRecord.CGPA)
		}
	})

	t.Run("should create applicant with zero values", func(t *testing.T) {
		student := Student{}
		record := AcademicRecord{}
		applicant := NewApplicant(student, record)

		if applicant == nil {
			t.Error("NewApplicant should not return nil")
		}
	})

	t.Run("should create applicant with negative values", func(t *testing.T) {
		student := Student{id: -1, name: ""}
		record := AcademicRecord{CGPA: -1.0}
		applicant := NewApplicant(student, record)

		if applicant.Student.id != -1 {
			t.Errorf("Expected student id -1, got %d", applicant.Student.id)
		}
	})
}

func TestApplicant_AddDrivesAppliedFor(t *testing.T) {
	t.Run("should add single drive", func(t *testing.T) {
		a := NewApplicant(Student{id: 1, name: "Alice"}, AcademicRecord{CGPA: 8.5})
		d1 := &Drive{id: 101}
		a.AddDrivesAppliedFor(d1)

		if len(a.drivesAppliedFor) != 1 {
			t.Errorf("expected 1 drive, got %d", len(a.drivesAppliedFor))
		}
		if a.drivesAppliedFor[0] != d1 {
			t.Error("drive not added correctly")
		}
	})

	t.Run("should add multiple drives", func(t *testing.T) {
		a := NewApplicant(Student{id: 1, name: "Alice"}, AcademicRecord{CGPA: 8.5})
		d1 := &Drive{id: 101}
		d2 := &Drive{id: 102}
		a.AddDrivesAppliedFor(d1)
		a.AddDrivesAppliedFor(d2)

		if len(a.drivesAppliedFor) != 2 {
			t.Errorf("expected 2 drives, got %d", len(a.drivesAppliedFor))
		}
	})

	t.Run("should handle nil drive", func(t *testing.T) {
		a := NewApplicant(Student{id: 1}, AcademicRecord{})
		a.AddDrivesAppliedFor(nil)

		if len(a.drivesAppliedFor) != 1 {
			t.Error("should add nil drive")
		}
		if a.drivesAppliedFor[0] != nil {
			t.Error("should store nil drive")
		}
	})

	t.Run("should add duplicate drives", func(t *testing.T) {
		a := NewApplicant(Student{id: 1}, AcademicRecord{})
		d1 := &Drive{id: 101}
		a.AddDrivesAppliedFor(d1)
		a.AddDrivesAppliedFor(d1)

		if len(a.drivesAppliedFor) != 2 {
			t.Errorf("expected 2 drives (duplicates allowed), got %d", len(a.drivesAppliedFor))
		}
	})
}

func TestApplicant_SetDrivesAppliedFor(t *testing.T) {
	t.Run("should set drives with valid slice", func(t *testing.T) {
		a := NewApplicant(Student{id: 2}, AcademicRecord{})
		d1 := &Drive{id: 201}
		d2 := &Drive{id: 202}
		a.SetDrivesAppliedFor([]*Drive{d1, d2})

		if len(a.drivesAppliedFor) != 2 {
			t.Error("SetDrivesAppliedFor failed")
		}
	})

	t.Run("should set empty slice", func(t *testing.T) {
		a := NewApplicant(Student{id: 2}, AcademicRecord{})
		a.AddDrivesAppliedFor(&Drive{id: 100}) // Add something first
		a.SetDrivesAppliedFor([]*Drive{})

		if len(a.drivesAppliedFor) != 0 {
			t.Error("SetDrivesAppliedFor should set empty slice")
		}
	})

	t.Run("should set nil slice", func(t *testing.T) {
		a := NewApplicant(Student{id: 2}, AcademicRecord{})
		a.SetDrivesAppliedFor(nil)

		if a.drivesAppliedFor != nil {
			t.Error("SetDrivesAppliedFor should set nil slice")
		}
	})

	t.Run("should replace existing drives", func(t *testing.T) {
		a := NewApplicant(Student{id: 2}, AcademicRecord{})
		d1 := &Drive{id: 201}
		d2 := &Drive{id: 202}
		d3 := &Drive{id: 203}

		a.SetDrivesAppliedFor([]*Drive{d1, d2})
		a.SetDrivesAppliedFor([]*Drive{d3})

		if len(a.drivesAppliedFor) != 1 || a.drivesAppliedFor[0] != d3 {
			t.Error("SetDrivesAppliedFor should replace existing drives")
		}
	})
}

func TestApplicant_DrivesAppliedFor(t *testing.T) {
	t.Run("should return all drives", func(t *testing.T) {
		a := NewApplicant(Student{id: 3}, AcademicRecord{})
		d1 := &Drive{id: 301}
		d2 := &Drive{id: 302}
		a.AddDrivesAppliedFor(d1)
		a.AddDrivesAppliedFor(d2)
		drives := a.DrivesAppliedFor()

		if len(drives) != 2 {
			t.Errorf("expected 2 drives, got %d", len(drives))
		}
	})

	t.Run("should return empty slice when no drives", func(t *testing.T) {
		a := NewApplicant(Student{id: 3}, AcademicRecord{})
		drives := a.DrivesAppliedFor()

		if drives == nil {
			t.Error("DrivesAppliedFor should return empty slice, not nil")
		}
		if len(drives) != 0 {
			t.Error("DrivesAppliedFor should return empty slice")
		}
	})

	t.Run("should return same reference", func(t *testing.T) {
		a := NewApplicant(Student{id: 3}, AcademicRecord{})
		d1 := &Drive{id: 301}
		a.AddDrivesAppliedFor(d1)

		drives1 := a.DrivesAppliedFor()
		drives2 := a.DrivesAppliedFor()

		if &drives1[0] != &drives2[0] {
			t.Error("DrivesAppliedFor should return same reference")
		}
	})
}

func TestApplicant_CompaniesAppliedFor(t *testing.T) {
	t.Run("should return unique company names", func(t *testing.T) {
		a := NewApplicant(Student{id: 4}, AcademicRecord{})
		c1 := &Company{id: 1, name: "CompanyA"}
		c2 := &Company{id: 2, name: "CompanyB"}
		d1 := &Drive{id: 401}
		d2 := &Drive{id: 402}
		c1.AddDrive(d1)
		c2.AddDrive(d2)

		pr := &PlacementRegistrar{
			companies: []*Company{c1, c2},
			applications: []*Application{
				{id: 1, driveId: 401, Applicant: a, status: Applied},
				{id: 2, driveId: 402, Applicant: a, status: Applied},
			},
		}

		names := a.CompaniesAppliedFor(pr)
		if len(names) != 2 {
			t.Errorf("expected 2 companies, got %d", len(names))
		}
	})

	t.Run("should handle duplicate company applications", func(t *testing.T) {
		a := NewApplicant(Student{id: 4}, AcademicRecord{})
		c1 := &Company{id: 1, name: "CompanyA"}
		d1 := &Drive{id: 401}
		d2 := &Drive{id: 402}
		c1.AddDrive(d1)
		c1.AddDrive(d2)

		pr := &PlacementRegistrar{
			companies: []*Company{c1},
			applications: []*Application{
				{id: 1, driveId: 401, Applicant: a, status: Applied},
				{id: 2, driveId: 402, Applicant: a, status: Applied},
			},
		}

		names := a.CompaniesAppliedFor(pr)
		if len(names) != 1 {
			t.Errorf("expected 1 unique company, got %d", len(names))
		}
	})

	t.Run("should handle empty applications", func(t *testing.T) {
		a := NewApplicant(Student{id: 4}, AcademicRecord{})
		pr := &PlacementRegistrar{
			companies:    []*Company{},
			applications: []*Application{},
		}

		names := a.CompaniesAppliedFor(pr)
		if len(names) != 0 {
			t.Error("should return empty slice for no applications")
		}
	})

	t.Run("should handle nil registrar", func(t *testing.T) {
		a := NewApplicant(Student{id: 4}, AcademicRecord{})

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for nil registrar")
			}
		}()

		a.CompaniesAppliedFor(nil)
	})
}

func TestApplicant_TotalNumberOfCompaniesAppliedFor(t *testing.T) {
	t.Run("should return correct count", func(t *testing.T) {
		a := NewApplicant(Student{id: 5}, AcademicRecord{})
		c1 := &Company{id: 1, name: "A"}
		c2 := &Company{id: 2, name: "B"}
		d1 := &Drive{id: 501}
		d2 := &Drive{id: 502}
		c1.AddDrive(d1)
		c2.AddDrive(d2)

		pr := &PlacementRegistrar{
			companies: []*Company{c1, c2},
			applications: []*Application{
				{id: 1, driveId: 501, Applicant: a, status: Applied},
				{id: 2, driveId: 502, Applicant: a, status: Applied},
			},
		}

		if a.TotalNumberOfCompaniesAppliedFor(pr) != 2 {
			t.Error("TotalNumberOfCompaniesAppliedFor failed")
		}
	})

	t.Run("should return zero for no applications", func(t *testing.T) {
		a := NewApplicant(Student{id: 5}, AcademicRecord{})
		pr := &PlacementRegistrar{
			companies:    []*Company{},
			applications: []*Application{},
		}

		if a.TotalNumberOfCompaniesAppliedFor(pr) != 0 {
			t.Error("should return 0 for no applications")
		}
	})
}

func TestApplicant_getAllRecivedOffersDrivesAndApplications(t *testing.T) {
	t.Run("should return offers for selected applications", func(t *testing.T) {
		a := NewApplicant(Student{id: 6}, AcademicRecord{})
		d1 := &Drive{id: 601}
		d2 := &Drive{id: 602}
		app1 := &Application{id: 1, driveId: 601, Applicant: a, status: Selected}
		app2 := &Application{id: 2, driveId: 602, Applicant: a, status: Applied}
		d1.applications = []*Application{app1}
		d2.applications = []*Application{app2}
		a.AddDrivesAppliedFor(d1)
		a.AddDrivesAppliedFor(d2)

		drives, apps := a.getAllReceivedOffersDrivesAndApplications()

		if len(drives) != 1 || len(apps) != 1 {
			t.Errorf("expected 1 offer, got %d drives and %d apps", len(drives), len(apps))
		}
		if drives[0] != d1 || apps[0] != app1 {
			t.Error("should return correct drive and application")
		}
	})

	t.Run("should return empty for no offers", func(t *testing.T) {
		a := NewApplicant(Student{id: 6}, AcademicRecord{})
		drives, apps := a.getAllReceivedOffersDrivesAndApplications()

		if len(drives) != 0 || len(apps) != 0 {
			t.Error("should return empty slices for no offers")
		}
	})

	t.Run("should handle multiple offers", func(t *testing.T) {
		a := NewApplicant(Student{id: 6}, AcademicRecord{})
		d1 := &Drive{id: 601}
		d2 := &Drive{id: 602}
		app1 := &Application{id: 1, driveId: 601, Applicant: a, status: Selected}
		app2 := &Application{id: 2, driveId: 602, Applicant: a, status: Selected}
		d1.applications = []*Application{app1}
		d2.applications = []*Application{app2}
		a.AddDrivesAppliedFor(d1)
		a.AddDrivesAppliedFor(d2)

		drives, apps := a.getAllReceivedOffersDrivesAndApplications()

		if len(drives) != 2 || len(apps) != 2 {
			t.Errorf("expected 2 offers, got %d drives and %d apps", len(drives), len(apps))
		}
	})

	t.Run("should handle applications from other applicants", func(t *testing.T) {
		a := NewApplicant(Student{id: 6}, AcademicRecord{})
		otherApplicant := NewApplicant(Student{id: 7}, AcademicRecord{})
		d1 := &Drive{id: 601}
		app1 := &Application{id: 1, driveId: 601, Applicant: a, status: Selected}
		app2 := &Application{id: 2, driveId: 601, Applicant: otherApplicant, status: Selected}
		d1.applications = []*Application{app1, app2}
		a.AddDrivesAppliedFor(d1)

		drives, apps := a.getAllReceivedOffersDrivesAndApplications()

		if len(drives) != 1 || len(apps) != 1 {
			t.Errorf("expected 1 offer for this applicant, got %d drives and %d apps", len(drives), len(apps))
		}
		if apps[0] != app1 {
			t.Error("should return only this applicant's offers")
		}
	})
}

func TestApplicant_getFinalOffer(t *testing.T) {
	t.Run("should return error when no offers exist", func(t *testing.T) {
		a := NewApplicant(Student{id: 7}, AcademicRecord{})
		_, err := a.getFinalOffer()

		if err == nil {
			t.Error("getFinalOffer should fail when no offers exist")
		}
		if err.Error() != "no offers yet" {
			t.Errorf("Expected error 'no offers yet', got '%v'", err.Error())
		}
	})

	t.Run("should return CTC when offers exist", func(t *testing.T) {
		a := NewApplicant(Student{id: 8}, AcademicRecord{})
		d := &Drive{id: 701, ctc: 100}
		app := &Application{id: 1, driveId: 701, Applicant: a, status: Selected}
		d.applications = []*Application{app}
		a.AddDrivesAppliedFor(d)

		ctc, err := a.getFinalOffer()
		if err != nil {
			t.Errorf("getFinalOffer should succeed when offers exist, got error: %v", err)
		}
		if ctc != 100 {
			t.Errorf("Expected CTC 100, got %d", ctc)
		}
	})

	t.Run("should return first offer when multiple offers exist", func(t *testing.T) {
		a := NewApplicant(Student{id: 9}, AcademicRecord{})
		d1 := &Drive{id: 801, ctc: 150}
		d2 := &Drive{id: 802, ctc: 200}
		app1 := &Application{id: 1, driveId: 801, Applicant: a, status: Selected}
		app2 := &Application{id: 2, driveId: 802, Applicant: a, status: Selected}
		d1.applications = []*Application{app1}
		d2.applications = []*Application{app2}
		a.AddDrivesAppliedFor(d1)
		a.AddDrivesAppliedFor(d2)

		ctc, err := a.getFinalOffer()
		if err != nil {
			t.Errorf("getFinalOffer should succeed with multiple offers, got error: %v", err)
		}
		if ctc != 150 {
			t.Errorf("Expected first offer CTC 150, got %d", ctc)
		}
	})

	t.Run("should handle zero CTC", func(t *testing.T) {
		a := NewApplicant(Student{id: 10}, AcademicRecord{})
		d := &Drive{id: 901, ctc: 0}
		app := &Application{id: 1, driveId: 901, Applicant: a, status: Selected}
		d.applications = []*Application{app}
		a.AddDrivesAppliedFor(d)

		ctc, err := a.getFinalOffer()
		if err != nil {
			t.Errorf("getFinalOffer should succeed with zero CTC, got error: %v", err)
		}
		if ctc != 0 {
			t.Errorf("Expected CTC 0, got %d", ctc)
		}
	})
}
