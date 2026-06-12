package github

type Repository interface {
	Save(Account) error
	FindAll() ([]Account, error)
	FindByName(name string) (*Account, error)
}