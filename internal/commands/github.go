package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/FatPandaC8/mitool/internal/github"
)

func Github(args []string) {

	if len(args) < 1 {
		fmt.Println("usage: mitool github <command>")
		return
	}

	switch args[0] {

	case "add":
		if len(args) < 2 {
			fmt.Println("usage: mitool github add <name>")
			return
		}

		name := args[1]

		username := ask("github username: ")
		email := ask("email: ")

		key, err := github.GenerateKey(name, email)

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
				Name:     name,
				Username: username,
				Email:    email,
				Host:     "github-" + name,
				KeyPath:  key,
			},
		)


		err = github.SaveAccounts(store)

		if err != nil {
			fmt.Println(err)
			return
		}


		fmt.Println(
			"account saved:",
			name,
		)

		// TODO: show the .ssh/mitool_+name here for quick copy



	case "remove":

		if len(args) < 2 {
			fmt.Println("usage: mitool github remove <name>")
			return
		}

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

		if len(args) < 2 {
			fmt.Println("usage: mitool github test <name>")
			return
		}


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
			fmt.Println(" username:", acc.Username)
			fmt.Println(" email:", acc.Email)
			fmt.Println(" host:", acc.Host)
			fmt.Println(" key:", acc.KeyPath)
			fmt.Println()
		}



	case "use":

		if len(args) < 2 {
			fmt.Println("usage: mitool github use <name>")
			return
		}


		err := GithubUse(args[1])

		if err != nil {
			fmt.Println(err)
		}



	default:
		fmt.Println(
			"unknown github command:",
			args[0],
		)
	}
}



func GithubUse(name string) error {

	acc, err := github.GetAccount(name)

	if err != nil {
		return err
	}


	err = runGitConfig(
		"user.name",
		acc.Username,
	)

	if err != nil {
		return err
	}


	err = runGitConfig(
		"user.email",
		acc.Email,
	)

	if err != nil {
		return err
	}


	err = setRemote(acc.Host)

	if err != nil {
		return err
	}


	fmt.Println(
		"using github account:",
		acc.Name,
	)


	return nil
}



func runGitConfig(
	key string,
	value string,
) error {

	cmd := exec.Command(
		"git",
		"config",
		"--local",
		key,
		value,
	)

	return cmd.Run()
}



func setRemote(host string) error {

	cmd := exec.Command(
		"git",
		"remote",
		"get-url",
		"origin",
	)

	out, err := cmd.Output()

	if err != nil {
		return err
	}


	oldURL := strings.TrimSpace(
		string(out),
	)


	if !strings.Contains(
		oldURL,
		"github.com",
	) {
		return nil
	}


	newURL := strings.Replace(
		oldURL,
		"github.com",
		host,
		1,
	)


	return exec.Command(
		"git",
		"remote",
		"set-url",
		"origin",
		newURL,
	).Run()
}

func ask(
	message string,
) string {

	fmt.Print(message)

	reader := bufio.NewReader(
		os.Stdin,
	)

	value, _ := reader.ReadString('\n')

	return strings.TrimSpace(value)
}