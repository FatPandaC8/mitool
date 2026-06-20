package github

type Account struct {
	Name  string `yaml:"name"`
    Email string `yaml:"email"`

    Host  string `yaml:"host"`
    Key   string `yaml:"key"`
}

type AccountStore struct {
	Accounts []Account `yaml:"accounts"`
}