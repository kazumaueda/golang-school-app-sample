package studentdm

import "context"

type StudentRepository interface {
	Create(ctx context.Context, student *Student) error
	FindByID(ctx context.Context, studentID string) (*Student, error)
}
