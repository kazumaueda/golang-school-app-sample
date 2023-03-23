package club

type ClubMember struct {
	clubID    string
	studentID string
}

func (c *ClubMember) ClubID() string {
	return c.clubID
}

func (c *ClubMember) StudentID() string {
	return c.studentID
}

func NewClubMember(clubID, studentID string) *ClubMember {
	return &ClubMember{
		clubID:    clubID,
		studentID: studentID,
	}
}
