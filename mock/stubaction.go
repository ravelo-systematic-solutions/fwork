package mock

type StubAction interface {
	With(any)
	WithError(err error)
}

type stubAction struct {
	err         error
	withoutData bool
	payload     any
}

func (s *stubAction) With(payload any) {
	s.payload = payload
}

func (s *stubAction) WithError(err error) {
	s.err = err
}
