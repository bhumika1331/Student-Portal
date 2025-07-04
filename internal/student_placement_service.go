package internal

import "fmt"

// StudentPlacementService is responsible for handling the placement process for a student in a specific drive.
// It encapsulates the student, the drive they are applying to, and the applicant information.
type StudentPlacementService struct {
	student   Student   // The student who is applying for placement
	drive     Drive     // The placement drive the student is applying to
	applicant Applicant // The applicant information (can be extended for more details)
	PlacementRegistrar
	Company
}

// NewStudentPlacementService creates and returns a new instance of StudentPlacementService.
// It takes a student and a drive as input and initializes the service for that student and drive.
func NewStudentPlacementService(student Student, drive Drive) *StudentPlacementService {
	return &StudentPlacementService{
		student:   student,
		drive:     drive,
		applicant: Applicant{Student: student},
	}
}

// Apply attempts to apply the student to the drive.
// It checks if the student (as an applicant) meets the eligibility criteria for the drive.
// If eligible, it appends the student's application to the drive and returns nil (no error).
// If not eligible, it returns an error indicating the student is not eligible for the drive.
// func (s *StudentPlacementService) Apply() error {
// 	elig := s.drive.Eligibility()               // Get the eligibility criteria for the drive
// 	if (&elig).CheckEligibility(&s.applicant) { // Check if the applicant meets the eligibility
// 		s.drive.AppendApplication(Application{id: s.student.id, driveId: s.drive.ID(), Applicant: Applicant{Student: s.student}}) // Add the application to the drive
// 		return nil
// 	}
// 	return errors.New("student is not eligible for the drive") // Return error if not eligible
// }

//Function for student to check companies applied

// Function to check for companies applicable
func (s *StudentPlacementService) CompaniesApplicable() []string {
	eligibleCompanies := make(map[string]struct{})
	for _, company := range s.PlacementRegistrar.companies {
		for _, drive := range company.drives {
			elig := drive.Eligibility()
			if elig.CheckEligibility(&s.applicant) {
				eligibleCompanies[company.name] = struct{}{}
				break // No need to check other drives for this company
			}
		}
	}
	companies := make([]string, 0, len(eligibleCompanies))
	for name := range eligibleCompanies {
		companies = append(companies, name)
	}
	return companies
}

func (s *StudentPlacementService) CompaniesApplied() interface{} {
	fmt.Println("Checking for companies applied")
	return s.applicant.CompaniesAppliedFor(&s.PlacementRegistrar)

}

// Function (Method) to view a particular offer details like CTC, Rolename, Category like Dream, Super Dream
func (s *StudentPlacementService) ViewOfferDetails(companyname string, driveid int) {
	fmt.Printf("The offer details are as follows \n  CTC:%d, Rolename:%s, Category:%s", s.drive.ctc, s.drive.roleName, s.drive.jobCategory)
}

// Register/Apply for a Placement Drive
func (s *StudentPlacementService) Apply() error {
	err := s.PlacementRegistrar.ApplyForDrive(s.student.id, s.Company.ID(), s.drive.id)
	return err
}

// GetDrive returns the drive associated with this placement service.
func (s *StudentPlacementService) GetDrive() Drive {
	return s.drive
}

// Here i am trying to view the shortlisted status i.e the status after shortlisting
func (s *StudentPlacementService) ViewShortlistStatus(id int) {
	applications := s.drive.getShortlistedApplications()
	for _, v := range applications {
		if v.id == id {
			fmt.Println("You just got shortlisted,prepare yourself for further rounds")
		}
	}

}

//TO-DO function to give students notifications related to new drives and results etc.
//Trying to create a notification interface and then creating different types of notifications
//and then sending them to the student
// TO-DO: Automatic trigger for notifications if last date of application is near
// Dont know how to do this, so will leave it for now

type Notification interface {
	Send() interface{}
}

// DriveNotification is a notification that is sent when anything related to drive is updated.
type DriveNotification struct {
	drive Drive
}

func (d *DriveNotification) Send() interface{} {
	return d.drive
}

// NewDriveNotification creates a new DriveNotification for a given drive.
func NewDriveNotification(drive Drive) Notification {
	return &DriveNotification{drive: drive}
}

// ResultNotification is a notification that is sent when the result of the drive is updated.
type ResultNotification struct {
	Days_left int
}

func (r *ResultNotification) Send() interface{} {
	return r.Days_left
}

func NewResultNotification(days_left int) Notification {
	return &ResultNotification{Days_left: days_left}
}
