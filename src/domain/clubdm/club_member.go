package clubdm

type ClubMember struct {
	studentID string
}

func newClubMember(studentID string) (clubMember *ClubMember, err error) {
	return &ClubMember{
		studentID: studentID,
	}, nil
}

func (c *ClubMember) StudentID() string {
	return c.studentID
}
