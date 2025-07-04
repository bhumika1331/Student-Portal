package internal

import "testing"

func TestNewCompany(t *testing.T) {
	t.Run("should create company with name", func(t *testing.T) {
		c := NewCompany("TestCo")
		if c.Name() != "TestCo" {
			t.Errorf("Expected name 'TestCo', got '%s'", c.Name())
		}
		if c.ID() <= 0 {
			t.Error("ID should be positive")
		}
		if c.Drives() == nil {
			t.Error("Drives should be initialized")
		}
		if len(c.Drives()) != 0 {
			t.Error("Drives should be empty initially")
		}
	})

	t.Run("should create company with empty name", func(t *testing.T) {
		c := NewCompany("")
		if c.Name() != "" {
			t.Error("Should handle empty name")
		}
		if c.ID() <= 0 {
			t.Error("ID should still be positive for empty name")
		}
	})

	t.Run("should generate unique IDs", func(t *testing.T) {
		c1 := NewCompany("Company1")
		c2 := NewCompany("Company2")
		if c1.ID() == c2.ID() {
			t.Error("Each company should have unique ID")
		}
	})

	t.Run("should handle special characters in name", func(t *testing.T) {
		specialName := "Test & Co. Ltd. #1"
		c := NewCompany(specialName)
		if c.Name() != specialName {
			t.Error("Should handle special characters in name")
		}
	})
}

func TestCompany_ID(t *testing.T) {
	t.Run("should return auto-generated positive ID", func(t *testing.T) {
		c := NewCompany("TestCo")
		if c.ID() <= 0 {
			t.Error("ID should be positive")
		}
	})

	t.Run("should return consistent ID", func(t *testing.T) {
		c := NewCompany("TestCo")
		id1 := c.ID()
		id2 := c.ID()
		if id1 != id2 {
			t.Error("ID should be consistent across calls")
		}
	})
}

func TestCompany_Name(t *testing.T) {
	t.Run("should return correct name", func(t *testing.T) {
		c := NewCompany("TestCo")
		if c.Name() != "TestCo" {
			t.Error("Name() incorrect")
		}
	})

	t.Run("should return empty name", func(t *testing.T) {
		c := NewCompany("")
		if c.Name() != "" {
			t.Error("Name() should return empty string")
		}
	})

	t.Run("should return name with special characters", func(t *testing.T) {
		c := NewCompany("Test & Co. Ltd.")
		if c.Name() != "Test & Co. Ltd." {
			t.Error("Name() should handle special characters")
		}
	})

	t.Run("should return very long name", func(t *testing.T) {
		longName := "This is a very long company name that might be used in some cases"
		c := NewCompany(longName)
		if c.Name() != longName {
			t.Error("Name() should handle long names")
		}
	})
}

func TestCompany_Drives(t *testing.T) {
	t.Run("should return empty drives initially", func(t *testing.T) {
		c := NewCompany("TestCo")
		drives := c.Drives()
		if len(drives) != 0 {
			t.Error("Drives() should return empty slice initially")
		}
	})

	t.Run("should return all drives after adding", func(t *testing.T) {
		c := NewCompany("TestCo")
		d1 := &Drive{id: 101}
		d2 := &Drive{id: 102}
		c.AddDrive(d1)
		c.AddDrive(d2)
		drives := c.Drives()

		if len(drives) != 2 {
			t.Errorf("Expected 2 drives, got %d", len(drives))
		}
		if drives[0] != d1 || drives[1] != d2 {
			t.Error("Drives() should return correct drives")
		}
	})

	t.Run("should return same reference", func(t *testing.T) {
		c := NewCompany("TestCo")
		d1 := &Drive{id: 101}
		c.AddDrive(d1)

		drives1 := c.Drives()
		drives2 := c.Drives()

		if &drives1[0] != &drives2[0] {
			t.Error("Drives() should return same reference")
		}
	})

	t.Run("should not be nil even when empty", func(t *testing.T) {
		c := NewCompany("TestCo")
		drives := c.Drives()
		if drives == nil {
			t.Error("Drives() should not return nil")
		}
	})
}

func TestCompany_AddDrive(t *testing.T) {
	t.Run("should add single drive", func(t *testing.T) {
		c := NewCompany("Test2")
		d1 := &Drive{id: 101}
		c.AddDrive(d1)

		drives := c.Drives()
		if len(drives) != 1 || drives[0] != d1 {
			t.Error("AddDrive failed to add single drive")
		}
	})

	t.Run("should add multiple drives", func(t *testing.T) {
		c := NewCompany("Test2")
		d1 := &Drive{id: 101}
		d2 := &Drive{id: 102}
		c.AddDrive(d1)
		c.AddDrive(d2)

		drives := c.Drives()
		if len(drives) != 2 || drives[0] != d1 || drives[1] != d2 {
			t.Error("AddDrive failed to add multiple drives")
		}
	})

	t.Run("should add nil drive", func(t *testing.T) {
		c := NewCompany("Test2")
		c.AddDrive(nil)

		drives := c.Drives()
		if len(drives) != 1 {
			t.Error("AddDrive should add nil drive")
		}
		if drives[0] != nil {
			t.Error("AddDrive should store nil drive")
		}
	})

	t.Run("should add duplicate drives", func(t *testing.T) {
		c := NewCompany("Test2")
		d1 := &Drive{id: 101}
		c.AddDrive(d1)
		c.AddDrive(d1)

		drives := c.Drives()
		if len(drives) != 2 {
			t.Error("AddDrive should allow duplicate drives")
		}
		if drives[0] != d1 || drives[1] != d1 {
			t.Error("AddDrive should store duplicate drives")
		}
	})

	t.Run("should maintain order of drives", func(t *testing.T) {
		c := NewCompany("Test2")
		drives := []*Drive{
			{id: 101}, {id: 102}, {id: 103}, {id: 104}, {id: 105},
		}

		for _, d := range drives {
			c.AddDrive(d)
		}

		result := c.Drives()
		if len(result) != 5 {
			t.Errorf("Expected 5 drives, got %d", len(result))
		}

		for i, d := range drives {
			if result[i] != d {
				t.Errorf("Drive order not maintained at index %d", i)
			}
		}
	})
}

func TestCompany_Integration(t *testing.T) {
	t.Run("should work correctly with multiple operations", func(t *testing.T) {
		c := NewCompany("Integration Test Co")

		// Verify initial state
		if c.Name() != "Integration Test Co" {
			t.Error("Name not set correctly")
		}
		if len(c.Drives()) != 0 {
			t.Error("Should start with no drives")
		}
		if c.ID() <= 0 {
			t.Error("Should have positive ID")
		}

		// Add drives and verify
		d1 := &Drive{id: 201}
		d2 := &Drive{id: 202}
		c.AddDrive(d1)
		c.AddDrive(d2)

		if len(c.Drives()) != 2 {
			t.Error("Should have 2 drives after adding")
		}

		// Verify drives are in correct order
		drives := c.Drives()
		if drives[0] != d1 || drives[1] != d2 {
			t.Error("Drives not in correct order")
		}
	})

	t.Run("should handle edge cases gracefully", func(t *testing.T) {
		// Company with very long name
		longName := string(make([]byte, 1000))
		c := NewCompany(longName)
		if c.Name() != longName {
			t.Error("Should handle very long names")
		}

		// Add many drives
		for i := 0; i < 100; i++ {
			c.AddDrive(&Drive{id: i})
		}

		if len(c.Drives()) != 100 {
			t.Error("Should handle many drives")
		}
	})
}

func TestCompany_IDUniqueness(t *testing.T) {
	t.Run("should generate unique IDs for multiple companies", func(t *testing.T) {
		companies := make([]*Company, 10)
		idSet := make(map[int]bool)

		for i := 0; i < 10; i++ {
			companies[i] = NewCompany("Company" + string(rune('A'+i)))
			id := companies[i].ID()

			if idSet[id] {
				t.Errorf("Duplicate ID found: %d", id)
			}
			idSet[id] = true
		}

		if len(idSet) != 10 {
			t.Errorf("Expected 10 unique IDs, got %d", len(idSet))
		}
	})
}
