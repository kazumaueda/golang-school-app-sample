package club

import "context"

type ClubRepository interface {
	Create(ctx context.Context, club *Club) (res *Club, err error)
	Update(ctx context.Context, club *Club) (res *Club, err error)
	FindByID(ctx context.Context, clubID string) (res *Club, err error)
}

type ClubMemberRepository interface {
	Create(ctx context.Context, clubMember *ClubMember) (res *ClubMember, err error)
}
