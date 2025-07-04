package internal

var newID = SeqID()

type Company struct {
	id     int
	name   string
	drives []*Drive
}

func NewCompany(name string) *Company {
	return &Company{id: newID(), name: name, drives: make([]*Drive, 0)}

}

func (c *Company) ID() int {
	return c.id
}

func (c *Company) Name() string {
	return c.name
}

func (c *Company) Drives() []*Drive {
	return c.drives
}

func (c *Company) AddDrive(drive *Drive) {
	c.drives = append(c.drives, drive)
}
