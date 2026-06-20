package github

type Account struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Email    string `yaml:"email"`

	Host    string `yaml:"host"`
	KeyPath string `yaml:"key_path"`
}

type AccountStore struct {
	Accounts []Account `yaml:"accounts"`
}