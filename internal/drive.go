package internal

import (
	"fmt"
	"time"
)

// Automatically generated ID generator
var nextID = SeqID()

// Enum for job category
type JobCategory int

const (
	Day JobCategory = iota
	Dream
	SuperDream
	Marquee
)

// Map for job category string representation
var JobCategoryStringMap = map[JobCategory]string{
	Day:        "Day Company",
	Dream:      "Dream",
	SuperDream: "Super Dream",
	Marquee:    "Marquee",
}

// String method for enum
func (jc JobCategory) String() string {
	return JobCategoryStringMap[jc]
}

// Eligibility struct
type Eligibility struct {
	requirement float64
}

// func (el Eligibility) checkEligibility(applicant *Applicant) bool {
// 	panic("unimplemented")
// }

// Drive struct
type Drive struct {
	id           int
	startDate    time.Time
	endDate      time.Time
	roleName     string
	eligibility  Eligibility
	ctc          int
	jobCategory  JobCategory
	applications []*Application
}

// --- Constructor Functions ---

func NewEligibility(minimumGPA float64) *Eligibility {
	return &Eligibility{requirement: minimumGPA}
}

func NewDrive(startDate time.Time, endDate time.Time, roleName string, minimumGPA float64, ctc int, jobCategory JobCategory) *Drive {
	return &Drive{id: nextID(), startDate: startDate, endDate: endDate, roleName: roleName, eligibility: *NewEligibility(minimumGPA), ctc: ctc, jobCategory: jobCategory}
}

// Elegibility setters and Getters
func (el *Eligibility) Requirement() float64 {
	return el.requirement
}

func (el *Eligibility) ChangeRequirement(newReq float64) {
	el.requirement = newReq
}

func (el *Eligibility) CheckEligibility(applicant *Applicant) bool {
	return applicant.CGPA >= el.requirement
}

// --- Drive Getters ---

func (dr Drive) ID() int {
	return dr.id
}

func (dr Drive) StartDate() time.Time {
	return dr.startDate
}

func (dr Drive) EndDate() time.Time {
	return dr.endDate
}

func (dr Drive) RoleName() string {
	return dr.roleName
}

func (dr Drive) Eligibility() *Eligibility {
	return &dr.eligibility
}

func (dr Drive) CTC() int {
	return dr.ctc
}

func (dr Drive) JobCategory() JobCategory {
	return dr.jobCategory
}

func (dr Drive) Applications() []*Application {
	return dr.applications
}

// --- Drive Setters ---

func (dr *Drive) SetStartDate(startDate time.Time) {
	dr.startDate = startDate
}

func (dr *Drive) SetEndDate(endDate time.Time) {
	dr.endDate = endDate
}

func (dr *Drive) SetRoleName(roleName string) {
	dr.roleName = roleName
}

func (dr *Drive) SetEligibility(minimumGPA float64) {
	dr.eligibility.ChangeRequirement(minimumGPA)
}

func (dr *Drive) SetCTC(ctc int) {
	dr.ctc = ctc
}

func (dr *Drive) SetJobCategory(jobCat JobCategory) {
	dr.jobCategory = jobCat
}

func (dr *Drive) AppendApplication(application *Application) {
	if application != nil{
		dr.applications = append(dr.applications, application)
	}
}

// --- Drive Functional Methods ---

func (dr *Drive) HasApplied(studentID int) bool {
	for _, e := range dr.applications {
		if e.Student.ID() == studentID {
			return true
		}
	}
	return false
}

func (dr *Drive) GetApplicationByID(id int) (*Application, error) {
	for _, e := range dr.applications {
		if e.ID() == id {
			return e, nil
		}
	}
	return nil, fmt.Errorf("no such application for given id")
}

func (dr *Drive) getSelectedApplications() []*Application {
	var arr []*Application
	for _, app := range dr.Applications() {
		if app.Status() == Selected {
			arr = append(arr, app)
		}
	}
	return arr
}

func (dr *Drive) getShortlistedApplications() []*Application {
	var arr []*Application
	for _, app := range dr.Applications() {
		if app.Status() == ShortListed {
			arr = append(arr, app)
		}
	}
	return arr
}

// Elegibility functions
func (el *Eligibility) checkEligibility(applicant *Applicant) bool {
	if el.requirement >= applicant.CGPA {
		return false
	} else {
		return true
	}
}
