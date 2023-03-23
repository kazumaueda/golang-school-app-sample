package club

import (
	"context"
	domainClub "golang-school-app-sample/src/domain/club"
	domainStudent "golang-school-app-sample/src/domain/student"
)

type ClubMemberUseCase struct {
	studentRepo    domainStudent.StudnetDBIF
	clubMemberRepo domainClub.ClubMemberDBIF
	clubRepo       domainClub.ClubDBIF
}

func NewClubUseCase(
	studentRepo domainStudent.StudnetDBIF,
	clubMemberRepo domainClub.ClubMemberDBIF,
	clubRepo domainClub.ClubDBIF,
) *ClubMemberUseCase {
	return &ClubMemberUseCase{studentRepo, clubMemberRepo, clubRepo}
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

	cm := domainClub.NewClubMember(club.ID(), student.ID())
	clubMember, err := c.clubMemberRepo.Create(ctx, cm)
	if err != nil {
		return nil, err
	}

	// TODO: 承認状態の更新漏れを防ぐためビジネスロジックの内部に隠蔽したい
	club.UpdateApprovalStatus()
	_, err = c.clubRepo.Update(ctx, club)
	if err != nil {
		return nil, err
	}

	return clubMember, nil
}
