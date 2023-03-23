package club

import (
	"context"
	domainClub "golang-school-app-sample/src/domain/club"
	domainStudent "golang-school-app-sample/src/domain/student"
)

type ClubMemberUseCase struct {
	studentRepo domainStudent.StudnetRepository
	clubRepo    domainClub.ClubRepository
}

func NewClubUseCase(
	studentRepo domainStudent.StudnetRepository,
	clubRepo domainClub.ClubRepository,
) *ClubMemberUseCase {
	return &ClubMemberUseCase{studentRepo, clubRepo}
}

func (c *ClubMemberUseCase) Create(ctx context.Context, studentID string, clubID string) (res *domainClub.ClubMember, err error) {
	club, err := c.clubRepo.FindByID(ctx, clubID)
	if err != nil {
		return nil, err
	}

	student, err := c.studentRepo.FindByID(ctx, studentID)
	if err != nil {
		return nil, err
	}

	err = club.AddClubMember(club.ID(), student.ID())
	if err != nil {
		return nil, err
	}

	newClub, err := c.clubRepo.Update(ctx, club)
	if err != nil {
		return nil, err
	}

	// MEMO: 部員をresで返すための実装が冗長
	newClubMembers := newClub.ClubMembers()
	var newClubMember *domainClub.ClubMember
	for _, cm := range newClubMembers {
		if cm.ClubID() == club.ID() && cm.StudentID() == student.ID() {
			newClubMember = &cm
		}
	}

	return newClubMember, nil
}
