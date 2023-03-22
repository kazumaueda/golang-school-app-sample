package club

import "context"

type ClubMemberDBIF interface {
	Create(ctx context.Context, clubMember *ClubMember) (res *ClubMember, err error)
}

type ClubMember struct {
	ClubID    string
	StudentID string
}

func NewClubMember(clubID, studentID string) *ClubMember {
	return &ClubMember{
		ClubID:    clubID,
		StudentID: studentID,
	}
}
