package student

import "context"

type StudnetRepository interface {
	Create(ctx context.Context, student *Student) (res *Student, err error)
	FindByID(ctx context.Context, studentID string) (res *Student, err error)
}
