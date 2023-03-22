package club

import "context"

type ClubDBIF interface {
	Create(ctx context.Context, club *Club) (res *Club, err error)
	Update(ctx context.Context, club *Club) (res *Club, err error)
	FindByID(ctx context.Context, clubID string) (res *Club, err error)
}

type Club struct {
	ID             string
	Name           string
	ApprovalStatus string
	ClubMembers    []ClubMember
}

func (c *Club) UpdateApprovalStatus() {
	if len(c.ClubMembers) >= 5 {
		c.ApprovalStatus = "approved"
	} else {
		c.ApprovalStatus = "notApproved"
	}
}
