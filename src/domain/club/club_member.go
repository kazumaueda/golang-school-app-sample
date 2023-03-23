package club

type ClubMember struct {
	clubID    string
	studentID string
}

// MEMO: Privateにして同一集約内のみインスタンスを生成できるようにしている
func newClubMember(clubID, studentID string) (clubMember *ClubMember, err error) {
	return &ClubMember{
		clubID:    clubID,
		studentID: studentID,
	}, nil
}

func (c *ClubMember) ClubID() string {
	return c.clubID
}

func (c *ClubMember) StudentID() string {
	return c.studentID
}
