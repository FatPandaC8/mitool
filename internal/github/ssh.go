package github

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/FatPandaC8/mitool/internal/config"
)

func GenerateKey(name, email string) (string, error) {

	home := config.HomeDir()

	keyPath := filepath.Join(
		home,
		".ssh",
		"mitool_"+name,
	)


	err := os.MkdirAll(
		filepath.Dir(keyPath),
		0700,
	)

	if err != nil {
		return "", err
	}


	cmd := exec.Command(
		"ssh-keygen",
		"-t",
		"ed25519",
		"-C",
		email,
		"-f",
		keyPath,
		"-N",
		"",
	)

	if err := cmd.Run(); err != nil {
		return "", err
	}

	
	cmd_show := exec.Command(
		"cat",
		keyPath+".pub",
	)

	output, err := cmd_show.CombinedOutput() // this capture both stdout and stderr

	if err != nil {
		return "", fmt.Errorf(
			"show key failed: %v\n%s",
			err,
			output,
		)
	}


	fmt.Println(string(output))

	fmt.Println(
		"[github-ssh-account] created:",
		keyPath,
	)

	return keyPath, nil
}

func TestSSH(name string) error {

	host := "github-" + name

	cmd := exec.Command(
		"ssh",
		"-T",
		"git@"+host,
	)

	// ssh -T returns exit code 1 even on success
	// because GitHub does not open a shell.
	output, err := cmd.CombinedOutput()

	if err != nil {
		// Check output because GitHub success is a special case
		if len(output) > 0 {
			return nil
		}

		return err
	}

	return nil
}

func TestSSHAccount(
	name string,
) error {

	store, err := LoadAccounts()

	if err != nil {
		return err
	}


	for _, acc := range store.Accounts {

		if acc.Name == name {

			return TestSSH(
				acc.Host,
			)
		}
	}


	return fmt.Errorf(
		"account not found: %s",
		name,
	)
}