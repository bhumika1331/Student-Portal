package internal

import (
	"testing"
	"time"
)

func TestJobCategory_String(t *testing.T) {
	t.Run("should return correct string for Day", func(t *testing.T) {
		if Day.String() != "Day Company" {
			t.Errorf("Expected 'Day Company', got '%s'", Day.String())
		}
	})

	t.Run("should return correct string for Dream", func(t *testing.T) {
		if Dream.String() != "Dream" {
			t.Errorf("Expected 'Dream', got '%s'", Dream.String())
		}
	})

	t.Run("should return correct string for SuperDream", func(t *testing.T) {
		if SuperDream.String() != "Super Dream" {
			t.Errorf("Expected 'Super Dream', got '%s'", SuperDream.String())
		}
	})

	t.Run("should return correct string for Marquee", func(t *testing.T) {
		if Marquee.String() != "Marquee" {
			t.Errorf("Expected 'Marquee', got '%s'", Marquee.String())
		}
	})

	t.Run("should handle invalid job category", func(t *testing.T) {
		invalidCategory := JobCategory(999)
		result := invalidCategory.String()
		if result != "" {
			t.Errorf("Expected empty string for invalid category, got '%s'", result)
		}
	})
}

func TestJobCategory_Constants(t *testing.T) {
	t.Run("should have correct enum values", func(t *testing.T) {
		if Day != 0 {
			t.Errorf("Day should be 0, got %d", Day)
		}
		if Dream != 1 {
			t.Errorf("Dream should be 1, got %d", Dream)
		}
		if SuperDream != 2 {
			t.Errorf("SuperDream should be 2, got %d", SuperDream)
		}
		if Marquee != 3 {
			t.Errorf("Marquee should be 3, got %d", Marquee)
		}
	})
}

func TestNewEligibility(t *testing.T) {
	t.Run("should create eligibility with positive GPA", func(t *testing.T) {
		el := NewEligibility(7.5)
		if el.Requirement() != 7.5 {
			t.Errorf("Expected 7.5, got %f", el.Requirement())
		}
	})

	t.Run("should create eligibility with zero GPA", func(t *testing.T) {
		el := NewEligibility(0.0)
		if el.Requirement() != 0.0 {
			t.Error("NewEligibility should handle zero GPA")
		}
	})

	t.Run("should create eligibility with negative GPA", func(t *testing.T) {
		el := NewEligibility(-1.0)
		if el.Requirement() != -1.0 {
			t.Error("NewEligibility should handle negative GPA")
		}
	})

	t.Run("should create eligibility with very high GPA", func(t *testing.T) {
		el := NewEligibility(10.0)
		if el.Requirement() != 10.0 {
			t.Error("NewEligibility should handle high GPA")
		}
	})

	t.Run("should create eligibility with decimal GPA", func(t *testing.T) {
		el := NewEligibility(8.75)
		if el.Requirement() != 8.75 {
			t.Error("NewEligibility should handle decimal GPA")
		}
	})
}

func TestEligibility_Requirement(t *testing.T) {
	t.Run("should return correct requirement", func(t *testing.T) {
		el := NewEligibility(6.0)
		if el.Requirement() != 6.0 {
			t.Error("Requirement() failed")
		}
	})

	t.Run("should return consistent value", func(t *testing.T) {
		el := NewEligibility(8.5)
		req1 := el.Requirement()
		req2 := el.Requirement()
		if req1 != req2 {
			t.Error("Requirement() should return consistent value")
		}
	})
}

func TestEligibility_ChangeRequirement(t *testing.T) {
	t.Run("should update requirement", func(t *testing.T) {
		el := NewEligibility(6.0)
		el.ChangeRequirement(8.0)
		if el.Requirement() != 8.0 {
			t.Error("ChangeRequirement did not update requirement")
		}
	})

	t.Run("should handle zero requirement", func(t *testing.T) {
		el := NewEligibility(7.0)
		el.ChangeRequirement(0.0)
		if el.Requirement() != 0.0 {
			t.Error("ChangeRequirement should handle zero")
		}
	})

	t.Run("should handle negative requirement", func(t *testing.T) {
		el := NewEligibility(7.0)
		el.ChangeRequirement(-1.0)
		if el.Requirement() != -1.0 {
			t.Error("ChangeRequirement should handle negative values")
		}
	})

	t.Run("should handle multiple changes", func(t *testing.T) {
		el := NewEligibility(6.0)
		el.ChangeRequirement(7.0)
		el.ChangeRequirement(8.0)
		el.ChangeRequirement(9.0)
		if el.Requirement() != 9.0 {
			t.Error("ChangeRequirement should handle multiple changes")
		}
	})
}

func TestEligibility_checkEligibility(t *testing.T) {
	t.Run("should pass for eligible applicant", func(t *testing.T) {
		el := NewEligibility(8.0)
		a := NewApplicant(Student{id: 1, name: "Bob"}, AcademicRecord{CGPA: 8.1})
		if !el.checkEligibility(a) {
			t.Error("Eligibility check should pass")
		}
	})

	t.Run("should fail for ineligible applicant", func(t *testing.T) {
		el := NewEligibility(8.0)
		a := NewApplicant(Student{id: 2}, AcademicRecord{CGPA: 7.0})
		if el.checkEligibility(a) {
			t.Error("Eligibility check should fail")
		}
	})

	t.Run("should pass for exact GPA match", func(t *testing.T) {
		el := NewEligibility(8.0)
		a := NewApplicant(Student{id: 3}, AcademicRecord{CGPA: 8.0})
		if el.checkEligibility(a) {
			t.Error("Eligibility check should fail for exact match (requirement > CGPA)")
		}
	})

	t.Run("should handle zero requirement", func(t *testing.T) {
		el := NewEligibility(0.0)
		a := NewApplicant(Student{id: 4}, AcademicRecord{CGPA: 5.0})
		if !el.checkEligibility(a) {
			t.Error("Zero requirement should pass for any positive CGPA")
		}
	})

	t.Run("should handle zero CGPA", func(t *testing.T) {
		el := NewEligibility(8.0)
		a := NewApplicant(Student{id: 5}, AcademicRecord{CGPA: 0.0})
		if el.checkEligibility(a) {
			t.Error("Should fail for zero CGPA")
		}
	})

	t.Run("should handle negative values", func(t *testing.T) {
		el := NewEligibility(-1.0)
		a := NewApplicant(Student{id: 6}, AcademicRecord{CGPA: 5.0})
		if !el.checkEligibility(a) {
			t.Error("Negative requirement should pass for positive CGPA")
		}
	})

	t.Run("should handle nil applicant", func(t *testing.T) {
		el := NewEligibility(8.0)
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic for nil applicant")
			}
		}()
		el.checkEligibility(nil)
	})
}

func TestNewDrive(t *testing.T) {
	t.Run("should create drive with all parameters", func(t *testing.T) {
		start := time.Now()
		end := start.Add(24 * time.Hour)
		dr := NewDrive(start, end, "SWE", 8.0, 100, Dream)

		if dr.StartDate() != start {
			t.Error("StartDate not set correctly")
		}
		if dr.EndDate() != end {
			t.Error("EndDate not set correctly")
		}
		if dr.RoleName() != "SWE" {
			t.Error("RoleName not set correctly")
		}
		if dr.Eligibility().Requirement() != 8.0 {
			t.Error("Eligibility not set correctly")
		}
		if dr.CTC() != 100 {
			t.Error("CTC not set correctly")
		}
		if dr.JobCategory() != Dream {
			t.Error("JobCategory not set correctly")
		}
		if dr.ID() <= 0 {
			t.Error("ID should be positive")
		}
	})

	t.Run("should create drive with zero values", func(t *testing.T) {
		dr := NewDrive(time.Time{}, time.Time{}, "", 0.0, 0, Day)
		if dr == nil {
			t.Error("NewDrive should not return nil")
		}
	})

	t.Run("should create drive with special characters in role", func(t *testing.T) {
		dr := NewDrive(time.Now(), time.Now(), "C++ Developer & QA", 7.0, 50, Marquee)
		if dr.RoleName() != "C++ Developer & QA" {
			t.Error("Should handle special characters in role name")
		}
	})

	t.Run("should generate unique IDs", func(t *testing.T) {
		dr1 := NewDrive(time.Now(), time.Now(), "Role1", 7.0, 100, Day)
		dr2 := NewDrive(time.Now(), time.Now(), "Role2", 8.0, 200, Dream)
		if dr1.ID() == dr2.ID() {
			t.Error("Each drive should have unique ID")
		}
	})
}

func TestDrive_Getters(t *testing.T) {
	start := time.Now()
	end := start.Add(48 * time.Hour)
	dr := NewDrive(start, end, "DevOps", 7.5, 150, SuperDream)

	t.Run("should return correct ID", func(t *testing.T) {
		if dr.ID() <= 0 {
			t.Error("ID should be positive")
		}
	})

	t.Run("should return correct start date", func(t *testing.T) {
		if dr.StartDate() != start {
			t.Error("StartDate getter failed")
		}
	})

	t.Run("should return correct end date", func(t *testing.T) {
		if dr.EndDate() != end {
			t.Error("EndDate getter failed")
		}
	})

	t.Run("should return correct role name", func(t *testing.T) {
		if dr.RoleName() != "DevOps" {
			t.Error("RoleName getter failed")
		}
	})

	t.Run("should return correct eligibility", func(t *testing.T) {
		el := dr.Eligibility()
		if el.Requirement() != 7.5 {
			t.Error("Eligibility getter failed")
		}
	})

	t.Run("should return correct CTC", func(t *testing.T) {
		if dr.CTC() != 150 {
			t.Error("CTC getter failed")
		}
	})

	t.Run("should return correct job category", func(t *testing.T) {
		if dr.JobCategory() != SuperDream {
			t.Error("JobCategory getter failed")
		}
	})

	t.Run("should return empty applications initially", func(t *testing.T) {
		apps := dr.Applications()
		if len(apps) != 0 {
			t.Error("Applications should be empty initially")
		}
	})
}

func TestDrive_Setters(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)

	t.Run("should set start date", func(t *testing.T) {
		newStart := time.Now().Add(time.Hour)
		dr.SetStartDate(newStart)
		if dr.StartDate() != newStart {
			t.Error("SetStartDate failed")
		}
	})

	t.Run("should set end date", func(t *testing.T) {
		newEnd := time.Now().Add(2 * time.Hour)
		dr.SetEndDate(newEnd)
		if dr.EndDate() != newEnd {
			t.Error("SetEndDate failed")
		}
	})

	t.Run("should set role name", func(t *testing.T) {
		dr.SetRoleName("Senior QA")
		if dr.RoleName() != "Senior QA" {
			t.Error("SetRoleName failed")
		}
	})

	t.Run("should set eligibility", func(t *testing.T) {
		dr.SetEligibility(9.0)
		if dr.Eligibility().Requirement() != 9.0 {
			t.Error("SetEligibility failed")
		}
	})

	t.Run("should set CTC", func(t *testing.T) {
		dr.SetCTC(200)
		if dr.CTC() != 200 {
			t.Error("SetCTC failed")
		}
	})

	t.Run("should set job category", func(t *testing.T) {
		dr.SetJobCategory(Marquee)
		if dr.JobCategory() != Marquee {
			t.Error("SetJobCategory failed")
		}
	})

	t.Run("should handle empty role name", func(t *testing.T) {
		dr.SetRoleName("")
		if dr.RoleName() != "" {
			t.Error("SetRoleName should handle empty string")
		}
	})

	t.Run("should handle zero CTC", func(t *testing.T) {
		dr.SetCTC(0)
		if dr.CTC() != 0 {
			t.Error("SetCTC should handle zero")
		}
	})

	t.Run("should handle negative CTC", func(t *testing.T) {
		dr.SetCTC(-100)
		if dr.CTC() != -100 {
			t.Error("SetCTC should handle negative values")
		}
	})
}

func TestDrive_AppendApplication(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)

	t.Run("should append single application", func(t *testing.T) {
		app := &Application{id: 1, Applicant: &Applicant{Student: Student{id: 1}}, status: Applied}
		dr.AppendApplication(app)

		apps := dr.Applications()
		if len(apps) != 1 || apps[0] != app {
			t.Error("AppendApplication failed")
		}
	})

	t.Run("should append multiple applications", func(t *testing.T) {
		dr2 := NewDrive(time.Now(), time.Now(), "Dev", 7.0, 100, Dream)
		app1 := &Application{id: 1, status: Applied}
		app2 := &Application{id: 2, status: ShortListed}

		dr2.AppendApplication(app1)
		dr2.AppendApplication(app2)

		apps := dr2.Applications()
		if len(apps) != 2 || apps[0] != app1 || apps[1] != app2 {
			t.Error("AppendApplication failed for multiple applications")
		}
	})

	t.Run("should not append nil application", func(t *testing.T) {
		dr3 := NewDrive(time.Now(), time.Now(), "Test", 7.0, 100, Day)
		dr3.AppendApplication(nil)

		apps := dr3.Applications()
		if len(apps) > 0  {
			t.Error("AppendApplication should not append nil application")
		}
	})

	t.Run("should maintain application order", func(t *testing.T) {
		dr4 := NewDrive(time.Now(), time.Now(), "Order", 7.0, 100, Day)
		apps := make([]*Application, 5)
		for i := 0; i < 5; i++ {
			apps[i] = &Application{id: i + 1}
			dr4.AppendApplication(apps[i])
		}

		result := dr4.Applications()
		for i, app := range apps {
			if result[i] != app {
				t.Errorf("Application order not maintained at index %d", i)
			}
		}
	})
}

func TestDrive_HasApplied(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)

	t.Run("should return true for existing applicant", func(t *testing.T) {
		app := &Application{id: 1, Applicant: &Applicant{Student: Student{id: 123}}, status: Applied}
		dr.AppendApplication(app)

		if !dr.HasApplied(123) {
			t.Error("HasApplied should return true for existing applicant")
		}
	})

	t.Run("should return false for non-existing applicant", func(t *testing.T) {
		dr2 := NewDrive(time.Now(), time.Now(), "Dev", 7.0, 100, Dream)
		if dr2.HasApplied(999) {
			t.Error("HasApplied should return false for non-existing applicant")
		}
	})

	t.Run("should return false for zero ID", func(t *testing.T) {
		if dr.HasApplied(0) {
			t.Error("HasApplied should return false for zero ID")
		}
	})

	t.Run("should return false for negative ID", func(t *testing.T) {
		if dr.HasApplied(-1) {
			t.Error("HasApplied should return false for negative ID")
		}
	})

	t.Run("should handle multiple applicants", func(t *testing.T) {
		dr3 := NewDrive(time.Now(), time.Now(), "Multi", 7.0, 100, Day)
		app1 := &Application{id: 1, Applicant: &Applicant{Student: Student{id: 100}}, status: Applied}
		app2 := &Application{id: 2, Applicant: &Applicant{Student: Student{id: 200}}, status: Applied}
		dr3.AppendApplication(app1)
		dr3.AppendApplication(app2)

		if !dr3.HasApplied(100) || !dr3.HasApplied(200) {
			t.Error("HasApplied should return true for both applicants")
		}
		if dr3.HasApplied(300) {
			t.Error("HasApplied should return false for non-existing applicant")
		}
	})

	t.Run("should handle nil applications gracefully", func(t *testing.T) {
		dr4 := NewDrive(time.Now(), time.Now(), "Nil", 7.0, 100, Day)
		dr4.AppendApplication(nil)

		// Should not panic and should return false
		if dr4.HasApplied(123) {
			t.Error("HasApplied should return false when application is nil")
		}
	})
}

func TestDrive_GetApplicationByID(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)

	t.Run("should return application for existing ID", func(t *testing.T) {
		app := &Application{id: 123, status: Applied}
		dr.AppendApplication(app)

		got, err := dr.GetApplicationByID(123)
		if err != nil || got != app {
			t.Error("GetApplicationByID failed for existing ID")
		}
	})

	t.Run("should return error for non-existing ID", func(t *testing.T) {
		_, err := dr.GetApplicationByID(999)
		if err == nil {
			t.Error("GetApplicationByID should fail for missing ID")
		}
		if err.Error() != "no such application for given id" {
			t.Errorf("Expected specific error message, got: %v", err.Error())
		}
	})

	t.Run("should handle zero ID", func(t *testing.T) {
		app := &Application{id: 0, status: Applied}
		dr.AppendApplication(app)

		got, err := dr.GetApplicationByID(0)
		if err != nil || got != app {
			t.Error("GetApplicationByID should handle zero ID")
		}
	})

	t.Run("should handle negative ID", func(t *testing.T) {
		app := &Application{id: -1, status: Applied}
		dr.AppendApplication(app)

		got, err := dr.GetApplicationByID(-1)
		if err != nil || got != app {
			t.Error("GetApplicationByID should handle negative ID")
		}
	})

	t.Run("should return first match for duplicate IDs", func(t *testing.T) {
		dr2 := NewDrive(time.Now(), time.Now(), "Duplicate", 7.0, 100, Day)
		app1 := &Application{id: 123, status: Applied}
		app2 := &Application{id: 123, status: Selected}
		dr2.AppendApplication(app1)
		dr2.AppendApplication(app2)

		got, err := dr2.GetApplicationByID(123)
		if err != nil || got != app1 {
			t.Error("GetApplicationByID should return first match")
		}
	})
}

func TestDrive_getSelectedApplications(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)

	t.Run("should return selected applications", func(t *testing.T) {
		app1 := &Application{id: 1, status: Selected}
		app2 := &Application{id: 2, status: Applied}
		app3 := &Application{id: 3, status: Selected}
		dr.AppendApplication(app1)
		dr.AppendApplication(app2)
		dr.AppendApplication(app3)

		selected := dr.getSelectedApplications()
		if len(selected) != 2 {
			t.Errorf("Expected 2 selected applications, got %d", len(selected))
		}
		if selected[0] != app1 || selected[1] != app3 {
			t.Error("getSelectedApplications returned wrong applications")
		}
	})

	t.Run("should return empty for no selected applications", func(t *testing.T) {
		dr2 := NewDrive(time.Now(), time.Now(), "NoSelected", 7.0, 100, Day)
		app1 := &Application{id: 1, status: Applied}
		app2 := &Application{id: 2, status: ShortListed}
		dr2.AppendApplication(app1)
		dr2.AppendApplication(app2)

		selected := dr2.getSelectedApplications()
		if len(selected) != 0 {
			t.Error("getSelectedApplications should return empty slice")
		}
	})

	t.Run("should handle all selected applications", func(t *testing.T) {
		dr3 := NewDrive(time.Now(), time.Now(), "AllSelected", 7.0, 100, Day)
		app1 := &Application{id: 1, status: Selected}
		app2 := &Application{id: 2, status: Selected}
		dr3.AppendApplication(app1)
		dr3.AppendApplication(app2)

		selected := dr3.getSelectedApplications()
		if len(selected) != 2 {
			t.Error("getSelectedApplications should return all applications")
		}
	})
}

func TestDrive_getShortlistedApplications(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)

	t.Run("should return shortlisted applications", func(t *testing.T) {
		app1 := &Application{id: 1, status: ShortListed}
		app2 := &Application{id: 2, status: Applied}
		app3 := &Application{id: 3, status: ShortListed}
		dr.AppendApplication(app1)
		dr.AppendApplication(app2)
		dr.AppendApplication(app3)

		shortlisted := dr.getShortlistedApplications()
		if len(shortlisted) != 2 {
			t.Errorf("Expected 2 shortlisted applications, got %d", len(shortlisted))
		}
		if shortlisted[0] != app1 || shortlisted[1] != app3 {
			t.Error("getShortlistedApplications returned wrong applications")
		}
	})

	t.Run("should return empty for no shortlisted applications", func(t *testing.T) {
		dr2 := NewDrive(time.Now(), time.Now(), "NoShort", 7.0, 100, Day)
		app1 := &Application{id: 1, status: Applied}
		app2 := &Application{id: 2, status: Selected}
		dr2.AppendApplication(app1)
		dr2.AppendApplication(app2)

		shortlisted := dr2.getShortlistedApplications()
		if len(shortlisted) != 0 {
			t.Error("getShortlistedApplications should return empty slice")
		}
	})

	t.Run("should handle all status types", func(t *testing.T) {
		dr3 := NewDrive(time.Now(), time.Now(), "AllStatus", 7.0, 100, Day)
		statuses := []ApplicationStatus{Applied, ShortListed, Cleared, Selected, Rejected}
		for i, status := range statuses {
			app := &Application{id: i + 1, status: status}
			dr3.AppendApplication(app)
		}

		shortlisted := dr3.getShortlistedApplications()
		if len(shortlisted) != 1 {
			t.Errorf("Expected 1 shortlisted application, got %d", len(shortlisted))
		}
		if shortlisted[0].Status() != ShortListed {
			t.Error("Should return only ShortListed applications")
		}
	})
}
