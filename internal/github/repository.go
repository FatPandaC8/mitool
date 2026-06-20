package github

import (
	"fmt"
	"os"

	"github.com/FatPandaC8/mitool/internal/config"
	"go.yaml.in/yaml/v4"
)

func LoadAccounts() (*AccountStore, error) {

	data, err := os.ReadFile(
		config.AccountsFile(),
	)

	if os.IsNotExist(err) {
		return &AccountStore{
			Accounts: []Account{},
		}, nil
	}

	if err != nil {
		return nil, err
	}

	var store AccountStore

	err = yaml.Unmarshal(
		data,
		&store,
	)

	return &store, err
}

func SaveAccounts(
	store *AccountStore,
) error {

	data, err := yaml.Marshal(store)

	if err != nil {
		return err
	}

	return os.WriteFile(
		config.AccountsFile(),
		data,
		0644,
	)
}

func RemoveAccount(
	name string,
) error {

	store, err := LoadAccounts()

	if err != nil {
		return err
	}


	var remaining []Account

	found := false

	for _, acc := range store.Accounts {

		if acc.Name == name {
			found = true
			continue
		}

		remaining = append(
			remaining,
			acc,
		)
	}

	if !found {
		return fmt.Errorf(
			"account not found: %s",
			name,
		)
	}

	store.Accounts = remaining

	return SaveAccounts(store)
}

func GetAccount(name string) (*Account, error) {
	store, err := LoadAccounts()

	if err != nil {
		return nil, err
	}

	// TODO: use another search
	for _, acc := range store.Accounts {
		if acc.Name == name {
			return &acc, nil
		}
	}

	return nil, fmt.Errorf(
		"[github-ssh-account]: account not found:%s",
		name,
	)
}