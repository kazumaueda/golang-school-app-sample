package clubdm

import "context"

type ClubRepository interface {
	Create(ctx context.Context, club *Club) error
	Update(ctx context.Context, club *Club) error
	FindByID(ctx context.Context, clubID string) (*Club, error)
}
