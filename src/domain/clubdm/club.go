package clubdm

type Club struct {
	id             string
	name           string
	approvalStatus string
	clubMembers    []ClubMember
}

func (c *Club) ID() string {
	return c.id
}

func (c *Club) Name() string {
	return c.name
}

func (c *Club) ApprovalStatus() string {
	return c.approvalStatus
}

func (c *Club) ClubMembers() []ClubMember {
	return c.clubMembers
}

func (c *Club) AddClubMember(studentID string) error {
	clubMember, err := newClubMember(studentID)
	if err != nil {
		return err
	}

	c.clubMembers = append(c.clubMembers, *clubMember)
	c.updateApprovalStatus()

	return nil
}

func (c *Club) updateApprovalStatus() {
	if len(c.clubMembers) >= 5 {
		c.approvalStatus = "approved"
	} else {
		c.approvalStatus = "notApproved"
	}
}
