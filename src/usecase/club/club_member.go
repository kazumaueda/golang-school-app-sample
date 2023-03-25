package club

import (
	"context"
	"golang-school-app-sample/src/domain/clubdm"
	"golang-school-app-sample/src/domain/studentdm"
)

type ClubMemberUseCase struct {
	studentRepo studentdm.StudentRepository
	clubRepo    clubdm.ClubRepository
}

func NewClubUseCase(
	studentRepo studentdm.StudentRepository,
	clubRepo clubdm.ClubRepository,
) *ClubMemberUseCase {
	return &ClubMemberUseCase{studentRepo, clubRepo}
}

func (c *ClubMemberUseCase) Create(ctx context.Context, studentID string, clubID string) error {
	club, err := c.clubRepo.FindByID(ctx, clubID)
	if err != nil {
		return err
	}

	student, err := c.studentRepo.FindByID(ctx, studentID)
	if err != nil {
		return err
	}

	err = club.AddClubMember(student.ID())
	if err != nil {
		return err
	}

	err = c.clubRepo.Update(ctx, club)
	if err != nil {
		return err
	}

	return nil
}
