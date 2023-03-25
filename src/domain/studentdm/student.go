package studentdm

type Student struct {
	id string
}

func (s *Student) ID() string {
	return s.id
}
