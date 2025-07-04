package internal

import (
	"fmt"
)

type Applicant struct {
	Student
	AcademicRecord
	drivesAppliedFor []*Drive
	offersReceived   []*Drive

}

func NewApplicant(st Student, ar AcademicRecord) *Applicant {
	return &Applicant{Student: st, AcademicRecord: ar, drivesAppliedFor: make([]*Drive, 0), offersReceived: make([]*Drive, 0)}
}

func (a *Applicant) getAllReceivedOffersDrivesAndApplications() ([]*Drive, []*Application) {
	var drarr []*Drive
	var pparr []*Application
	for _, d := range a.drivesAppliedFor {
		for _, app := range d.applications {
			if app.Status() == Selected && app.Applicant.ID() == a.ID() {
				drarr = append(drarr, d)
				pparr = append(pparr, app)
			}
		}
	}
	return drarr, pparr
}

func (a *Applicant) getFinalOffer() (int, error) {
	drArr, _ := a.getAllReceivedOffersDrivesAndApplications()
	if len(drArr) == 0 {
		return -1, fmt.Errorf("no offers yet")
	} else {
		return drArr[0].CTC(), nil
	}
}

func (a *Applicant) DrivesAppliedFor() []*Drive {
	return a.drivesAppliedFor
}

func (a *Applicant) AddDrivesAppliedFor(drive *Drive) {
	a.drivesAppliedFor = append(a.drivesAppliedFor, drive)
}

func (a *Applicant) SetDrivesAppliedFor(drives []*Drive) {
	a.drivesAppliedFor = drives
}

func (a *Applicant) CompaniesAppliedFor(pr *PlacementRegistrar) []string {
	companySet := make(map[string]struct{})

	for _, app := range pr.applications {
		if app.Applicant.Student.id == a.Student.id {
			for _, company := range pr.companies {
				for _, drive := range company.drives {
					if drive.id == app.driveId {
						companySet[company.name] = struct{}{} // in the companyset the struct only checks if the key is present and not the value , its more memory efficent and holds all company name in a map
					}
				}
			}
		}
	}
	companies := make([]string, 0, len(companySet))
	for name := range companySet {
		companies = append(companies, name)
	}
	return companies
}

func (a *Applicant) TotalNumberOfCompaniesAppliedFor(pr *PlacementRegistrar) int {
	return len(a.CompaniesAppliedFor(pr))
}
