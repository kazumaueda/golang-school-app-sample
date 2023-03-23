package club

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

func (c *Club) UpdateApprovalStatus() {
	if len(c.clubMembers) >= 5 {
		c.approvalStatus = "approved"
	} else {
		c.approvalStatus = "notApproved"
	}
}
