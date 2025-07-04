package internal

import (
	"encoding/json"
	"errors"

	"os"
)

// UpdateStudentName updates the name of a student by ID in the given slice.
func UpdateStudentName(students []Student, id int, newName string) error {
	for i, s := range students {
		if s.ID() == id {
			students[i].name = newName
			return nil
		}
	}
	return errors.New("student not found")
}

// FindStudentByID returns a pointer to the student with the given ID, or nil if not found.
func FindStudentByID(students []Student, id int) *Student {
	for i, s := range students {
		if s.ID() == id {
			return &students[i]
		}
	}
	return nil
}

// FindStudentsByName returns a slice of students matching the given name (case-insensitive).
func FindStudentsByName(students []Student, name string) []Student {
	var result []Student
	for _, s := range students {
		if s.Name() == name {
			result = append(result, s)
		}
	}
	return result
}

// DeleteStudentByID removes a student by ID from the slice and returns the new slice.
func DeleteStudentByID(students []Student, id int) ([]Student, error) {
	for i, s := range students {
		if s.ID() == id {
			return append(students[:i], students[i+1:]...), nil
		}
	}
	return students, errors.New("student not found")
}

// SerializeStudents serializes the students slice to JSON and writes it to filename(user input).
func SerializeStudents(filename string, students []Student) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(students)

	if err != nil {
		return err
	}
	return nil
}

// DeserializeStudents deserializes JSON data into a slice of students.
func DeserializeStudents(data []byte) ([]Student, error) {
	var students []Student
	err := json.Unmarshal(data, &students)
	return students, err
}
