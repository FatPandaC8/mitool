package github

type SSHManager struct{}

func (s *SSHManager) GenerateKey(
    name string,
    email string,
) (string, error)

func (s *SSHManager) AppendConfig(
    host string,
    keyPath string,
) error