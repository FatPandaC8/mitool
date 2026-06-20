package commands

import (
	"fmt"

	"github.com/FatPandaC8/mitool/internal/github"
)

func Github(args []string) {

	if len(args) < 2 {
		fmt.Println(
			"usage: mitool github add <name>",
		)
		return
	}


	switch args[0] {

	case "add":
		name := args[1]

		key, err := github.GenerateKey(name)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = github.AddSSHConfig(
			name,
			key,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		store, err := github.LoadAccounts()
		if err != nil {
			fmt.Println(err)
			return
		}

		store.Accounts = append(
			store.Accounts,
			github.Account{
				Name: name,
				Host: "github-" + name,
				Key: key,
			},
		)

		err = github.SaveAccounts(store)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("account saved")

	case "remove":
		name := args[1]

		err := github.RemoveAccount(name)

		if err != nil {
			fmt.Println(err)
			return
		}


		err = github.RemoveSSHConfig(name)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(
			"removed account:",
			name,
		)

	case "test":

		name := args[1]

		err := github.TestSSHAccount(name)

		if err != nil {
			fmt.Println("SSH test failed:")
			fmt.Println(err)
			return
		}

		fmt.Println(
			"SSH works for",
			name,
		)

	case "list":
		store, err := github.LoadAccounts()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, acc := range store.Accounts {
			fmt.Println(acc.Name)
			fmt.Println(" host:", acc.Host)
			fmt.Println(" key:", acc.Key)
			fmt.Println()
		}

	default:
		fmt.Println("unknown github command")
	}
}