package github

type Service struct {
    repo Repository
    ssh  *SSHManager
}

func NewService(
	repo Repository,
	ssh *SSHManager,
) *Service

func (s *Service) AddAccount(
    name string,
    email string,
) (*Account, error)