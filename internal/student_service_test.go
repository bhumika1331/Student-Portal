package internal

import (
	"os"
	"testing"
)

func sampleStudents() []Student {
	return []Student{
		NewStudent(1, "Alice"),
		NewStudent(2, "Bob"),
		NewStudent(3, "Charlie"),
	}
}

func TestStudentService_UpdateStudentName(t *testing.T) {
	students := sampleStudents()
	err := UpdateStudentName(students, 2, "Bobby")
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if students[1].Name() != "Bobby" {
		t.Errorf("expected name to be 'Bobby', got: %s", students[1].Name())
	}
}

func TestStudentService_UpdateStudentName_NotFound(t *testing.T) {
	students := sampleStudents()
	err := UpdateStudentName(students, 999, "Ghost")
	if err == nil {
		t.Error("expected error for non-existent student ID")
	}
}

func TestStudentService_UpdateStudentName_EmptySlice(t *testing.T) {
	err := UpdateStudentName([]Student{}, 1, "Test")
	if err == nil {
		t.Error("expected error for empty slice, got nil")
	}
}

func TestStudentService_FindStudentByID_Found(t *testing.T) {
	students := sampleStudents()
	s := FindStudentByID(students, 3)
	if s == nil || s.Name() != "Charlie" {
		t.Error("expected to find Charlie")
	}
}

func TestStudentService_FindStudentByID_NotFound(t *testing.T) {
	students := sampleStudents()
	s := FindStudentByID(students, 999)
	if s != nil {
		t.Error("expected nil for non-existent student")
	}
}

func TestStudentService_FindStudentByID_EmptySlice(t *testing.T) {
	res := FindStudentByID([]Student{}, 1)
	if res != nil {
		t.Error("expected nil for empty slice")
	}
}

func TestStudentService_FindStudentsByName_Basic(t *testing.T) {
	students := sampleStudents()
	students = append(students, NewStudent(4, "Alice"))
	found := FindStudentsByName(students, "Alice")
	if len(found) != 2 {
		t.Errorf("expected 2 students named Alice, got: %d", len(found))
	}
}

func TestStudentService_FindStudentsByName_Unknown(t *testing.T) {
	students := sampleStudents()
	found := FindStudentsByName(students, "Unknown")
	if len(found) != 0 {
		t.Errorf("expected 0 results for unknown name, got: %d", len(found))
	}
}

func TestStudentService_FindStudentsByName_EmptySlice(t *testing.T) {
	res := FindStudentsByName([]Student{}, "Test")
	if len(res) != 0 {
		t.Errorf("expected 0, got %d", len(res))
	}
}

func TestStudentService_FindStudentsByName_CaseSensitive(t *testing.T) {
	students := []Student{NewStudent(1, "Alice")}
	res := FindStudentsByName(students, "alice")
	if len(res) != 0 {
		t.Error("expected 0 for case-sensitive mismatch")
	}
}

func TestStudentService_DeleteStudentByID_Basic(t *testing.T) {
	students := sampleStudents()
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
}

func TestStudentService_DeleteStudentByID_NotFound(t *testing.T) {
	students := sampleStudents()
	_, err := DeleteStudentByID(students, 999)
	if err == nil {
		t.Error("expected error when deleting non-existent student")
	}
}

func TestStudentService_DeleteStudentByID_EmptySlice(t *testing.T) {
	newList, err := DeleteStudentByID([]Student{}, 1)
	if err == nil {
		t.Error("expected error for empty slice")
	}
	if len(newList) != 0 {
		t.Errorf("expected 0, got %d", len(newList))
	}
}

func TestStudentService_DeleteStudentByID_DeleteOnlyStudent(t *testing.T) {
	students := []Student{NewStudent(1, "Solo")}
	newList, err := DeleteStudentByID(students, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(newList) != 0 {
		t.Errorf("expected 0, got %d", len(newList))
	}
}

func TestStudentService_SerializeStudents_Basic(t *testing.T) {
	students := sampleStudents()
	filename := "test_students.json"
	err := SerializeStudents(filename, students)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	_ = os.Remove(filename)
}

func TestStudentService_SerializeStudents_InvalidFilename(t *testing.T) {
	students := []Student{NewStudent(1, "Alice")}
	err := SerializeStudents("/invalid_path/\x00students.json", students)
	if err == nil {
		t.Error("expected error for invalid filename")
	}
}

func TestStudentService_DeserializeStudents_Valid(t *testing.T) {
	jsonData := `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`
	students, err := DeserializeStudents([]byte(jsonData))
	if err != nil {
		t.Errorf("deserialization failed: %v", err)
	}
	if len(students) != 2 || students[0].Name() != "Alice" {
		t.Error("unexpected deserialized student data")
	}
}

func TestStudentService_DeserializeStudents_InvalidJSON(t *testing.T) {
	badData := []byte(`{invalid json}`)
	_, err := DeserializeStudents(badData)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestStudentService_DeserializeStudents_EmptyData(t *testing.T) {
	students, err := DeserializeStudents([]byte(""))
	if err == nil {
		t.Error("expected error for empty data")
	}
	if students != nil && len(students) != 0 {
		t.Errorf("expected empty slice, got %v", students)
	}
}
