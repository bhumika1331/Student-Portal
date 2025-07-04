package internal

import (
	"fmt"
)

type PlacementRegistrar struct {
	companies    []*Company
	applications []*Application
	applicants   []*Applicant
}

type ReportByStudent struct {
	applicant        *Applicant
	offersRecived    []*Application
	eligibileRoles   []*Drive
	finalOffer       *Drive
	ctcForFinalOffer int
}

type ReportByDrive struct {
	drive                *Drive
	company              *Company
	driveCTC             int
	noOfSelectedStudents int
}

type FullPlacementReport struct {
	allComapanies         []*Company
	totalComapanies       int
	totalOffersMade       int
	allOffersMade         []*Application
	totalOffersByCatagory map[JobCategory]int
}

func (pr PlacementRegistrar) GenerateReportByDrive() ReportByDrive {
	var reportByDrive ReportByDrive
	for _, c := range pr.companies {
		for _, d := range c.Drives() {
			reportByDrive.company = c
			reportByDrive.drive = d
			reportByDrive.driveCTC = d.CTC()
			reportByDrive.noOfSelectedStudents = len(d.getSelectedApplications())
		}
	}
	return reportByDrive
}

func (pr PlacementRegistrar) GenerateReportByStudent() []ReportByStudent {
	var ReportsByStudent []ReportByStudent
	for _, e := range pr.applicants {
		report := ReportByStudent{}
		report.applicant = e
		for _, d := range pr.AllDrives() {
			if d.eligibility.checkEligibility(e) {
				report.eligibileRoles = append(report.eligibileRoles, d)
			}
		}
		for _, a := range pr.applications {
			if a.Applicant.ID() == e.ID() && a.status == Selected {
				report.offersRecived = append(report.offersRecived, a)
			}
		}
		if len(report.offersRecived) > 0 {
			firstOfferDriveID := report.offersRecived[0].driveId
			for _, d := range pr.AllDrives() {
				if d.ID() == firstOfferDriveID {
					report.finalOffer = d
					report.ctcForFinalOffer = d.CTC()
					break
				}
			}
		} else {
			report.finalOffer = nil
			report.ctcForFinalOffer = 0
		}
		// report.finalOffer = e.DrivesAppliedFor()[0]
		// report.ctcForFinalOffer = report.finalOffer.CTC()
		ReportsByStudent = append(ReportsByStudent, report)
	}
	return ReportsByStudent
}

func (pr PlacementRegistrar) GenerateFullReport() FullPlacementReport {
	var report FullPlacementReport
	allOffersBycatagory := make(map[JobCategory]int)
	report.allComapanies = pr.companies
	report.totalComapanies = len(report.allComapanies)
	for _, a := range pr.applications {
		if a.Status() == Selected {
			report.allOffersMade = append(report.allOffersMade, a)
		}
	}
	report.totalOffersMade = len(report.allOffersMade)
	for _, d := range pr.AllDrives() {
		jc := d.JobCategory()
		switch jc {
		case Day:
			allOffersBycatagory[Day]++
		case Dream:
			allOffersBycatagory[Dream]++
		case SuperDream:
			allOffersBycatagory[SuperDream]++
		case Marquee:
			allOffersBycatagory[Marquee]++
		}
	}
	report.totalOffersByCatagory = allOffersBycatagory
	return report
}

func (pr PlacementRegistrar) AllDrives() []*Drive {
	var alldrives []*Drive
	for _, c := range pr.companies {
		alldrives = append(alldrives, c.Drives()...)
	}
	return alldrives
}

func (pr *PlacementRegistrar) AddCompany(company *Company) { // AddCompany will help us add new company to PlacementRegistrar
	pr.companies = append(pr.companies, company)
}

func (pr *PlacementRegistrar) GetCompanyByID(id int) (*Company, error) {
	for i := range pr.companies {
		if pr.companies[i].id == id {
			return pr.companies[i], nil
		}
	}
	return nil, fmt.Errorf("company with id %d is not found", id)
}

func (pr *PlacementRegistrar) UpdateCompany(UpdatedCompany *Company) error {
	if UpdatedCompany == nil {
		return fmt.Errorf("cannot update nil company")
	}

	for i := range pr.companies {
		if pr.companies[i].id == UpdatedCompany.id {
			pr.companies[i] = UpdatedCompany
			return nil
		}
	}

	return fmt.Errorf("company with id %d not found to update", UpdatedCompany.id)
}

func (pr *PlacementRegistrar) AddDriveToCompany(companyID int, drive *Drive) error {
	for i := range pr.companies {
		if pr.companies[i].id == companyID {
			pr.companies[i].drives = append(pr.companies[i].drives, drive)
			return nil

		}
	}
	return fmt.Errorf("company with id %d not found", companyID)
}

func (pr *PlacementRegistrar) ApplicantByID(studentID int) (*Applicant, error) {
	for i := range pr.applicants {
		if pr.applicants[i].ID() == studentID {
			return pr.applicants[i], nil
		}
	}
	return nil, fmt.Errorf("applicant with id %d not found", studentID)
}

func (pr *PlacementRegistrar) CompanyByID(id int) (*Company, error) {
	for i := range pr.companies {
		if pr.companies[i].ID() == id {
			return pr.companies[i], nil
		}
	}
	return nil, fmt.Errorf("company with id %d not found", id)
}

func (pr *PlacementRegistrar) DriveByID(companyID, driveID int) (*Drive, error) {
	company, err := pr.CompanyByID(companyID)
	if err != nil {
		return nil, err
	}

	for i := range company.Drives() {
		if company.Drives()[i].ID() == driveID {
			drives := company.Drives()
			return drives[i], nil
		}
	}
	return nil, fmt.Errorf("drive by id  %d and company with id %d not found", driveID, companyID)
}

func (pr *PlacementRegistrar) ApplyForDrive(studentID, companyID, driveID int) error {
	applicant, err := pr.ApplicantByID(studentID)
	if err != nil {
		return fmt.Errorf("applicant not found: %v", err)
	}
	drive, err := pr.DriveByID(companyID, driveID)
	if err != nil {
		return err
	}

	if drive.HasApplied(studentID) {
		return fmt.Errorf("applicant applied already")
	}

	if !drive.eligibility.checkEligibility(applicant) {
		return fmt.Errorf("applicant is not meet the criteria ")
	}

	application := &Application{
		id:        len(pr.applications) + 1,
		driveId:   driveID,
		Applicant: applicant,
		status:    Applied,
	}

	pr.applications = append(pr.applications, application)
	drive.AppendApplication(application)

	for i := range pr.applicants {
		if pr.applicants[i].ID() == studentID {
			pr.applicants[i].AddDrivesAppliedFor(drive)
			break
		}
	}
	return nil
}

func (pr *PlacementRegistrar) UpdateApplicationStatus(studentID, driverID int, newStatus ApplicationStatus) error {
	for i := range pr.applications {
		app := pr.applications[i]
		if app.Student.id == studentID && app.driveId == driverID {
			app.status = newStatus
			return nil
		}
	}
	return fmt.Errorf("student with id %d status cannot be updated for drive with id %d", studentID, driverID)
}
