package v1

type Service struct {
	Name      string
	Listeners []string
	Endpoints []map[string]string
}

func (s *Service) Validate() error {
	return nil
}
