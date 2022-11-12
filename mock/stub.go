package mock

type Mock interface {
	Query() *stubAction
}

type mock struct {
	queryStub *stubAction
}

func (s *mock) Query() *stubAction {
	s.queryStub = &stubAction{}
	return s.queryStub
}

func (s *mock) GetQueryStub() *stubAction {
	return s.queryStub
}

func NewSub() *mock {
	return &mock{}
}
