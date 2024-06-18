package auth

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(username, password string) (string, error) {
	return "", nil
}
