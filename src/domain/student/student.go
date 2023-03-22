package student

import "context"

type StudnetDBIF interface {
	Create(ctx context.Context, student *Student) (res *Student, err error)
	FindByID(ctx context.Context, studentID string) (res *Student, err error)
}

type Student struct {
	ID string
}
