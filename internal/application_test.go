package internal

import "testing"

func TestApplication_ID(t *testing.T) {
	t.Run("should return positive ID", func(t *testing.T) {
		app := &Application{id: 1, status: Applied}
		if app.ID() != 1 {
			t.Error("ID() did not return correct id")
		}
	})

	t.Run("should return zero ID", func(t *testing.T) {
		app := &Application{id: 0, status: Applied}
		if app.ID() != 0 {
			t.Error("ID() should return zero")
		}
	})

	t.Run("should return negative ID", func(t *testing.T) {
		app := &Application{id: -1, status: Applied}
		if app.ID() != -1 {
			t.Error("ID() should return negative id")
		}
	})
}

func TestApplication_Status(t *testing.T) {
	t.Run("should return Applied status", func(t *testing.T) {
		app := &Application{id: 1, status: Applied}
		if app.Status() != Applied {
			t.Error("Status() did not return Applied")
		}
	})

	t.Run("should return ShortListed status", func(t *testing.T) {
		app := &Application{id: 1, status: ShortListed}
		if app.Status() != ShortListed {
			t.Error("Status() did not return ShortListed")
		}
	})

	t.Run("should return Cleared status", func(t *testing.T) {
		app := &Application{id: 1, status: Cleared}
		if app.Status() != Cleared {
			t.Error("Status() did not return Cleared")
		}
	})

	t.Run("should return Selected status", func(t *testing.T) {
		app := &Application{id: 1, status: Selected}
		if app.Status() != Selected {
			t.Error("Status() did not return Selected")
		}
	})

	t.Run("should return Rejected status", func(t *testing.T) {
		app := &Application{id: 1, status: Rejected}
		if app.Status() != Rejected {
			t.Error("Status() did not return Rejected")
		}
	})
}

func TestApplication_Fields(t *testing.T) {
	t.Run("should handle all fields correctly", func(t *testing.T) {
		applicant := NewApplicant(Student{id: 1, name: "Test"}, AcademicRecord{})
		app := &Application{
			id:        123,
			driveId:   456,
			Applicant: applicant,
			status:    Selected,
		}

		if app.ID() != 123 {
			t.Error("ID field not set correctly")
		}
		if app.driveId != 456 {
			t.Error("driveId field not set correctly")
		}
		if app.Applicant != applicant {
			t.Error("Applicant field not set correctly")
		}
		if app.Status() != Selected {
			t.Error("Status field not set correctly")
		}
	})

	t.Run("should handle nil applicant", func(t *testing.T) {
		app := &Application{
			id:        123,
			driveId:   456,
			Applicant: nil,
			status:    Applied,
		}

		if app.Applicant != nil {
			t.Error("Should handle nil applicant")
		}
	})
}

func TestApplicationStatus_Constants(t *testing.T) {
	t.Run("should have correct enum values", func(t *testing.T) {
		if Applied != 0 {
			t.Errorf("Applied should be 0, got %d", Applied)
		}
		if ShortListed != 1 {
			t.Errorf("ShortListed should be 1, got %d", ShortListed)
		}
		if Cleared != 2 {
			t.Errorf("Cleared should be 2, got %d", Cleared)
		}
		if Selected != 3 {
			t.Errorf("Selected should be 3, got %d", Selected)
		}
		if Rejected != 4 {
			t.Errorf("Rejected should be 4, got %d", Rejected)
		}
	})
}
